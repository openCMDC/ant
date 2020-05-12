// Generated from SqlParser.g4 by ANTLR 4.7.

package parser // SqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSqlParserListener is a complete listener for a parse tree produced by SqlParser.
type BaseSqlParserListener struct{}

var _ SqlParserListener = &BaseSqlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSqlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSqlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSqlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSqlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSelectStatement is called when production selectStatement is entered.
func (s *BaseSqlParserListener) EnterSelectStatement(ctx *SelectStatementContext) {}

// ExitSelectStatement is called when production selectStatement is exited.
func (s *BaseSqlParserListener) ExitSelectStatement(ctx *SelectStatementContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseSqlParserListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseSqlParserListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterOrderByExpression is called when production orderByExpression is entered.
func (s *BaseSqlParserListener) EnterOrderByExpression(ctx *OrderByExpressionContext) {}

// ExitOrderByExpression is called when production orderByExpression is exited.
func (s *BaseSqlParserListener) ExitOrderByExpression(ctx *OrderByExpressionContext) {}

// EnterSelectElements is called when production selectElements is entered.
func (s *BaseSqlParserListener) EnterSelectElements(ctx *SelectElementsContext) {}

// ExitSelectElements is called when production selectElements is exited.
func (s *BaseSqlParserListener) ExitSelectElements(ctx *SelectElementsContext) {}

// EnterSelectStarElement is called when production selectStarElement is entered.
func (s *BaseSqlParserListener) EnterSelectStarElement(ctx *SelectStarElementContext) {}

// ExitSelectStarElement is called when production selectStarElement is exited.
func (s *BaseSqlParserListener) ExitSelectStarElement(ctx *SelectStarElementContext) {}

// EnterSelectExprElement is called when production selectExprElement is entered.
func (s *BaseSqlParserListener) EnterSelectExprElement(ctx *SelectExprElementContext) {}

// ExitSelectExprElement is called when production selectExprElement is exited.
func (s *BaseSqlParserListener) ExitSelectExprElement(ctx *SelectExprElementContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseSqlParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseSqlParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterGroupbyClause is called when production groupbyClause is entered.
func (s *BaseSqlParserListener) EnterGroupbyClause(ctx *GroupbyClauseContext) {}

// ExitGroupbyClause is called when production groupbyClause is exited.
func (s *BaseSqlParserListener) ExitGroupbyClause(ctx *GroupbyClauseContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseSqlParserListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseSqlParserListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseSqlParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseSqlParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterLimitClauseAtom is called when production limitClauseAtom is entered.
func (s *BaseSqlParserListener) EnterLimitClauseAtom(ctx *LimitClauseAtomContext) {}

// ExitLimitClauseAtom is called when production limitClauseAtom is exited.
func (s *BaseSqlParserListener) ExitLimitClauseAtom(ctx *LimitClauseAtomContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseSqlParserListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseSqlParserListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterUid is called when production uid is entered.
func (s *BaseSqlParserListener) EnterUid(ctx *UidContext) {}

// ExitUid is called when production uid is exited.
func (s *BaseSqlParserListener) ExitUid(ctx *UidContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseSqlParserListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseSqlParserListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseSqlParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseSqlParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseSqlParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseSqlParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterStringConstant is called when production stringConstant is entered.
func (s *BaseSqlParserListener) EnterStringConstant(ctx *StringConstantContext) {}

// ExitStringConstant is called when production stringConstant is exited.
func (s *BaseSqlParserListener) ExitStringConstant(ctx *StringConstantContext) {}

// EnterNumberConstant is called when production numberConstant is entered.
func (s *BaseSqlParserListener) EnterNumberConstant(ctx *NumberConstantContext) {}

// ExitNumberConstant is called when production numberConstant is exited.
func (s *BaseSqlParserListener) ExitNumberConstant(ctx *NumberConstantContext) {}

// EnterBoolConstant is called when production boolConstant is entered.
func (s *BaseSqlParserListener) EnterBoolConstant(ctx *BoolConstantContext) {}

// ExitBoolConstant is called when production boolConstant is exited.
func (s *BaseSqlParserListener) ExitBoolConstant(ctx *BoolConstantContext) {}

// EnterAggregateFunctionCall is called when production aggregateFunctionCall is entered.
func (s *BaseSqlParserListener) EnterAggregateFunctionCall(ctx *AggregateFunctionCallContext) {}

// ExitAggregateFunctionCall is called when production aggregateFunctionCall is exited.
func (s *BaseSqlParserListener) ExitAggregateFunctionCall(ctx *AggregateFunctionCallContext) {}

// EnterUdfFunctionCall is called when production udfFunctionCall is entered.
func (s *BaseSqlParserListener) EnterUdfFunctionCall(ctx *UdfFunctionCallContext) {}

// ExitUdfFunctionCall is called when production udfFunctionCall is exited.
func (s *BaseSqlParserListener) ExitUdfFunctionCall(ctx *UdfFunctionCallContext) {}

// EnterCommonAggregateFunc is called when production commonAggregateFunc is entered.
func (s *BaseSqlParserListener) EnterCommonAggregateFunc(ctx *CommonAggregateFuncContext) {}

// ExitCommonAggregateFunc is called when production commonAggregateFunc is exited.
func (s *BaseSqlParserListener) ExitCommonAggregateFunc(ctx *CommonAggregateFuncContext) {}

// EnterCountFunc1 is called when production countFunc1 is entered.
func (s *BaseSqlParserListener) EnterCountFunc1(ctx *CountFunc1Context) {}

// ExitCountFunc1 is called when production countFunc1 is exited.
func (s *BaseSqlParserListener) ExitCountFunc1(ctx *CountFunc1Context) {}

// EnterCountFunc2 is called when production countFunc2 is entered.
func (s *BaseSqlParserListener) EnterCountFunc2(ctx *CountFunc2Context) {}

// ExitCountFunc2 is called when production countFunc2 is exited.
func (s *BaseSqlParserListener) ExitCountFunc2(ctx *CountFunc2Context) {}

// EnterFunctionArgs is called when production functionArgs is entered.
func (s *BaseSqlParserListener) EnterFunctionArgs(ctx *FunctionArgsContext) {}

// ExitFunctionArgs is called when production functionArgs is exited.
func (s *BaseSqlParserListener) ExitFunctionArgs(ctx *FunctionArgsContext) {}

// EnterLogicalExpression is called when production logicalExpression is entered.
func (s *BaseSqlParserListener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production logicalExpression is exited.
func (s *BaseSqlParserListener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterPredicateExpression is called when production predicateExpression is entered.
func (s *BaseSqlParserListener) EnterPredicateExpression(ctx *PredicateExpressionContext) {}

// ExitPredicateExpression is called when production predicateExpression is exited.
func (s *BaseSqlParserListener) ExitPredicateExpression(ctx *PredicateExpressionContext) {}

// EnterInPredicate is called when production inPredicate is entered.
func (s *BaseSqlParserListener) EnterInPredicate(ctx *InPredicateContext) {}

// ExitInPredicate is called when production inPredicate is exited.
func (s *BaseSqlParserListener) ExitInPredicate(ctx *InPredicateContext) {}

// EnterIsNullPredicate is called when production isNullPredicate is entered.
func (s *BaseSqlParserListener) EnterIsNullPredicate(ctx *IsNullPredicateContext) {}

// ExitIsNullPredicate is called when production isNullPredicate is exited.
func (s *BaseSqlParserListener) ExitIsNullPredicate(ctx *IsNullPredicateContext) {}

// EnterBinaryComparasionPredicate is called when production binaryComparasionPredicate is entered.
func (s *BaseSqlParserListener) EnterBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) {
}

// ExitBinaryComparasionPredicate is called when production binaryComparasionPredicate is exited.
func (s *BaseSqlParserListener) ExitBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) {
}

// EnterBetweenPredicate is called when production betweenPredicate is entered.
func (s *BaseSqlParserListener) EnterBetweenPredicate(ctx *BetweenPredicateContext) {}

// ExitBetweenPredicate is called when production betweenPredicate is exited.
func (s *BaseSqlParserListener) ExitBetweenPredicate(ctx *BetweenPredicateContext) {}

// EnterLikePredicate is called when production likePredicate is entered.
func (s *BaseSqlParserListener) EnterLikePredicate(ctx *LikePredicateContext) {}

// ExitLikePredicate is called when production likePredicate is exited.
func (s *BaseSqlParserListener) ExitLikePredicate(ctx *LikePredicateContext) {}

// EnterRegexpPredicate is called when production regexpPredicate is entered.
func (s *BaseSqlParserListener) EnterRegexpPredicate(ctx *RegexpPredicateContext) {}

// ExitRegexpPredicate is called when production regexpPredicate is exited.
func (s *BaseSqlParserListener) ExitRegexpPredicate(ctx *RegexpPredicateContext) {}

// EnterExpressionAtomPredicate is called when production expressionAtomPredicate is entered.
func (s *BaseSqlParserListener) EnterExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) {}

// ExitExpressionAtomPredicate is called when production expressionAtomPredicate is exited.
func (s *BaseSqlParserListener) ExitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) {}

// EnterUnaryExpressionAtom is called when production unaryExpressionAtom is entered.
func (s *BaseSqlParserListener) EnterUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) {}

// ExitUnaryExpressionAtom is called when production unaryExpressionAtom is exited.
func (s *BaseSqlParserListener) ExitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) {}

// EnterColumnNameExpressionAtom is called when production columnNameExpressionAtom is entered.
func (s *BaseSqlParserListener) EnterColumnNameExpressionAtom(ctx *ColumnNameExpressionAtomContext) {}

// ExitColumnNameExpressionAtom is called when production columnNameExpressionAtom is exited.
func (s *BaseSqlParserListener) ExitColumnNameExpressionAtom(ctx *ColumnNameExpressionAtomContext) {}

// EnterConstantExpressionAtom is called when production constantExpressionAtom is entered.
func (s *BaseSqlParserListener) EnterConstantExpressionAtom(ctx *ConstantExpressionAtomContext) {}

// ExitConstantExpressionAtom is called when production constantExpressionAtom is exited.
func (s *BaseSqlParserListener) ExitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) {}

// EnterFunctionCallExpressionAtom is called when production functionCallExpressionAtom is entered.
func (s *BaseSqlParserListener) EnterFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) {
}

// ExitFunctionCallExpressionAtom is called when production functionCallExpressionAtom is exited.
func (s *BaseSqlParserListener) ExitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) {
}

// EnterBracketedExpressionAtom is called when production bracketedExpressionAtom is entered.
func (s *BaseSqlParserListener) EnterBracketedExpressionAtom(ctx *BracketedExpressionAtomContext) {}

// ExitBracketedExpressionAtom is called when production bracketedExpressionAtom is exited.
func (s *BaseSqlParserListener) ExitBracketedExpressionAtom(ctx *BracketedExpressionAtomContext) {}

// EnterMathExpressionAtom is called when production mathExpressionAtom is entered.
func (s *BaseSqlParserListener) EnterMathExpressionAtom(ctx *MathExpressionAtomContext) {}

// ExitMathExpressionAtom is called when production mathExpressionAtom is exited.
func (s *BaseSqlParserListener) ExitMathExpressionAtom(ctx *MathExpressionAtomContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseSqlParserListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseSqlParserListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseSqlParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseSqlParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterLogicalOperator is called when production logicalOperator is entered.
func (s *BaseSqlParserListener) EnterLogicalOperator(ctx *LogicalOperatorContext) {}

// ExitLogicalOperator is called when production logicalOperator is exited.
func (s *BaseSqlParserListener) ExitLogicalOperator(ctx *LogicalOperatorContext) {}

// EnterMathOperator is called when production mathOperator is entered.
func (s *BaseSqlParserListener) EnterMathOperator(ctx *MathOperatorContext) {}

// ExitMathOperator is called when production mathOperator is exited.
func (s *BaseSqlParserListener) ExitMathOperator(ctx *MathOperatorContext) {}
