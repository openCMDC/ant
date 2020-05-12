package execution

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type AntlrErr struct {
	msg string
}

func (a *AntlrErr) Error() string {
	return a.msg
}

func NewAntlrErr(msg string) *AntlrErr{
	return &AntlrErr{msg}
}

func NewTypeNoMatchedErr(requireType string, actuallyType antlr.Tree) *AntlrErr {
	return NewAntlrErr(fmt.Sprintf("type mismatch, required {%s}, actually {%T}", requireType, actuallyType))
}

func NewResolveFuncNotFoundErr(toBeResolvedString interface{}) *AntlrErr{
	return NewAntlrErr(fmt.Sprintf("can not found proper resolve function for string {%s}", toBeResolvedString))
}

func NewColumnNotExistedErr(columnName string) *AntlrErr{
	return NewAntlrErr(fmt.Sprintf("column {%s} not existed", columnName))
}

