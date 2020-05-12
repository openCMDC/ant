// Generated from SqlParser.g4 by ANTLR 4.7.

package parser // SqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSqlParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSqlParserVisitor) VisitSelectStatement(ctx *SelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitOrderByExpression(ctx *OrderByExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitSelectElements(ctx *SelectElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitSelectStarElement(ctx *SelectStarElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitSelectExprElement(ctx *SelectExprElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitGroupbyClause(ctx *GroupbyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitLimitClauseAtom(ctx *LimitClauseAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitColumnName(ctx *ColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitUid(ctx *UidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitStringConstant(ctx *StringConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitNumberConstant(ctx *NumberConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitBoolConstant(ctx *BoolConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitAggregateFunctionCall(ctx *AggregateFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitUdfFunctionCall(ctx *UdfFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitCommonAggregateFunc(ctx *CommonAggregateFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitCountFunc1(ctx *CountFunc1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitCountFunc2(ctx *CountFunc2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitFunctionArgs(ctx *FunctionArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitLogicalExpression(ctx *LogicalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitPredicateExpression(ctx *PredicateExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitInPredicate(ctx *InPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitIsNullPredicate(ctx *IsNullPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitBetweenPredicate(ctx *BetweenPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitLikePredicate(ctx *LikePredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitRegexpPredicate(ctx *RegexpPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitColumnNameExpressionAtom(ctx *ColumnNameExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitBracketedExpressionAtom(ctx *BracketedExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitMathExpressionAtom(ctx *MathExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitUnaryOperator(ctx *UnaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitLogicalOperator(ctx *LogicalOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitMathOperator(ctx *MathOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}
