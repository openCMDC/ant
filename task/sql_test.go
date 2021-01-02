package task

import (
	"ant/core"
	"ant/core/types"
	"fmt"
	"testing"
)

func TestSql(test *testing.T) {
	//is := antlr.NewInputStream(`SELECT func1(c1), c2 FROM http WHERE (c2 + 3) / 2`)
	//sql := `SELECT func1(c1), c2 FROM http WHERE c2 / (3 + 2) ORDER BY c3`
	//
	//lastValidIndex := len(sql)
	//
	//limitIndex := strings.Index(sql, "LIMIT")
	//limitClause := ""
	//if limitIndex > 0 {
	//	limitClause = sql[limitIndex:lastValidIndex]
	//	lastValidIndex = limitIndex
	//}
	//
	//orderByIndex := strings.Index(sql, "ORDER BY")
	//orderByClause := ""
	//if orderByIndex > 0 {
	//	orderByClause = sql[orderByIndex:lastValidIndex]
	//	lastValidIndex = orderByIndex
	//}
	//
	//havingIndex := strings.Index(sql, "HAVING")
	//havingClause := ""
	//if havingIndex > 0 {
	//	havingClause = sql[havingIndex:lastValidIndex]
	//	lastValidIndex = havingIndex
	//}
	//
	//whereIndex := strings.Index(sql, "WHERE")
	//whereClause := ""
	//if whereIndex > 0 {
	//	whereClause = sql[whereIndex:lastValidIndex]
	//	lastValidIndex = whereIndex
	//}
	//
	//selectClause := sql[0:lastValidIndex]
	//
	//if len(selectClause) > 0 {
	//	ast := parseAst(selectClause)
	//	antlr.ParseTreeWalkerDefault.Walk(&parser.BaseSqlParserListener{}, ast.SelectStatement())
	//}
	//
	//fmt.Println("-----------------------------------------------------------")
	//
	//if len(whereClause) > 0 {
	//	ast := parseAst(whereClause)
	//	antlr.ParseTreeWalkerDefault.Walk(&parser.BaseSqlParserListener{}, ast.WhereClause())
	//	//execution.ExecFuncByTypeAndParseRes2Bool(nil, ast.WhereClause())
	//}
	//
	//fmt.Println("-----------------------------------------------------------")
	//
	//if len(havingClause) > 0 {
	//	ast := parseAst(havingClause)
	//	antlr.ParseTreeWalkerDefault.Walk(&parser.BaseSqlParserListener{}, ast.HavingClause())
	//}
	//
	//fmt.Println("-----------------------------------------------------------")
	//
	//if len(orderByClause) > 0 {
	//	ast := parseAst(orderByClause)
	//	antlr.ParseTreeWalkerDefault.Walk(&parser.BaseSqlParserListener{}, ast.OrderByClause())
	//}
	//
	//fmt.Println("-----------------------------------------------------------")
	//
	//if len(limitClause) > 0 {
	//	ast := parseAst(limitClause)
	//	antlr.ParseTreeWalkerDefault.Walk(&parser.BaseSqlParserListener{}, ast.OrderByClause())
	//}
	fields := make([]*core.Field, 0, 0)
	fields = append(fields, &core.Field{Name: "id", Alias: "", DataType: core.Number})
	fields = append(fields, &core.Field{Name: "name", Alias: "", DataType: core.String})
	fields = append(fields, &core.Field{Name: "age", Alias: "", DataType: core.Number})
	t := &core.Table{
		Name:   "test",
		Desc:   "test",
		Fields: fields,
	}
	rows := make([]*core.Row, 0, 10)
	rows = append(rows, &core.Row{
		Meta:    t,
		Content: []interface{}{1, "wang", 23},
	})
	rows = append(rows, &core.Row{
		Meta:    t,
		Content: []interface{}{2, "li", 19},
	})
	rows = append(rows, &core.Row{
		Meta:    t,
		Content: []interface{}{3, "wang", 70},
	})

	task, err := NewTask(types.Task{}, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res := task.AcceptRows(rows)
	fmt.Println(res)
	//sqlctx := &execution.SqlCtx{0, r}
	//
	//ast := parseAst(`id`)
	//fmt.Println(ast)
	//res, err := execution.ExecFuncByType(sqlctx, ast.WhereClause().(*parser.WhereClauseContext))
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Println(res)
	//}
	//antlr.ParseTreeWalkerDefault.Walk(&parser.Mylistener{}, ast.Pre())
	//antlr.ParseTreeWalkerDefault.Walk(&parser.Mylistener{}, ast.fu())

	// Create the Parser
	//for {
	//	t := lexer.NextToken()
	//	if t.GetTokenType() == antlr.TokenEOF {
	//		break
	//	}
	//	fmt.Printf("%s (%q)\n",
	//		lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	//}
	//p := parser.NewSqlParser(stream)
	//p1 := parser.NewSqlParser(stream)
	//antlr.ParseTreeWalkerDefault.Walk(&parser.BaseSqlParserListener{}, p.WhereClause())
	//antlr.ParseTreeWalkerDefault.Walk(&parser.BaseSqlParserListener{}, p.SelectStatement())
	//antlr.ParseTreeWalkerDefault.Walk(&parser.BaseSqlParserListener{}, p1.SelectStatement())

	// Finally parse the expression
	//var myvisitor parser.SqlParserVisitor = &parser.BaseSqlParserVisitor{}
	////antlr.ParseTreeWalkerDefault.(&myvisitor, p.SelectStatement())
	//selectStatment := p.SelectStatement()
	//selectStatment.Accept(myvisitor)
	//eles := p.SelectElements()
	//orderBy := p.OrderByClause()
	//myvisitor.Visit(selectStatment)
	//myvisitor.Visit(eles)
	//myvisitor.Visit(orderBy)
}
