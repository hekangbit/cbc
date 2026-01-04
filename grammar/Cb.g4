grammar Cb;
prog:   statement* EOF;

statement:  expr ';'
            | Identifier '=' expr ';'
            ;

expr:   expr ('*'|'/') expr
    |   expr ('+'|'-') expr
    |   '(' expr ')'
    |   IntLiteral
    |   Identifier
    ;

Identifier : [_a-zA-Z][_a-zA-Z0-9]*;

StringLiteral: '"' SChar*? '"';

IntLiteral: '0' | [1-9][0-9]* ;
DoubleLiteral: [0-9]+ ;

LINE_COMMENT: '//' .*? '\n' -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;

WS:     [ \t\n\r]+ -> skip;

fragment SChar
    : ~["\\\r\n]
    | '\\' [0-3][0-7][0-7]
    | '\\' [0-7] [0-7]?
    | '\\' [btnfr"'\\]
    | '\\\n'   // Added line
    | '\\\r\n' // Added line
    ;

fragment DIGIT : [0-9];