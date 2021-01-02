package execution

import "ant/core"

type SqlCtx struct {
	Status byte
	Row    *core.Row
	Rows   *core.GroupByedRows
}
