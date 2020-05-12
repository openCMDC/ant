// Generated from SqlParser.g4 by ANTLR 4.7.

package parser // SqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SqlParserListener is a complete listener for a parse tree produced by SqlParser.
type SqlParserListener interface {
	antlr.ParseTreeListener

	// EnterSelectStatement is called when entering the selectStatement production.
	EnterSelectStatement(c *SelectStatementContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterOrderByExpression is called when entering the orderByExpression production.
	EnterOrderByExpression(c *OrderByExpressionContext)

	// EnterSelectElements is called when entering the selectElements production.
	EnterSelectElements(c *SelectElementsContext)

	// EnterSelectStarElement is called when entering the selectStarElement production.
	EnterSelectStarElement(c *SelectStarElementContext)

	// EnterSelectExprElement is called when entering the selectExprElement production.
	EnterSelectExprElement(c *SelectExprElementContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterGroupbyClause is called when entering the groupbyClause production.
	EnterGroupbyClause(c *GroupbyClauseContext)

	// EnterHavingClause is called when entering the havingClause production.
	EnterHavingClause(c *HavingClauseContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterLimitClauseAtom is called when entering the limitClauseAtom production.
	EnterLimitClauseAtom(c *LimitClauseAtomContext)

	// EnterColumnName is called when entering the columnName production.
	EnterColumnName(c *ColumnNameContext)

	// EnterUid is called when entering the uid production.
	EnterUid(c *UidContext)

	// EnterDecimalLiteral is called when entering the decimalLiteral production.
	EnterDecimalLiteral(c *DecimalLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterStringConstant is called when entering the stringConstant production.
	EnterStringConstant(c *StringConstantContext)

	// EnterNumberConstant is called when entering the numberConstant production.
	EnterNumberConstant(c *NumberConstantContext)

	// EnterBoolConstant is called when entering the boolConstant production.
	EnterBoolConstant(c *BoolConstantContext)

	// EnterAggregateFunctionCall is called when entering the aggregateFunctionCall production.
	EnterAggregateFunctionCall(c *AggregateFunctionCallContext)

	// EnterUdfFunctionCall is called when entering the udfFunctionCall production.
	EnterUdfFunctionCall(c *UdfFunctionCallContext)

	// EnterCommonAggregateFunc is called when entering the commonAggregateFunc production.
	EnterCommonAggregateFunc(c *CommonAggregateFuncContext)

	// EnterCountFunc1 is called when entering the countFunc1 production.
	EnterCountFunc1(c *CountFunc1Context)

	// EnterCountFunc2 is called when entering the countFunc2 production.
	EnterCountFunc2(c *CountFunc2Context)

	// EnterFunctionArgs is called when entering the functionArgs production.
	EnterFunctionArgs(c *FunctionArgsContext)

	// EnterLogicalExpression is called when entering the logicalExpression production.
	EnterLogicalExpression(c *LogicalExpressionContext)

	// EnterPredicateExpression is called when entering the predicateExpression production.
	EnterPredicateExpression(c *PredicateExpressionContext)

	// EnterInPredicate is called when entering the inPredicate production.
	EnterInPredicate(c *InPredicateContext)

	// EnterIsNullPredicate is called when entering the isNullPredicate production.
	EnterIsNullPredicate(c *IsNullPredicateContext)

	// EnterBinaryComparasionPredicate is called when entering the binaryComparasionPredicate production.
	EnterBinaryComparasionPredicate(c *BinaryComparasionPredicateContext)

	// EnterBetweenPredicate is called when entering the betweenPredicate production.
	EnterBetweenPredicate(c *BetweenPredicateContext)

	// EnterLikePredicate is called when entering the likePredicate production.
	EnterLikePredicate(c *LikePredicateContext)

	// EnterRegexpPredicate is called when entering the regexpPredicate production.
	EnterRegexpPredicate(c *RegexpPredicateContext)

	// EnterExpressionAtomPredicate is called when entering the expressionAtomPredicate production.
	EnterExpressionAtomPredicate(c *ExpressionAtomPredicateContext)

	// EnterUnaryExpressionAtom is called when entering the unaryExpressionAtom production.
	EnterUnaryExpressionAtom(c *UnaryExpressionAtomContext)

	// EnterColumnNameExpressionAtom is called when entering the columnNameExpressionAtom production.
	EnterColumnNameExpressionAtom(c *ColumnNameExpressionAtomContext)

	// EnterConstantExpressionAtom is called when entering the constantExpressionAtom production.
	EnterConstantExpressionAtom(c *ConstantExpressionAtomContext)

	// EnterFunctionCallExpressionAtom is called when entering the functionCallExpressionAtom production.
	EnterFunctionCallExpressionAtom(c *FunctionCallExpressionAtomContext)

	// EnterBracketedExpressionAtom is called when entering the bracketedExpressionAtom production.
	EnterBracketedExpressionAtom(c *BracketedExpressionAtomContext)

	// EnterMathExpressionAtom is called when entering the mathExpressionAtom production.
	EnterMathExpressionAtom(c *MathExpressionAtomContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterLogicalOperator is called when entering the logicalOperator production.
	EnterLogicalOperator(c *LogicalOperatorContext)

	// EnterMathOperator is called when entering the mathOperator production.
	EnterMathOperator(c *MathOperatorContext)

	// ExitSelectStatement is called when exiting the selectStatement production.
	ExitSelectStatement(c *SelectStatementContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitOrderByExpression is called when exiting the orderByExpression production.
	ExitOrderByExpression(c *OrderByExpressionContext)

	// ExitSelectElements is called when exiting the selectElements production.
	ExitSelectElements(c *SelectElementsContext)

	// ExitSelectStarElement is called when exiting the selectStarElement production.
	ExitSelectStarElement(c *SelectStarElementContext)

	// ExitSelectExprElement is called when exiting the selectExprElement production.
	ExitSelectExprElement(c *SelectExprElementContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitGroupbyClause is called when exiting the groupbyClause production.
	ExitGroupbyClause(c *GroupbyClauseContext)

	// ExitHavingClause is called when exiting the havingClause production.
	ExitHavingClause(c *HavingClauseContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitLimitClauseAtom is called when exiting the limitClauseAtom production.
	ExitLimitClauseAtom(c *LimitClauseAtomContext)

	// ExitColumnName is called when exiting the columnName production.
	ExitColumnName(c *ColumnNameContext)

	// ExitUid is called when exiting the uid production.
	ExitUid(c *UidContext)

	// ExitDecimalLiteral is called when exiting the decimalLiteral production.
	ExitDecimalLiteral(c *DecimalLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitStringConstant is called when exiting the stringConstant production.
	ExitStringConstant(c *StringConstantContext)

	// ExitNumberConstant is called when exiting the numberConstant production.
	ExitNumberConstant(c *NumberConstantContext)

	// ExitBoolConstant is called when exiting the boolConstant production.
	ExitBoolConstant(c *BoolConstantContext)

	// ExitAggregateFunctionCall is called when exiting the aggregateFunctionCall production.
	ExitAggregateFunctionCall(c *AggregateFunctionCallContext)

	// ExitUdfFunctionCall is called when exiting the udfFunctionCall production.
	ExitUdfFunctionCall(c *UdfFunctionCallContext)

	// ExitCommonAggregateFunc is called when exiting the commonAggregateFunc production.
	ExitCommonAggregateFunc(c *CommonAggregateFuncContext)

	// ExitCountFunc1 is called when exiting the countFunc1 production.
	ExitCountFunc1(c *CountFunc1Context)

	// ExitCountFunc2 is called when exiting the countFunc2 production.
	ExitCountFunc2(c *CountFunc2Context)

	// ExitFunctionArgs is called when exiting the functionArgs production.
	ExitFunctionArgs(c *FunctionArgsContext)

	// ExitLogicalExpression is called when exiting the logicalExpression production.
	ExitLogicalExpression(c *LogicalExpressionContext)

	// ExitPredicateExpression is called when exiting the predicateExpression production.
	ExitPredicateExpression(c *PredicateExpressionContext)

	// ExitInPredicate is called when exiting the inPredicate production.
	ExitInPredicate(c *InPredicateContext)

	// ExitIsNullPredicate is called when exiting the isNullPredicate production.
	ExitIsNullPredicate(c *IsNullPredicateContext)

	// ExitBinaryComparasionPredicate is called when exiting the binaryComparasionPredicate production.
	ExitBinaryComparasionPredicate(c *BinaryComparasionPredicateContext)

	// ExitBetweenPredicate is called when exiting the betweenPredicate production.
	ExitBetweenPredicate(c *BetweenPredicateContext)

	// ExitLikePredicate is called when exiting the likePredicate production.
	ExitLikePredicate(c *LikePredicateContext)

	// ExitRegexpPredicate is called when exiting the regexpPredicate production.
	ExitRegexpPredicate(c *RegexpPredicateContext)

	// ExitExpressionAtomPredicate is called when exiting the expressionAtomPredicate production.
	ExitExpressionAtomPredicate(c *ExpressionAtomPredicateContext)

	// ExitUnaryExpressionAtom is called when exiting the unaryExpressionAtom production.
	ExitUnaryExpressionAtom(c *UnaryExpressionAtomContext)

	// ExitColumnNameExpressionAtom is called when exiting the columnNameExpressionAtom production.
	ExitColumnNameExpressionAtom(c *ColumnNameExpressionAtomContext)

	// ExitConstantExpressionAtom is called when exiting the constantExpressionAtom production.
	ExitConstantExpressionAtom(c *ConstantExpressionAtomContext)

	// ExitFunctionCallExpressionAtom is called when exiting the functionCallExpressionAtom production.
	ExitFunctionCallExpressionAtom(c *FunctionCallExpressionAtomContext)

	// ExitBracketedExpressionAtom is called when exiting the bracketedExpressionAtom production.
	ExitBracketedExpressionAtom(c *BracketedExpressionAtomContext)

	// ExitMathExpressionAtom is called when exiting the mathExpressionAtom production.
	ExitMathExpressionAtom(c *MathExpressionAtomContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitLogicalOperator is called when exiting the logicalOperator production.
	ExitLogicalOperator(c *LogicalOperatorContext)

	// ExitMathOperator is called when exiting the mathOperator production.
	ExitMathOperator(c *MathOperatorContext)
}
