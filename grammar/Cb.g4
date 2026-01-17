grammar Cb;
prog
    : importStmts topDefs EOF;

importStmts
    : (importStmt)*;
importStmt
    : 'import' Identifier('.'Identifier)* ';';

topDefs
    : (defVars | defFunc | typedef)*;
defVars
    : 'static'? type Identifier ('=' expr)? (',' Identifier ('=' expr)?)* ';';
defFunc
    : 'static'? typeref Identifier '(' params ')' block;
typedef
    : 'typedef' type Identifier ';';

stmt
    : expr ';'
    | block
    | ifStmt
    | whileStmt
    | forStmt
    | breakStmt
    | continueStmt
    | gotoStmt
    | returnStmt
    ;
block
    : '{' (defVars)* (stmt)* '}';
ifStmt
    : 'if' '(' expr ')' stmt ('else' stmt)?;
whileStmt
    : 'while' '(' expr ')' stmt;
forStmt
    : 'for' '(' expr ';' expr ';' expr ')' stmt;
breakStmt
    : 'break' ';';
continueStmt
    : 'continue' ';';
gotoStmt
    : 'goto' Identifier ';';
returnStmt
    : 'return' (expr)? ';';

type
    : typeref;
typeref
    : typerefBase ('['']' | '[' IntLiteral ']' | '*' | '(' paramTyperefs')')*;
typerefBase
    : 'void'
    | 'char'
    | 'short'
    | 'int'
    | 'long'
    | 'unsigned char'
    | 'unsigned short'
    | 'unsigned' 'int'
    | 'unsigned' 'long'
    | 'struct' Identifier
    ;
params
    : 'void' | fixedParams (',' '...')? ;
fixedParams
    : param (',' param)*;
param
    : type Identifier;
paramTyperefs
    : 'void' | fixedparamTyperefs (',' '...')? ;
fixedparamTyperefs
    : typeref (',' typeref)*;

expr
    : term assignOp expr #AssignExpr
    | expr10 #CondExpr
    ;
assignOp
    : '='
    | '+='
    | '-='
    | '*='
    | '/='
    ;
expr10
    : expr9 ('?' expr ':' expr10)?
    ;
expr9
    : expr8 ('||' expr8)*
    ;
expr8
    : expr7 ('&&' expr7)*
    ;
expr7
    : expr6 ('>' expr6 | '<' expr6 | '>=' expr6 | '<=' expr6 | '==' expr6 | '!=' expr6)*
    ;
expr6
    : expr5 ('|' expr5)*
    ;
expr5
    : expr4 ('^' expr4)*
    ;
expr4
    : expr3 ('&' expr3)*
    ;
expr3
    : expr2 ('>>' expr2 | '<<' expr2)*
    ;
expr2
    : expr1 ('+' expr1 | '-' expr1)*
    ;
expr1
    : term ('*' term | '/' term | '%' term)*
    ;

term
    : castExpr
    | unary
    ;
castExpr
    : '(' type ')' term;
unary
    : '++' unary
    | '--' unary
    | '+' unary
    | '-' unary
    | '!' unary
    | '~' unary
    | '*' unary
    | '&' unary
    | 'sizeof' '(' type ')'
    | 'sizeof' unary
    | postfix
    ;
postfix
    : primary ('++' | '--' | '[' expr ']' | '.' Identifier | '->' Identifier | '(' args ')')*
    ;
args
    : expr (',' expr)*
    ;
primary
    : IntLiteral
    | Character
    | StringLiteral
    | Identifier
    | '(' expr ')'
    ;


MUL : '*' ;
DIV : '/' ;
ADD : '+' ;
SUB : '-' ;

Identifier
    : [_a-zA-Z][_a-zA-Z0-9]*;
Character
    : '\'' Char '\'';
StringLiteral
    : '"' SChar* '"';
IntLiteral
    : '0' | [1-9][0-9]* ;
DoubleLiteral
    : [0-9]+ ;

LineComment
    : '//' ~[\r\n]* -> skip;
BlockComment
    : '/*' .*? '*/' -> skip;
WhiteSpace
    : [ \t\n\r]+ -> skip;

fragment DIGIT
    : [0-9];
fragment Char
    : ~['\\\r\n]
    | '\\' EscapedSequence
    ;
fragment SChar
    : ~["\\\r\n]          // Ordinary characters (excluding double quotes, backslashes, and line breaks)
    | EscapedSequence     // Escaped sequence
    | LineContinuation    // Line continuation
    ;

fragment EscapedSequence
    : '\\' CommonEscape
    | '\\' OctalEscape
    ;

fragment CommonEscape
    : ['"\\]           // special char
    | [bfnrt]          // control char
    | 'v'              // vertical table
    | '0'              // empty char
    ;

fragment OctalEscape
    : [0-3][0-7][0-7]  // \ooo (1-3 bits octal number)
    | [0-7][0-7]?      // \o or \oo
    ;

fragment LineContinuation
    : '\\' '\r'? '\n'  // line concatenation with backslashes
    ;