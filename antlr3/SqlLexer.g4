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

lexer grammar SqlLexer;

channels { MYSQLCOMMENT, ERRORCHANNEL }

// SKIP

SPACE:                               [ \t\r\n]+    -> channel(HIDDEN);
SPEC_MYSQL_COMMENT:                  '/*!' .+? '*/' -> channel(MYSQLCOMMENT);
COMMENT_INPUT:                       '/*' .*? '*/' -> channel(HIDDEN);
LINE_COMMENT:                        (
                                       ('-- ' | '#') ~[\r\n]* ('\r'? '\n' | EOF)
                                       | '--' ('\r'? '\n' | EOF)
                                     ) -> channel(HIDDEN);


// Keywords
// Common Keywords

ADD:                                 'ADD';
ALL:                                 'ALL';
ALTER:                               'ALTER';
ALWAYS:                              'ALWAYS';
ANALYZE:                             'ANALYZE';
AND:                                 'AND';
AS:                                  'AS';
ASC:                                 'ASC';
BEFORE:                              'BEFORE';
BETWEEN:                             'BETWEEN';
BOTH:                                'BOTH';
BY:                                  'BY';
CALL:                                'CALL';
CASCADE:                             'CASCADE';
CASE:                                'CASE';
CAST:                                'CAST';
CHANGE:                              'CHANGE';
CHARACTER:                           'CHARACTER';
CHECK:                               'CHECK';
COLLATE:                             'COLLATE';
COLUMN:                              'COLUMN';
CONDITION:                           'CONDITION';
CONSTRAINT:                          'CONSTRAINT';
CONTINUE:                            'CONTINUE';
CONVERT:                             'CONVERT';
CREATE:                              'CREATE';
CROSS:                               'CROSS';
CURRENT:                             'CURRENT';
CURRENT_USER:                        'CURRENT_USER';
CURSOR:                              'CURSOR';
DATABASE:                            'DATABASE';
DATABASES:                           'DATABASES';
DECLARE:                             'DECLARE';
DEFAULT:                             'DEFAULT';
DELAYED:                             'DELAYED';
DELETE:                              'DELETE';
DESC:                                'DESC';
DESCRIBE:                            'DESCRIBE';
DETERMINISTIC:                       'DETERMINISTIC';
DIAGNOSTICS:                         'DIAGNOSTICS';
DISTINCT:                            'DISTINCT';
DISTINCTROW:                         'DISTINCTROW';
DROP:                                'DROP';
EACH:                                'EACH';
ELSE:                                'ELSE';
ELSEIF:                              'ELSEIF';
ENCLOSED:                            'ENCLOSED';
ESCAPED:                             'ESCAPED';
EXISTS:                              'EXISTS';
EXIT:                                'EXIT';
EXPLAIN:                             'EXPLAIN';
FALSE:                               'FALSE';
FETCH:                               'FETCH';
FOR:                                 'FOR';
FORCE:                               'FORCE';
FOREIGN:                             'FOREIGN';
FROM:                                'FROM';
FULLTEXT:                            'FULLTEXT';
GENERATED:                           'GENERATED';
GET:                                 'GET';
GRANT:                               'GRANT';
GROUP:                               'GROUP';
HAVING:                              'HAVING';
HIGH_PRIORITY:                       'HIGH_PRIORITY';
IF:                                  'IF';
IGNORE:                              'IGNORE';
IN:                                  'IN';
INDEX:                               'INDEX';
INFILE:                              'INFILE';
INNER:                               'INNER';
INOUT:                               'INOUT';
INSERT:                              'INSERT';
INTERVAL:                            'INTERVAL';
INTO:                                'INTO';
IS:                                  'IS';
ITERATE:                             'ITERATE';
JOIN:                                'JOIN';
KEY:                                 'KEY';
KEYS:                                'KEYS';
KILL:                                'KILL';
LEADING:                             'LEADING';
LEAVE:                               'LEAVE';
LEFT:                                'LEFT';
LIKE:                                'LIKE';
LIMIT:                               'LIMIT';
LINEAR:                              'LINEAR';
LINES:                               'LINES';
LOAD:                                'LOAD';
LOCK:                                'LOCK';
LOOP:                                'LOOP';
LOW_PRIORITY:                        'LOW_PRIORITY';
MASTER_BIND:                         'MASTER_BIND';
MASTER_SSL_VERIFY_SERVER_CERT:       'MASTER_SSL_VERIFY_SERVER_CERT';
MATCH:                               'MATCH';
MAXVALUE:                            'MAXVALUE';
MODIFIES:                            'MODIFIES';
NATURAL:                             'NATURAL';
NOT:                                 'NOT';
NO_WRITE_TO_BINLOG:                  'NO_WRITE_TO_BINLOG';
NULL_LITERAL:                        'NULL';
NUMBER:                              'NUMBER';
ON:                                  'ON';
OPTIMIZE:                            'OPTIMIZE';
OPTION:                              'OPTION';
OPTIONALLY:                          'OPTIONALLY';
OR:                                  'OR';
ORDER:                               'ORDER';
OUT:                                 'OUT';
OUTER:                               'OUTER';
OUTFILE:                             'OUTFILE';
PARTITION:                           'PARTITION';
PRIMARY:                             'PRIMARY';
PROCEDURE:                           'PROCEDURE';
PURGE:                               'PURGE';
RANGE:                               'RANGE';
READ:                                'READ';
READS:                               'READS';
REFERENCES:                          'REFERENCES';
REGEXP:                              'REGEXP';
RELEASE:                             'RELEASE';
RENAME:                              'RENAME';
REPEAT:                              'REPEAT';
REPLACE:                             'REPLACE';
REQUIRE:                             'REQUIRE';
RESIGNAL:                            'RESIGNAL';
RESTRICT:                            'RESTRICT';
RETURN:                              'RETURN';
REVOKE:                              'REVOKE';
RIGHT:                               'RIGHT';
RLIKE:                               'RLIKE';
SCHEMA:                              'SCHEMA';
SCHEMAS:                             'SCHEMAS';
SELECT:                              'SELECT';
SET:                                 'SET';
SEPARATOR:                           'SEPARATOR';
SHOW:                                'SHOW';
SIGNAL:                              'SIGNAL';
SPATIAL:                             'SPATIAL';
SQL:                                 'SQL';
SQLEXCEPTION:                        'SQLEXCEPTION';
SQLSTATE:                            'SQLSTATE';
SQLWARNING:                          'SQLWARNING';
SQL_BIG_RESULT:                      'SQL_BIG_RESULT';
SQL_CALC_FOUND_ROWS:                 'SQL_CALC_FOUND_ROWS';
SQL_SMALL_RESULT:                    'SQL_SMALL_RESULT';
SSL:                                 'SSL';
STACKED:                             'STACKED';
STARTING:                            'STARTING';
STRAIGHT_JOIN:                       'STRAIGHT_JOIN';
TABLE:                               'TABLE';
TERMINATED:                          'TERMINATED';
THEN:                                'THEN';
TO:                                  'TO';
TRAILING:                            'TRAILING';
TRIGGER:                             'TRIGGER';
TRUE:                                'TRUE';
UNDO:                                'UNDO';
UNION:                               'UNION';
UNIQUE:                              'UNIQUE';
UNLOCK:                              'UNLOCK';
UNSIGNED:                            'UNSIGNED';
UPDATE:                              'UPDATE';
USAGE:                               'USAGE';
USE:                                 'USE';
USING:                               'USING';
VALUES:                              'VALUES';
WHEN:                                'WHEN';
WHERE:                               'WHERE';
WHILE:                               'WHILE';
WITH:                                'WITH';
WRITE:                               'WRITE';
XOR:                                 'XOR';
ZEROFILL:                            'ZEROFILL';

AVG:                                 'AVG';
COUNT:                               'COUNT';
MAX:                                 'MAX';
MIN:                                 'MIN';
SUM:                                 'SUM';

OFFSET:                              'OFFSET';
ESCAPE:                              'ESCAPE';

// Operators
// Operators. Assigns

VAR_ASSIGN:                          ':=';
PLUS_ASSIGN:                         '+=';
MINUS_ASSIGN:                        '-=';
MULT_ASSIGN:                         '*=';
DIV_ASSIGN:                          '/=';
MOD_ASSIGN:                          '%=';
AND_ASSIGN:                          '&=';
XOR_ASSIGN:                          '^=';
OR_ASSIGN:                           '|=';


// Operators. Arithmetics

STAR:                                '*';
DIVIDE:                              '/';
MODULE:                              '%';
PLUS:                                '+';
MINUSMINUS:                          '--';
MINUS:                               '-';
DIV:                                 'DIV';
MOD:                                 'MOD';


// Operators. Comparation

EQUAL_SYMBOL:                        '=';
GREATER_SYMBOL:                      '>';
LESS_SYMBOL:                         '<';
EXCLAMATION_SYMBOL:                  '!';


// Operators. Bit

BIT_NOT_OP:                          '~';
BIT_OR_OP:                           '|';
BIT_AND_OP:                          '&';
BIT_XOR_OP:                          '^';


// Constructors symbols

DOT:                                 '.';
LR_BRACKET:                          '(';
RR_BRACKET:                          ')';
COMMA:                               ',';
SEMI:                                ';';
AT_SIGN:                             '@';
ZERO_DECIMAL:                        '0';
ONE_DECIMAL:                         '1';
TWO_DECIMAL:                         '2';
SINGLE_QUOTE_SYMB:                   '\'';
DOUBLE_QUOTE_SYMB:                   '"';
REVERSE_QUOTE_SYMB:                  '`';
COLON_SYMB:                          ':';

fragment QUOTE_SYMB
    : SINGLE_QUOTE_SYMB | DOUBLE_QUOTE_SYMB | REVERSE_QUOTE_SYMB
    ;

// File's sizes


FILESIZE_LITERAL:                    DEC_DIGIT+ ('K'|'M'|'G'|'T');



// Literal Primitives


START_NATIONAL_STRING_LITERAL:       'N' SQUOTA_STRING;
STRING_LITERAL:                      DQUOTA_STRING | SQUOTA_STRING | BQUOTA_STRING;
DECIMAL_LITERAL:                     DEC_DIGIT+;
HEXADECIMAL_LITERAL:                 'X' '\'' (HEX_DIGIT HEX_DIGIT)+ '\''
                                     | '0X' HEX_DIGIT+;

REAL_LITERAL:                        (DEC_DIGIT+)? '.' DEC_DIGIT+
                                     | DEC_DIGIT+ '.' EXPONENT_NUM_PART
                                     | (DEC_DIGIT+)? '.' (DEC_DIGIT+ EXPONENT_NUM_PART)
                                     | DEC_DIGIT+ EXPONENT_NUM_PART;
NULL_SPEC_LITERAL:                   '\\' 'N';
//BIT_STRING:                          BIT_STRINfunctionCallG_L;




// Hack for dotID
// Prevent recognize string:         .123somelatin AS ((.123), FLOAT_LITERAL), ((somelatin), ID)
//  it must recoginze:               .123somelatin AS ((.), DOT), (123somelatin, ID)

DOT_ID:                              '.' ID_LITERAL;



// Identifiers

ID:                                  ID_LITERAL;
// DOUBLE_QUOTE_ID:                  '"' ~'"'+ '"';
REVERSE_QUOTE_ID:                    '`' ~'`'+ '`';
STRING_USER_NAME:                    (
                                       SQUOTA_STRING | DQUOTA_STRING
                                       | BQUOTA_STRING | ID_LITERAL
                                     ) '@'
                                     (
                                       SQUOTA_STRING | DQUOTA_STRING
                                       | BQUOTA_STRING | ID_LITERAL
                                     );
LOCAL_ID:                            '@'
                                (
                                  [A-Z0-9._$]+
                                  | SQUOTA_STRING
                                  | DQUOTA_STRING
                                  | BQUOTA_STRING
                                );
GLOBAL_ID:                           '@' '@'
                                (
                                  [A-Z0-9._$]+
                                  | BQUOTA_STRING
                                );


fragment EXPONENT_NUM_PART:          'E' [-+]? DEC_DIGIT+;
//fragment ID_LITERAL:                 [A-Z_$0-9]*?[A-Z_$]+?[A-Z_$0-9]*;
fragment DQUOTA_STRING:              '"' ( '\\'. | '""' | ~('"'| '\\') )* '"';
fragment SQUOTA_STRING:              '\'' ('\\'. | '\'\'' | ~('\'' | '\\'))* '\'';
fragment BQUOTA_STRING:              '`' ( '\\'. | '``' | ~('`'|'\\'))* '`';
fragment HEX_DIGIT:                  [0-9A-F];
fragment DEC_DIGIT:                  [0-9];
fragment BIT_STRING_L:               'B' '\'' [01]+ '\'';

fragment DIGIT:[0-9];   //匹配数字
fragment LETTER:[a-zA-Z];  //匹配字母
ID_LITERAL    //匹配只含有数字字母和下划线的文本
    : (LETTER | DIGIT | '_')+
    ;


// Last tokens must generate Errors

ERROR_RECONGNIGION:                  .    -> channel(ERRORCHANNEL);
