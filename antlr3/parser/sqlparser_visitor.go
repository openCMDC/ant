// Generated from SqlParser.g4 by ANTLR 4.7.

package parser // SqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SqlParser.
type SqlParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SqlParser#selectStatement.
	VisitSelectStatement(ctx *SelectStatementContext) interface{}

	// Visit a parse tree produced by SqlParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by SqlParser#orderByExpression.
	VisitOrderByExpression(ctx *OrderByExpressionContext) interface{}

	// Visit a parse tree produced by SqlParser#selectElements.
	VisitSelectElements(ctx *SelectElementsContext) interface{}

	// Visit a parse tree produced by SqlParser#selectStarElement.
	VisitSelectStarElement(ctx *SelectStarElementContext) interface{}

	// Visit a parse tree produced by SqlParser#selectExprElement.
	VisitSelectExprElement(ctx *SelectExprElementContext) interface{}

	// Visit a parse tree produced by SqlParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by SqlParser#groupbyClause.
	VisitGroupbyClause(ctx *GroupbyClauseContext) interface{}

	// Visit a parse tree produced by SqlParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by SqlParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by SqlParser#limitClauseAtom.
	VisitLimitClauseAtom(ctx *LimitClauseAtomContext) interface{}

	// Visit a parse tree produced by SqlParser#columnName.
	VisitColumnName(ctx *ColumnNameContext) interface{}

	// Visit a parse tree produced by SqlParser#uid.
	VisitUid(ctx *UidContext) interface{}

	// Visit a parse tree produced by SqlParser#decimalLiteral.
	VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{}

	// Visit a parse tree produced by SqlParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by SqlParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by SqlParser#stringConstant.
	VisitStringConstant(ctx *StringConstantContext) interface{}

	// Visit a parse tree produced by SqlParser#numberConstant.
	VisitNumberConstant(ctx *NumberConstantContext) interface{}

	// Visit a parse tree produced by SqlParser#boolConstant.
	VisitBoolConstant(ctx *BoolConstantContext) interface{}

	// Visit a parse tree produced by SqlParser#aggregateFunctionCall.
	VisitAggregateFunctionCall(ctx *AggregateFunctionCallContext) interface{}

	// Visit a parse tree produced by SqlParser#udfFunctionCall.
	VisitUdfFunctionCall(ctx *UdfFunctionCallContext) interface{}

	// Visit a parse tree produced by SqlParser#commonAggregateFunc.
	VisitCommonAggregateFunc(ctx *CommonAggregateFuncContext) interface{}

	// Visit a parse tree produced by SqlParser#countFunc1.
	VisitCountFunc1(ctx *CountFunc1Context) interface{}

	// Visit a parse tree produced by SqlParser#countFunc2.
	VisitCountFunc2(ctx *CountFunc2Context) interface{}

	// Visit a parse tree produced by SqlParser#functionArgs.
	VisitFunctionArgs(ctx *FunctionArgsContext) interface{}

	// Visit a parse tree produced by SqlParser#logicalExpression.
	VisitLogicalExpression(ctx *LogicalExpressionContext) interface{}

	// Visit a parse tree produced by SqlParser#predicateExpression.
	VisitPredicateExpression(ctx *PredicateExpressionContext) interface{}

	// Visit a parse tree produced by SqlParser#inPredicate.
	VisitInPredicate(ctx *InPredicateContext) interface{}

	// Visit a parse tree produced by SqlParser#isNullPredicate.
	VisitIsNullPredicate(ctx *IsNullPredicateContext) interface{}

	// Visit a parse tree produced by SqlParser#binaryComparasionPredicate.
	VisitBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) interface{}

	// Visit a parse tree produced by SqlParser#betweenPredicate.
	VisitBetweenPredicate(ctx *BetweenPredicateContext) interface{}

	// Visit a parse tree produced by SqlParser#likePredicate.
	VisitLikePredicate(ctx *LikePredicateContext) interface{}

	// Visit a parse tree produced by SqlParser#regexpPredicate.
	VisitRegexpPredicate(ctx *RegexpPredicateContext) interface{}

	// Visit a parse tree produced by SqlParser#expressionAtomPredicate.
	VisitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) interface{}

	// Visit a parse tree produced by SqlParser#unaryExpressionAtom.
	VisitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) interface{}

	// Visit a parse tree produced by SqlParser#columnNameExpressionAtom.
	VisitColumnNameExpressionAtom(ctx *ColumnNameExpressionAtomContext) interface{}

	// Visit a parse tree produced by SqlParser#constantExpressionAtom.
	VisitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) interface{}

	// Visit a parse tree produced by SqlParser#functionCallExpressionAtom.
	VisitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) interface{}

	// Visit a parse tree produced by SqlParser#bracketedExpressionAtom.
	VisitBracketedExpressionAtom(ctx *BracketedExpressionAtomContext) interface{}

	// Visit a parse tree produced by SqlParser#mathExpressionAtom.
	VisitMathExpressionAtom(ctx *MathExpressionAtomContext) interface{}

	// Visit a parse tree produced by SqlParser#unaryOperator.
	VisitUnaryOperator(ctx *UnaryOperatorContext) interface{}

	// Visit a parse tree produced by SqlParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by SqlParser#logicalOperator.
	VisitLogicalOperator(ctx *LogicalOperatorContext) interface{}

	// Visit a parse tree produced by SqlParser#mathOperator.
	VisitMathOperator(ctx *MathOperatorContext) interface{}
}
