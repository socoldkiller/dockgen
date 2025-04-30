grammar Bash;

commandLine
    : pipeline EOF
    | EOF;

pipeline
    : command (PIPE command)*
    ;

optionWithArg
    : option (arg)?
    ;

command
    : prog? assign
    | redir* prog (optionWithArg | arg)* redir* (logicalOp command)?
    ;

prog
    : WORD
    ;

option
    : OPTION
    ;

arg
    : WORD
    | variable
    ;

redir
    : (LT | GT | DGT | LTGT) WORD
    ;

variable
    : '$' '('* WORD ')'*
    ;

logicalOp
    : ANDAND    // &&
    | OROR      // ||
    | SEMI      // ;
    ;

assign
    : WORD ASSIGN (WORD | variable)
    ;



ASSIGN: '=';
PIPE: '|';
LT: '<';
GT: '>';
DGT: '>>';
LTGT: '<>';
ANDAND: '&&';
OROR: '||';
SEMI: ';';

fragment DASH: '-';
OPTION: DASH+ ~[ \t\r\n|<>;&=()$`'"\\]+;

WORD: ~[ \t\r\n|<>;&=()`\\]+;

// ------------ Whitespace ------------

WS: [ \t\r\n]+ -> skip;
