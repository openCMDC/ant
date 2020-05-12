package core

import (
	"bytes"
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
	Alias    string
	DataType FieldType
}

//len(Meta.fields) must equals len(Content)
type Row struct {
	Meta    *Table
	Content []interface{}
}

type GroupByedRows struct {
	GroupKVs map[string]interface{}
	GroupByK []string
	GroupByV []interface{}
	rows     []*Row
}

func NewGroupByedRows() *GroupByedRows {
	return &GroupByedRows{
		GroupKVs: make(map[string]interface{}),
		rows:     make([]*Row, 0),
	}
}

func (g *GroupByedRows) ParseKV2String() string {
	var buffer bytes.Buffer
	for _, k := range g.GroupByK {
		buffer.WriteString(fmt.Sprintf("%s->%s", k, g.GroupKVs[k]))
	}
	return buffer.String()
}

func (g *GroupByedRows) AddKV(k string, v interface{}) {
	g.GroupKVs[k] = v
	g.GroupByK = append(g.GroupByK, k)
	g.GroupByV = append(g.GroupByV, v)
}

func (g *GroupByedRows) AddRows(r *Row) {
	g.rows = append(g.rows, r)
}

func (g *GroupByedRows) GetGroupByKeys() []string {
	return g.GroupByK
}

func (g *GroupByedRows) GetGroupByValues() []interface{} {
	return g.GroupByV
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
