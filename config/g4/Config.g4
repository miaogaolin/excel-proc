grammar Config;

stat: NL* fields NL* sections  # FieldStat
| NL* sections # NotFieldStat
;

sections: (NL* section NL*)+;
section: .+? (NL NL | EOF) # GetSection;

fields: (field NL+)+;
field: fieldName ' '* ':' ' '* fieldValue;
fieldName: ID DEC?;
fieldValue: STRING | FLOAT | DEC;

ID: [a-zA-Z_][a-zA-Z_0-9]*;
COMMENT: '//' ~[\r\n]* -> skip;
NL: '\r'? '\n';
FLOAT: DIGIT+ '.' DIGIT* | '.' DIGIT+;
DEC: DIGIT+;
STRING: '"' (ESC | .)*? '"' | '\'' (ESC | .)*? '\'';
OTHER: .;
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