package execution

type SqlCtx struct {
	Status    byte
	Row   	  *Row
	Rows	  *GroupByedRows
}