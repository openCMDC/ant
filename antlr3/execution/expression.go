package execution

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)
import "ant/antlr3/parser"

var PredicateExpressionContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.PredicateExpressionContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.PredicateExpressionContext", t)
	}
	return ExecFuncByType(ctx, c.Predicate())
}

var LogicalExpressionContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.LogicalExpressionContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.LogicalExpressionContext", t)
	}
	lbool, err := ExecFuncByTypeAndParseRes2Bool(ctx, c.GetLeft())
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("reslove left expression {%s} in LogicalExpression failed of {%s}",
			c.GetLeft().GetText(), err.Error()))
	}
	rbool, err := ExecFuncByTypeAndParseRes2Bool(ctx, c.GetRight())
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("reslove right expression {%s} in LogicalExpression failed of {%s}",
			c.GetRight().GetText(), err.Error()))
	}
	op := c.LogicalOperator().GetText()
	var res bool
	switch op {
	case "and", "AND", "&&":
		res = lbool && rbool
	case "or", "OR", "||":
		res = lbool || rbool
	}
	return res, nil
}

func init() {
	RegisterParseFunc("PredicateExpressionContext", PredicateExpressionContextFunc)
	RegisterParseFunc("LogicalExpressionContext", LogicalExpressionContextFunc)
}
