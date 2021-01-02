package execution

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"regexp"
	"strings"
)
import "ant/antlr3/parser"

var InPredicateContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.InPredicateContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.InPredicateContext", t)
	}
	not := c.NOT()
	//c.Predicate()
	leftRes, err := ExecFuncByType(ctx, c.GetLeft())
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("resolve left expression {%s} in InPredicate failed of %s",
			c.GetLeft().GetText(), err.Error()))
	}

	args := make([]interface{}, 0, 0)
	for _, ex := range c.AllExpressionAtom() {
		r, err := ExecFuncByType(ctx, ex)
		if err != nil {
			return nil, NewAntlrErr(fmt.Sprintf("compute function arg {%s} faild of {%s}", ex.GetText(), err.Error()))
		}
		args = append(args, r)
	}

	if not == nil {
		for _, item := range args {
			if item == leftRes {
				return true, nil
			}
		}
		return false, nil
	} else {
		for _, item := range args {
			if item == leftRes {
				return false, nil
			}
		}
		return true, nil
	}
}

var IsNullPredicateContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.IsNullPredicateContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.IsNullPredicateContext", t)
	}
	res, err := ExecFuncByType(ctx, c.ExpressionAtom())
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("resolve left expression {%s} in IsNullPredicate failed of %s",
			c.ExpressionAtom().GetText(), err.Error()))
	}
	if c.NOT() != nil {
		return res != nil, nil
	} else {
		return res == nil, nil
	}
}

var BinaryComparasionPredicateContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.BinaryComparasionPredicateContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.BinaryComparasionPredicateContext", t)
	}
	//todo
	lf, err := ExecFuncByTypeAndParseRes2String(ctx, c.GetLeft())
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("resolve left expression {%s} in BinaryComparasionPredicate failed of {%s}",
			c.GetLeft().GetText(), err.Error()))
	}

	rf, err := ExecFuncByTypeAndParseRes2String(ctx, c.GetRight())
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("resolve right expression {%s} in BinaryComparasionPredicate failed of %s",
			c.GetRight().GetText(), err.Error()))
	}

	op := c.ComparisonOperator().GetText()
	var res bool
	switch op {
	case "=":
		res = lf == rf
	case ">":
		res = lf > rf
	case "<":
		res = lf < rf
	case ">=":
		res = lf >= rf
	case "<=":
		res = lf <= rf
	case "!=":
		res = lf != rf
	}
	return res, nil
}

var BetweenPredicateContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.BetweenPredicateContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.BetweenPredicateContext", t)
	}
	lf, err := ExecFuncByTypeAndParseRes2Float(ctx, c.GetLeft())
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("resolve left expression {%s} in BetweenPredicate failed of %s",
			c.GetLeft().GetText(), err.Error()))
	}

	ff, err := ExecFuncByTypeAndParseRes2Float(ctx, c.GetFrom())
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("resolve right expression {%s} in BetweenPredicate failed of %s",
			c.GetFrom().GetText(), err.Error()))
	}

	tf, err := ExecFuncByTypeAndParseRes2Float(ctx, c.GetTo())
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("resolve right expression {%s} in BetweenPredicate failed of %s",
			c.GetTo().GetText(), err.Error()))
	}

	var res bool
	if lf >= ff && lf < tf {
		res = true
	}
	if c.NOT() != nil {
		res = !res
	}
	return res, nil
}

var LikePredicateContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.LikePredicateContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.LikePredicateContext", t)
	}
	leftStr, err := ExecFuncByTypeAndParseRes2String(ctx, c.GetLeft())
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("resolve left expression {%s} in LikePredicate failed of %s",
			c.GetLeft().GetText(), err.Error()))
	}
	likeStr, err := ExecFuncByTypeAndParseRes2String(ctx, c.GetLikeStr())
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("resolve likeStr expression {%s} in LikePredicate failed of %s",
			c.GetLeft().GetText(), err.Error()))
	}

	s := strings.ReplaceAll(likeStr, `%`, `.*`)
	s = strings.ReplaceAll(likeStr, `_`, `.*`)
	match, err := regexp.MatchString(s, leftStr)
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("regexp match failed for {%s} and {%s}", leftStr, s))
	}
	if c.NOT() != nil {
		match = !match
	}
	return match, nil
}

var RegexpPredicateContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.RegexpPredicateContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.RegexpPredicateContext", t)
	}
	leftStr, err := ExecFuncByTypeAndParseRes2String(ctx, c.GetLeft())
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("resolve left expression {%s} in LikePredicate failed of %s",
			c.GetLeft().GetText(), err.Error()))
	}
	regStr, err := ExecFuncByTypeAndParseRes2String(ctx, c.GetRegStr())
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("resolve regStr expression {%s} in LikePredicate failed of %s",
			c.GetLeft().GetText(), err.Error()))
	}

	match, err := regexp.MatchString(regStr, leftStr)
	if err != nil {
		return nil, NewAntlrErr(fmt.Sprintf("regexp match failed for {%s} and {%s}", leftStr, regStr))
	}
	if c.NOT() != nil {
		match = !match
	}
	return match, nil
}

var ExpressionAtomPredicateContextFunc = func(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	c, e := t.(*parser.ExpressionAtomPredicateContext)
	if !e {
		return nil, NewTypeNoMatchedErr("parser.ExpressionAtomPredicateContext", t)
	}
	c.GetText()
	return ExecFuncByType(ctx, c.ExpressionAtom())
}

func init() {
	RegisterParseFunc("InPredicateContext", InPredicateContextFunc)
	RegisterParseFunc("IsNullPredicateContext", IsNullPredicateContextFunc)
	RegisterParseFunc("BinaryComparasionPredicateContext", BinaryComparasionPredicateContextFunc)
	RegisterParseFunc("BetweenPredicateContext", BetweenPredicateContextFunc)
	RegisterParseFunc("LikePredicateContext", LikePredicateContextFunc)
	RegisterParseFunc("RegexpPredicateContext", RegexpPredicateContextFunc)
	RegisterParseFunc("ExpressionAtomPredicateContext", ExpressionAtomPredicateContextFunc)
}
