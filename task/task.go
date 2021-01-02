package task

import (
	"ant/antlr3/execution"
	"ant/antlr3/parser"
	"ant/core"
	"ant/core/types"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	log "github.com/sirupsen/logrus"
	"strings"
	"sync"
	"time"
)

type TaskRuntime struct {
	types.Task
	manager      *Manager
	rowCache     []*core.Row
	ticker       *time.Ticker
	shutdownChan chan struct{}

	lock       sync.RWMutex
	selectCtx  *parser.SelectElementsContext
	whereCtx   *parser.WhereClauseContext
	groupByCtx *parser.GroupbyClauseContext
	havingCtx  *parser.HavingClauseContext
	OrderByCtx *parser.OrderByClauseContext
	limitCtx   *parser.LimitClauseContext
}

func (t *TaskRuntime) AcceptRows(rows []*core.Row) []*core.Row {
	if rows == nil || len(rows) == 0 {
		return nil
	}

	filterdRows := t.doFilter(rows)
	if len(filterdRows) == 0 {
		return nil
	}

	if t.groupByCtx == nil {
		return t.processWithoutGroupBy(filterdRows)
	} else {
		return t.processWithGroupBy(filterdRows)
	}
}

func (t *TaskRuntime) ShouldForward(row *core.Row) bool {
	//todo
	return false
}

func (t *TaskRuntime) AcceptRow(row *core.Row) {
	//todo
	if t.Status != types.Running {
		return
	}
	t.lock.Lock()
	defer t.lock.Unlock()
	t.rowCache = append(t.rowCache, row)
}

func (t *TaskRuntime) Start() {
	t.Status = types.Running
}

func (t *TaskRuntime) Stop() {
	t.Status = types.Stopped
}

func (t *TaskRuntime) commitResult() {
	//todo
	t.lock.RLock()
	defer t.lock.RUnlock()
	if len(t.rowCache) == 0 {
		//todo 上报空数据，齐全度问题
		return
	}
	result := t.AcceptRows(t.rowCache)
	if result == nil {
		//todo 上报空数据，齐全度问题
		return
	}
	t.manager.CommitTaskResult(&t.Task, result)
}

func (t *TaskRuntime) Go() {
	go func() {
		for {
			select {
			case <-t.ticker.C:
				t.commitResult()
			case <-t.shutdownChan:
				log.WithField("taskName", t.Name).Debug("stop task")
			}
		}
	}()
}

func (t *TaskRuntime) Close() {
	t.shutdownChan <- struct{}{}
}

func (t *TaskRuntime) processWithoutGroupBy(rows []*core.Row) []*core.Row {
	selectRows := make([]*core.Row, 0, 10)
	if len(rows) == 0 {
		return selectRows
	}

	srcTable := rows[0].Meta
	as := t.selectCtx.AllSelectElement()
	fields := parseSelectElements2FieldArr(as, srcTable.Fields)
	table := &core.Table{
		Name:   srcTable.Name,
		Desc:   srcTable.Desc,
		Fields: fields,
	}

	for _, r := range rows {
		newRow := &core.Row{
			Meta:    table,
			Content: []interface{}{},
		}
		sqlCtx := &execution.SqlCtx{0, r, nil}
		for _, selectElement := range as {
			switch realType := selectElement.(type) {
			case *parser.SelectStarElementContext:
				newRow.Content = append(newRow.Content, r.Content...)
			case *parser.SelectExprElementContext:
				res, err := execution.ExecFuncByType(sqlCtx, realType.ExpressionAtom())
				if err != nil {
					fmt.Errorf("processWithoutGroupBy failed of {%s}", err.Error())
					newRow.Content = append(newRow.Content, nil)
				} else {
					newRow.Content = append(newRow.Content, res)
				}
			}
		}
		selectRows = append(selectRows, newRow)
	}
	return selectRows
}

func (t *TaskRuntime) processWithGroupBy(rows []*core.Row) []*core.Row {
	//todo group by 怎么处理
	allCtx := t.groupByCtx.AllExpression()
	groupByEdRows := make(map[string]*core.GroupByedRows)
	var groupByKeys []*core.Field
	for _, row := range rows {
		gbr := core.NewGroupByedRows()
		sqlCtx := &execution.SqlCtx{0, row, nil}
		for _, expr := range allCtx {
			key := expr.GetText()
			val, err := execution.ExecFuncByType(sqlCtx, expr)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			gbr.AddKV(key, val)
		}
		if len(groupByKeys) == 0 {
			for _, name := range gbr.GroupByK {
				groupByKeys = append(groupByKeys, &core.Field{Name: name, Alias: "", DataType: core.String})
			}
		}
		if v, ok := groupByEdRows[gbr.ParseKV2String()]; ok {
			v.AddRows(row)
		} else {
			groupByEdRows[gbr.ParseKV2String()] = gbr
			gbr.AddRows(row)
		}
	}
	selectRows := make([]*core.Row, 0)
	if len(groupByEdRows) == 0 {
		return selectRows
	}

	srcTable := rows[0].Meta
	as := t.selectCtx.AllSelectElement()
	fields := parseSelectElements2FieldArr(as, groupByKeys)
	table := &core.Table{
		Name:   srcTable.Name,
		Desc:   srcTable.Desc,
		Fields: fields,
	}

	for _, v := range groupByEdRows {
		newRow := &core.Row{
			Meta:    table,
			Content: []interface{}{},
		}
		sqlCtx := &execution.SqlCtx{0, nil, v}
		for _, selectElement := range as {
			switch realType := selectElement.(type) {
			case *parser.SelectStarElementContext:
				newRow.Content = append(newRow.Content, v.GroupByV...)
			case *parser.SelectExprElementContext:
				res, err := execution.ExecFuncByType(sqlCtx, realType.ExpressionAtom())
				if err != nil {
					fmt.Errorf("processWithoutGroupBy failed of {%s}", err.Error())
					newRow.Content = append(newRow.Content, nil)
				} else {
					newRow.Content = append(newRow.Content, res)
				}
			}
		}
		selectRows = append(selectRows, newRow)
	}
	return selectRows
}

func parseSelectElements2FieldArr(selectElements []parser.ISelectElementContext, selectStarFields []*core.Field) []*core.Field {
	fields := make([]*core.Field, 0, 0)
	for _, selectElement := range selectElements {
		switch realType := selectElement.(type) {
		case *parser.SelectStarElementContext:
			fields = append(fields, selectStarFields...)
		case *parser.SelectExprElementContext:
			ea := realType.ExpressionAtom()
			var cn string
			if cat, ok := ea.(*parser.ColumnNameExpressionAtomContext); ok {
				cn = cat.ColumnName().(*parser.ColumnNameContext).GetCn().GetText()
			} else {
				cn = ea.GetText()
			}
			field := &core.Field{Name: cn, Alias: "", DataType: core.String}
			if realType.AS() != nil {
				alia := realType.Uid().GetText()
				field.Alias = alia
			}
			fields = append(fields, field)
		}
	}
	return fields
}

func (t *TaskRuntime) doFilter(rows []*core.Row) []*core.Row {
	if t.whereCtx == nil {
		return rows
	}
	filterdRows := make([]*core.Row, 0, len(rows))
	ex := t.whereCtx.Expression()
	ctx := &execution.SqlCtx{
		Status: 0,
		Row:    nil,
	}
	for _, r := range rows {
		ctx.Row = r
		res, err := execution.ExecFuncByTypeAndParseRes2Bool(ctx, ex)
		if err != nil {
			fmt.Errorf("do filter failed of {%s}", err.Error())
			continue
		}
		if res {
			filterdRows = append(filterdRows, r)
		}
	}
	return filterdRows
}

func (t *TaskRuntime) doGroupBy(rows []*core.Row) map[string]*core.GroupByedRows {
	if t.groupByCtx == nil {
		return nil
	}
	groupByExpressions := t.groupByCtx.AllExpression()
	ctx := &execution.SqlCtx{
		Status: 0,
		Row:    nil,
	}

	resMap := make(map[string]*core.GroupByedRows)

	for _, r := range rows {
		ctx.Row = r
		groupKv := core.NewGroupByedRows()
		for _, ex := range groupByExpressions {
			v, err := execution.ExecFuncByType(ctx, ex)
			if err != nil {
				fmt.Errorf("group by failed of {%s}", err.Error())
				continue
			}
			groupKv.AddKV(ex.GetText(), v)
		}
		kstring := groupKv.ParseKV2String()
		if kv, ok := resMap[kstring]; ok {
			kv.AddRows(r)
		} else {
			groupKv.AddRows(r)
			resMap[kstring] = groupKv
		}
	}
	return resMap
}

func NewTask(task types.Task, manager *Manager) (*TaskRuntime, error) {

	taskRuntime := new(TaskRuntime)
	taskRuntime.Task = task

	//todo 检测sql语句是否合法
	sql := task.SqlStr
	lastValidIndex := len(sql)
	limitIndex := strings.Index(sql, "LIMIT")
	limitClause := ""
	if limitIndex > 0 {
		limitClause = sql[limitIndex:lastValidIndex]
		lastValidIndex = limitIndex
		ast := parseAst(limitClause)
		taskRuntime.limitCtx = ast.LimitClause().(*parser.LimitClauseContext)
	}

	orderByIndex := strings.Index(sql, "ORDER BY")
	orderByClause := ""
	if orderByIndex > 0 {
		orderByClause = sql[orderByIndex:lastValidIndex]
		lastValidIndex = orderByIndex
		ast := parseAst(orderByClause)
		taskRuntime.OrderByCtx = ast.OrderByClause().(*parser.OrderByClauseContext)
	}

	havingIndex := strings.Index(sql, "HAVING")
	havingClause := ""
	if havingIndex > 0 {
		havingClause = sql[havingIndex:lastValidIndex]
		lastValidIndex = havingIndex
		ast := parseAst(havingClause)
		taskRuntime.havingCtx = ast.HavingClause().(*parser.HavingClauseContext)
	}

	groupByIndex := strings.Index(sql, "GROUP BY")
	groupByClause := ""
	if groupByIndex > 0 {
		groupByClause = sql[groupByIndex:lastValidIndex]
		lastValidIndex = groupByIndex
		ast := parseAst(groupByClause)
		taskRuntime.groupByCtx = ast.GroupbyClause().(*parser.GroupbyClauseContext)
	}

	whereIndex := strings.Index(sql, "WHERE")
	whereClause := ""
	if whereIndex > 0 {
		whereClause = sql[whereIndex:lastValidIndex]
		lastValidIndex = whereIndex
		ast := parseAst(whereClause)
		taskRuntime.whereCtx = ast.WhereClause().(*parser.WhereClauseContext)
	}

	selectClause := sql[0:lastValidIndex]
	if len(selectClause) > 0 {
		ast := parseAst(selectClause)
		statmentCtx := ast.SelectStatement().(*parser.SelectStatementContext)
		taskRuntime.selectCtx = statmentCtx.SelectElements().(*parser.SelectElementsContext)
		//taskRuntime.Table = statmentCtx.GetTableName().GetText()
	}

	taskRuntime.rowCache = make([]*core.Row, 0, 128)
	taskRuntime.ticker = time.NewTicker(time.Duration(task.Interval) * time.Second)
	taskRuntime.shutdownChan = make(chan struct{})
	taskRuntime.manager = manager
	return taskRuntime, nil
}

func parseAst(str string) *parser.SqlParser {
	is := antlr.NewInputStream(str)
	lexer := parser.NewSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSqlParser(stream)
	return p
}
