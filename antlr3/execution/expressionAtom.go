package execution

import (
	"../parser"
	"bytes"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"math"
	"strconv"
	"strings"
)

var ConstantExpressionAtomContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.ConstantExpressionAtomContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.ConstantExpressionAtomContext", t)
	}
	return ExecFuncByType(ctx, c.Constant())
}

var StringConstContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.StringConstantContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.StringConstantContext", t)
	}
	content := c.StringLiteral().GetText()
	return content, nil
}

var NumberConstantContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.NumberConstantContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.NumberConstantContext", t)
	}
	content := c.DecimalLiteral().GetText()
	number, err := strconv.ParseFloat(content, 64)
	if err != nil {
		return nil, err
	}
	if c.GetMinusFlag() != nil {
		number = -number
	}
	return number, nil
}

var BoolConstantContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.BoolConstantContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.BoolConstantContext", t)
	}
	content := c.BooleanLiteral().GetText()
	return strings.EqualFold(content, "true"), nil
}

var ColumnNameContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.ColumnNameContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.ColumnNameContext", t)
	}
	cn := c.GetCn().GetText()
	table := ctx.Row.Meta
	contents := ctx.Row.Content

	for i := 0; i < len(table.Fields); i++ {
		if table.Fields[i].Name == cn {
			return contents[i], nil
		}
	}
	return nil, fmt.Errorf("can't find proper colume %s", cn);
}

var ColumnNameExpressionAtomContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.ColumnNameExpressionAtomContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.ColumnNameExpressionAtomContext", t)
	}
	cn := c.ColumnName()
	return ColumnNameContextFunc(ctx, cn)
}

var FunctionCallExpressionAtomContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.FunctionCallExpressionAtomContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.FunctionCallExpressionAtomContext", t)
	}
	fc := c.FunctionCall()
	return ExecFuncByType(ctx, fc)
}

var FunctionArgsContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.FunctionArgsContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.FunctionArgsContext", t)
	}
	args := make([]interface{}, 0, 0)
	for _, ex := range c.AllExpression() {
		r, err := ExecFuncByType(ctx, ex)
		if err != nil {
			return nil, NewAntlrErr(fmt.Sprintf("compute function arg {%s} faild of {%s}", ex.GetText(), err.Error()))
		}
		args = append(args, r)
	}
	return args, nil
}

var UdfFunctionCallContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.UdfFunctionCallContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.UdfFunctionCallContext", t)
	}
	fn := c.Uid().GetText()
	udf, ok := GetUDF(fn)
	if !ok {
		return nil, NewAntlrErr(fmt.Sprintf("can not find udf with key {%s}", fn))
	}
	argsExpr := c.FunctionArgs()
	args, err := FunctionArgsContextFunc(ctx, argsExpr)
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("exec udf {%s} failed of %s", fn, err.Error()))
	}
	args2, _ := args.([]interface{})
	return udf(args2)
}

var AggregateFunctionCallContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.AggregateFunctionCallContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.AggregateFunctionCallContext", t)
	}
	return ExecFuncByType(ctx, c.AggregateWindowedFunction())
}

var CommonAggregateFuncContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.CommonAggregateFuncContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.AggregateFunctionCallContext", t)
	}
	aggFuncName := c.GetAggFunc().GetText()
	aggFunc := aggregateFuncMap[aggFuncName]
	if aggFunc == nil {
		return nil, NewAntlrErr(fmt.Sprintf("aggregate func {%s} not found", aggFuncName))
	}
	argExpr := c.Expression()
	aggregator := "ALL"
	if c.GetAggregator() != nil {
		aggregator = c.GetAggregator().GetText()
	}

	args := make([]interface{}, 0, len(ctx.Rows.rows))
	for _, r := range ctx.Rows.rows {
		nctx := &SqlCtx{0, r, nil}
		res, err := ExecFuncByType(nctx, argExpr)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		args = append(args, res)
	}
	return aggFunc(args, aggregator), nil
}

var CountFunc1ContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	_, e := t.(*parser.CountFunc1Context)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.AggregateFunctionCallContext", t)
	}
	return len(ctx.Rows.rows), nil
}

var CountFunc2ContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.CountFunc2Context)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.AggregateFunctionCallContext", t)
	}
	aggregator := c.GetAggregator().GetText()
	if strings.EqualFold(aggregator, "ALL") {
		//todo 这里需要确定一下count(column1, column2)表达的语义是什么
		return len(ctx.Rows.rows), nil
	}
	functionArgs := c.FunctionArgs().(*parser.FunctionArgsContext)
	temp := make(map[string]bool)
	for _, r := range ctx.Rows.rows {
		nctx := &SqlCtx{0, r, nil}
		var buffer bytes.Buffer
		for _, arg := range functionArgs.AllExpression() {
			res, err := ExecFuncByType(nctx, arg)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			buffer.WriteString(fmt.Sprintf("%s", res))
		}
		temp[buffer.String()] = true
	}
	return len(temp), nil
}

var UnaryExpressionAtomContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.UnaryExpressionAtomContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.UnaryExpressionAtomContext", t)
	}
	ea := c.ExpressionAtom()
	res, err := ExecFuncByType(ctx, ea)
	if err != nil {
		return nil, err
	}
	up := c.UnaryOperator().GetText()
	//todo 如何定义下面这些行为
	switch up {
	case "!":
	case "~":
	case "+":
	case "-":
	case "NOT":
	case "not":
	}
	return res, nil
}

var BracketedExpressionAtomContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.BracketedExpressionAtomContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.BracketedExpressionAtomContext", t)
	}
	return ExecFuncByType(ctx, c.Expression())
}

var MathExpressionAtomContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.MathExpressionAtomContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.MathExpressionAtomContext", t)
	}
	le := c.GetLeft()
	re := c.GetRight()
	lRes, err := ExecFuncByType(ctx, le)
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("resolve left expression {%s} in MathExpression failed of {%s}",
			le.GetText(), err.Error()))
	}
	lf, err := parseVal2Float64(lRes)
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("parse left expression result {%s} to float/int failed of {%s}",
			lRes, err.Error()))
	}
	rRes, err := ExecFuncByType(ctx, re)
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("resolve right expression {%s} in MathExpression failed of {%s}",
			re.GetText(), err.Error()))
	}
	rf, err := parseVal2Float64(lRes)
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("parse right expression result {%s} to float/int failed of {%s}",
			rRes, err.Error()))
	}
	op := c.MathOperator().GetText()
	var res float64
	switch op {
	case "*":
		res = lf * rf
	case "/":
		res = lf / rf
	case "%":
		res = float64(int64(lf) % int64(rf))
	case "+":
		res = lf + rf
	case "-":
		res = lf - rf
	default:
		return nil, NewAntlrErr(fmt.Sprintf("unsupported math opertion {%s}", op))
	}
	return res, nil
}

var aggregateFuncMap = make(map[string]func(args []interface{}, aggregator string) interface{})

func avgFunc(args []interface{}, aggregator string) interface{} {
	sum, count := 0.0, 0
	for _, arg := range args {
		val, err := parseVal2Float64(arg)
		if err != nil {
			fmt.Errorf("parse %s to float failed ", arg)
			continue
		}
		sum += val
		count++
	}
	if count == 0 {
		return 0
	}
	return sum / float64(count)
}

func maxFunc(args []interface{}, aggregator string) interface{} {
	var max float64
	first := true
	for _, arg := range args {
		val, err := parseVal2Float64(arg)
		if err != nil {
			fmt.Errorf("parse %s to float failed ", arg)
			continue
		}
		if first {
			max = val
		} else {
			max = math.Max(max, val)
		}
	}
	return max
}

func minFunc(args []interface{}, aggregator string) interface{} {
	var max float64
	first := true
	for _, arg := range args {
		val, err := parseVal2Float64(arg)
		if err != nil {
			fmt.Errorf("parse %s to float failed ", arg)
			continue
		}
		if first {
			max = val
		} else {
			max = math.Min(max, val)
		}
	}
	return max
}

func sumFunc(args []interface{}, aggregator string) interface{} {
	var sum float64
	for _, arg := range args {
		val, err := parseVal2Float64(arg)
		if err != nil {
			fmt.Errorf("parse %s to float failed ", arg)
			continue
		}
		sum += val
	}
	return sum
}

func init() {
	RegisterParseFunc("ConstantExpressionAtomContext", ConstantExpressionAtomContextFunc)
	RegisterParseFunc("StringConstContext", StringConstContextFunc)
	RegisterParseFunc("NumberConstantContext", NumberConstantContextFunc)
	RegisterParseFunc("BoolConstantContext", BoolConstantContextFunc)
	RegisterParseFunc("ColumnNameContext", ColumnNameContextFunc)
	RegisterParseFunc("ColumnNameExpressionAtomContext", ColumnNameExpressionAtomContextFunc)
	RegisterParseFunc("FunctionCallExpressionAtomContext", FunctionCallExpressionAtomContextFunc)
	RegisterParseFunc("FunctionArgsContext", FunctionArgsContextFunc)
	RegisterParseFunc("UdfFunctionCallContext", UdfFunctionCallContextFunc)
	RegisterParseFunc("AggregateFunctionCallContext", AggregateFunctionCallContextFunc)
	RegisterParseFunc("CommonAggregateFuncContext", CommonAggregateFuncContextFunc)
	RegisterParseFunc("CountFunc1Context", CountFunc1ContextFunc)
	RegisterParseFunc("CountFunc2Context", CountFunc2ContextFunc)
	RegisterParseFunc("UnaryExpressionAtomContext", UnaryExpressionAtomContextFunc)
	RegisterParseFunc("BracketedExpressionAtomContext", BracketedExpressionAtomContextFunc)
	RegisterParseFunc("MathExpressionAtomContext", MathExpressionAtomContextFunc)

	aggregateFuncMap["AVG"] = avgFunc
	aggregateFuncMap["MAX"] = maxFunc
	aggregateFuncMap["MIN"] = minFunc
	aggregateFuncMap["SUM"] = sumFunc
}
