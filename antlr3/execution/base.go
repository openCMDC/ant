package execution

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
	"strconv"
	"strings"
)

type execFunc func(ctx *SqlCtx, t antlr.Tree) (interface{}, error)

var expressionMap = make(map[string]execFunc)

func RegisterParseFunc(name string, function execFunc) {
	expressionMap[name] = function
}

func ExecFuncByType(ctx *SqlCtx, t antlr.Tree) (interface{}, error) {
	typeString := reflect.TypeOf(t).String()
	lastDot := strings.LastIndex(typeString, `.`)
	shortType := typeString[lastDot+1:]
	f, ok := expressionMap[shortType]
	if !ok {
		type yummyInterface interface {
			GetText() string
		}
		y, ok := t.(yummyInterface)
		if ok {
			return nil, NewResolveFuncNotFoundErr(y.GetText())
		} else {
			return nil, NewResolveFuncNotFoundErr(t)
		}

	}
	return f(ctx, t)
}

func ExecFuncByTypeAndParseRes2Float(ctx *SqlCtx, t antlr.Tree) (float64, error) {
	res, err := ExecFuncByType(ctx, t)
	if err != nil {
		return -1, NewAntlrErr(fmt.Sprintf("resolve string {%s} failed of {%s}", t, err.Error()))
	}
	val, err := parseVal2Float64(res)
	if err != nil {
		return -1, NewAntlrErr(fmt.Sprintf("parse {%s} to float failed of {%s}",
			res, err.Error()))
	}
	return val, nil
}

func ExecFuncByTypeAndParseRes2String(ctx *SqlCtx, t antlr.Tree) (string, error) {
	res, err := ExecFuncByType(ctx, t)
	if err != nil {
		return "", NewAntlrErr(fmt.Sprintf("resolve string {%s} failed of {%s}", t, err.Error()))
	}
	val, ok := res.(string)
	if !ok {
		return "", NewAntlrErr(fmt.Sprintf("parse {%s} to string failed", res))
	}
	return val, nil
}

func ExecFuncByTypeAndParseRes2Bool(ctx *SqlCtx, t antlr.Tree) (bool, error) {
	res, err := ExecFuncByType(ctx, t)
	if err != nil {
		return false, NewAntlrErr(fmt.Sprintf("resolve string {%s} failed of {%s}", t, err.Error()))
	}
	val, ok := res.(bool)
	if !ok {
		return false, NewAntlrErr(fmt.Sprintf("parse {%s} to bool failed", res))
	}
	return val, nil
}

func parseVal2Float64(source interface{}) (float64, error) {
	var res float64
	var err error = nil
	switch v := source.(type) {
	case int8:
		res = float64(v)
	case uint8:
		res = float64(v)
	case int16:
		res = float64(v)
	case uint16:
		res = float64(v)
	case int32:
		res = float64(v)
	case uint32:
		res = float64(v)
	case int64:
		res = float64(v)
	case uint64:
		res = float64(v)
	case int:
		res = float64(v)
	case uint:
		res = float64(v)
	case float32:
		res = float64(v)
	case float64:
		res = v
	case string:
		res, err = strconv.ParseFloat(v, 64)
	default:
		err = NewAntlrErr(fmt.Sprintf("can't parse {%s} to float", source))
	}
	return res, err
}

func parseVal2Int64(source interface{}) (int64, error) {
	var res int64
	var err error = nil
	switch v := source.(type) {
	case int8:
		res = int64(v)
	case uint8:
		res = int64(v)
	case int16:
		res = int64(v)
	case uint16:
		res = int64(v)
	case int32:
		res = int64(v)
	case uint32:
		res = int64(v)
	case int64:
		res = int64(v)
	case uint64:
		res = int64(v)
	case int:
		res = int64(v)
	case uint:
		res = int64(v)
	case float32:
		res = int64(v)
	case float64:
		res = int64(v)
	case string:
		res, err = strconv.ParseInt(v, 10, 64)
	default:
		err = NewAntlrErr(fmt.Sprintf("can't parse {%s} to float", source))
	}
	return res, err
}
