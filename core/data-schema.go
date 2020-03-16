package core

import (
	"fmt"
	"strings"
)

type FieldType uint8

const Number FieldType = 0
const String FieldType = 1
const Map FieldType = 2

type Table struct {
	Name   string
	Desc   string
	Fields []*Field
}
type Field struct {
	Name     string
	DataType FieldType
}

//len(Meta.fields) must equals len(Content)
type Row struct {
	Meta    *Table
	Content []interface{}
}

func (r *Row) String() string {
	var str strings.Builder
	if len(r.Meta.Fields) != len(r.Content) {
		str.WriteString("warning!!! length of table fields doesn't equals to real content length \n")
	}
	var tableHeader []string
	var tableBody []string
	for i := 0; i < len(r.Meta.Fields) && i < len(r.Content); i++ {
		filed := r.Meta.Fields[i]
		tableHeader = append(tableHeader, filed.Name)
		content := r.Content[i]
		tableBody = append(tableBody, parseToShortStr(content))
	}
	str.WriteString(strings.Join(tableHeader, "|\t"))
	str.WriteString("\n")
	str.WriteString(strings.Join(tableBody, "|\t"))
	return str.String()
}
func parseToShortStr(src interface{}) string {
	str := fmt.Sprintf("%s", src)
	if len(str) <= 20 {
		return str
	}
	return str[0:20]
}
