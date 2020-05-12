/*
MySQL (Positive Technologies) grammar
The MIT License (MIT).
Copyright (c) 2015-2017, Ivan Kochurkin (kvanttt@gmail.com), Positive Technologies.
Copyright (c) 2017, Ivan Khudyashev (IHudyashov@ptsecurity.com)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

parser grammar SqlParser;

options { tokenVocab=SqlLexer; }

// details
selectStatement
    : SELECT selectElements FROM tableName=uid
      whereClause? groupbyClause? havingClause? orderByClause? limitClause? EOF
    ;
orderByClause
    : ORDER BY orderByExpression (',' orderByExpression)*
    ;

orderByExpression
    : expression order=(ASC | DESC)?
    ;

//tableSources
//    : tableSourceItem (',' tableSourceItem)*
//    ;


//tableSourceItem
//    : tableName=uid
//       (AS? alias=uid)?                                             #atomTableItem
//    | '(' tableSources ')'                                          #tableSourcesItem
//    ;

selectElements
    : selectElement  (',' selectElement)*
    ;

selectElement
    : (uid '.')? '*'                                               #selectStarElement
    | expressionAtom (AS uid)?                                     #selectExprElement
    ;

whereClause
    : WHERE expression
    ;

groupbyClause
    : GROUP BY expression (',' expression)*
    ;

havingClause
    : HAVING havingExpr=expression
    ;

limitClause
    : LIMIT
    (
      (offset=limitClauseAtom ',')? limit=limitClauseAtom
      | limit=limitClauseAtom OFFSET offset=limitClauseAtom
    )
    ;

limitClauseAtom
	: decimalLiteral
	;

//fullId
//    : uid (DOT_ID | '.' uid)?
//    ;

//tableName
//    : uid
//    ;

columnName
    : (tn=uid '.')? cn=uid
    ;
uid
    :ID
    ;
//
//dottedId
//    : DOT_ID
//    | '.' uid
//    ;


//    Literals

decimalLiteral
    : DECIMAL_LITERAL
    | ZERO_DECIMAL | ONE_DECIMAL | TWO_DECIMAL
    ;
stringLiteral
//    : (
//        STRING_CHARSET_NAME? STRING_LITERAL
//        | START_NATIONAL_STRING_LITERAL
//      ) STRING_LITERAL+
    : STRING_LITERAL+
//    | (
//        STRING_CHARSET_NAME? STRING_LITERAL
//        | START_NATIONAL_STRING_LITERAL
//      ) (COLLATE collationName)?
    ;

booleanLiteral
    : TRUE | FALSE;

//hexadecimalLiteral
//    : STRING_CHARSET_NAME? HEXADECIMAL_LITERAL;

//nullNotnull
//    : NOT? (NULL_LITERAL | NULL_SPEC_LITERAL)
//    ;

constant
    : stringLiteral                                                 #stringConstant
    | minusFlag='-'? decimalLiteral                                 #numberConstant
    | booleanLiteral                                                #boolConstant
//    | hexadecimalLiteral
//    | REAL_LITERAL | BIT_STRING
//    | NOT? nullLiteral=(NULL_LITERAL | NULL_SPEC_LITERAL)
    ;

//expressions
//    : expression (',' expression)*
//    ;

functionCall
    : aggregateWindowedFunction                                     #aggregateFunctionCall
    | uid '(' functionArgs? ')'                                     #udfFunctionCall
    ;


aggregateWindowedFunction
    : aggFunc=(AVG | MAX | MIN | SUM)
      '(' aggregator=(ALL | DISTINCT)? expression ')'               #commonAggregateFunc
    | COUNT '(' starArg='*' ')'                                     #countFunc1
    | COUNT '(' aggregator=(DISTINCT | ALL)? functionArgs ')'       #countFunc2
    ;

functionArgs
    : expression (','expression)*
    ;


//expression
//    : notOperator=(NOT | '!') expression                            #notExpression
//    | expression logicalOperator expression                         #logicalExpression
//    | predicate IS NOT? testValue=(TRUE | FALSE | UNKNOWN)          #isExpression
//    | predicate                                                     #predicateExpression
//    ;

expression
    : left=expression logicalOperator right=expression                         #logicalExpression
//    | predicate IS NOT? testValue=(TRUE | FALSE | UNKNOWN)          #isExpression
    | predicate                                                     #predicateExpression
    ;



//predicate
//    : predicate NOT? IN '(' expression (',' expression)* ')'        #inPredicate
//    | predicate IS nullNotnull                                      #isNullPredicate
//    | left=predicate comparisonOperator right=predicate             #binaryComparasionPredicate
//    | predicate NOT? BETWEEN predicate AND predicate                #betweenPredicate
//    | predicate SOUNDS LIKE predicate                               #soundsLikePredicate
//    | predicate NOT? LIKE predicate (ESCAPE STRING_LITERAL)?        #likePredicate
//    | predicate NOT? regex=(REGEXP | RLIKE) predicate               #regexpPredicate
//    | expressionAtom                                                #expressionAtomPredicate
//    ;

predicate
    : left=expressionAtom NOT? IN '(' expressionAtom (',' expressionAtom)* ')'        #inPredicate
    | expressionAtom IS NOT? NULL_LITERAL                                #isNullPredicate
    | left=expressionAtom comparisonOperator right=expressionAtom        #binaryComparasionPredicate
    | left=expressionAtom NOT? BETWEEN from=expressionAtom AND to=expressionAtom      #betweenPredicate
//    | left=expressionAtom SOUNDS LIKE likeStr=expressionAtom                          #soundsLikePredicate
    | left=expressionAtom NOT? LIKE likeStr=expressionAtom (ESCAPE STRING_LITERAL)?   #likePredicate
    | left=expressionAtom NOT? regex=(REGEXP | RLIKE) regStr=expressionAtom          #regexpPredicate
    | expressionAtom                                                     #expressionAtomPredicate
    ;


expressionAtom
    : constant                                                      #constantExpressionAtom
    | columnName                                                    #columnNameExpressionAtom
    | functionCall                                                  #functionCallExpressionAtom
    | unaryOperator expressionAtom                                  #unaryExpressionAtom
    | '(' expression ')'                                            #bracketedExpressionAtom
//    | left=expressionAtom bitOperator right=expressionAtom          #bitExpressionAtom
    | left=expressionAtom mathOperator right=expressionAtom         #mathExpressionAtom
    ;

unaryOperator
    : '!' | '~' | '+' | '-' | NOT
    ;

comparisonOperator
    : '=' | '>' | '<' | '<' '=' | '>' '='
    | '!' '='
    ;

logicalOperator
//    : AND | '&' '&' | XOR | OR | '|' '|'
    : AND | '&' '&' | OR | '|' '|'
    ;

//bitOperator
//    : '<' '<' | '>' '>' | '&' | '^' | '|'
//    ;

mathOperator
    : '*' | '/' | '%' | '+' | '-'
    ;
