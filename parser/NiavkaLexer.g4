lexer grammar NiavkaLexer;

END: 'кінець';
DIIA: 'дія';

SKIP_SPACES: ( ' ' | '\t' ) -> skip;
NL: ( '\r'? '\n' );
OPEN_PAREN: '(';
CLOSE_PAREN: ')';
OPEN_ARRAY: '[';
CLOSE_ARRAY: ']';
COMMA: ',';
ASSIGN: '=';
PLUS: '+';
MINUS: '-';
MUL: '*';
DIV: '/';
DOT: '.';
ID: ID_START ID_CONTINUE*;
NUMBER: FLOAT | INTEGER;
INTEGER: DIGIT+;
FLOAT: INTEGER '.' INTEGER;
STRING: '"' ( ~[\\"] | ESCAPE_CHAR )* '"';

fragment DIGIT
    : '0'..'9'
    ;

fragment ID_START
    : 'a'..'z' | 'A'..'Z'
    | 'а'..'я' | 'А'..'Я' | 'і' | 'І' | 'ї' | 'Ї' | 'є' | 'Є' | 'ґ' | 'Ґ'
    | '_'
    ;

fragment ID_CONTINUE
    : ID_START
    | '0'..'9'
    | 'ʼ'
    | ID_START
    ;

fragment ESCAPE_CHAR
    : '\\' [0btnfr"'\\]
    ;
