package core

type FieldType uint8

type Table struct {
	Name   string
	Desc   string
	Fields []*Field
}

type Field struct {
	Name     string
	DataType FieldType
}

type Row interface {
	Meta() *Table
	Content() []interface{}
}
