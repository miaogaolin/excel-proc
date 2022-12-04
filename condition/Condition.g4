grammar Condition;

expr:
	expr op = (AND | OR) expr									# AndOr
	| conditionKey op = (IN | NOTIN) array						# GetArrayExpr
	| conditionKey op = (LIKE | NOTLIKE | EQ | NOTEQ) STRING	# GetStringExpr
	| conditionKey op = (EQ | GT | LT | NOTEQ) number			# GetNumExpr
	| '(' expr ')'												# Parent;

conditionKey: '{' COL '}';

array:
	'[' STRING (',' STRING)* ']' // 字符串数组
	| '[' number (',' number)* ']'; // 数字数组
number: FLOAT | DEC;

LIKE: '=~';
NOTLIKE: '!~';
EQ: '==';
NOTEQ: '!=';
GT: '>';
LT: '<';
IN: I N;
NOTIN: NOT ' '? IN;
NOT: N O T;
AND: A N D;
OR: O R;
COL: 'col' DEC;

ID: [a-zA-Z]+;
COMMENT: '//' ~[\r\n]* -> skip;
NL: '\r'? '\n';
FLOAT: DIGIT+ '.' DIGIT* | '.' DIGIT+;
DEC: DIGIT+;
STRING: '"' (ESC | .)*? '"' | '\'' (ESC | .)*? '\'';
WS: [ \t\r\n]+ -> skip;

// fragment 表示只能由词法调用
fragment ESC: '\\"' | '\\\\';
// 匹配字符\"和\\
fragment DIGIT: [0-9];

// 不区分大小写
fragment A: [aA];
fragment B: [bB];
fragment C: [cC];
fragment D: [dD];
fragment E: [eE];
fragment F: [fF];
fragment G: [gG];
fragment H: [hH];
fragment I: [iI];
fragment J: [jJ];
fragment K: [kK];
fragment L: [lL];
fragment M: [mM];
fragment N: [nN];
fragment O: [oO];
fragment P: [pP];
fragment Q: [qQ];
fragment R: [rR];
fragment S: [sS];
fragment T: [tT];
fragment U: [uU];
fragment V: [vV];
fragment W: [wW];
fragment X: [xX];
fragment Y: [yY];
fragment Z: [zZ];