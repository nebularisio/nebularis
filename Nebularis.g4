grammar Nebularis;

// TODO: Complete operators
// TODO: Complete primitive types
// TODO: Complete literals
// TODO: Statements: for, do-while, try/error
// TODO: Const and globals
// TODO: Remove semicolons
// TODO: Coalesce operator
// TODO: fallthrough keyword
// TODO: Generic invocation
// TODO: variadic functions
// TODO: Switch expression
// TODO: pure lambda
// TODO: Reconcile type parameter ordering (lambda v.s. type decl v.s. fn decl)

compilationUnit:
    module=moduleDeclaration?
    imports+=importStatement*
    declarations+=declaration*
    EOF;

moduleDeclaration:
    attributes+=attribute* Module name=qualifiedName SemiColon;

importStatement:
    attributes+=attribute* Import module=moduleReference (As alias=Identifier)? SemiColon;

moduleReference:
    name=qualifiedName;

declaration:
    typeDecl=typeDeclaration
    | fnDecl=functionDeclaration;

functionDeclaration:
    attributes+=attribute*
    extern=Extern? pub=Public? pure=Pure? Fn
    typeParams=typeParameters?
    (ParenOpen target=referencedType ParenClose)?
    name=Identifier
    params=functionParameters
    returnType=typeSpec?
    constraints+=whereClause*
    body=codeBlock;

functionParameters:
    ParenOpen (params+=functionParameter (Comma params+=functionParameter)*)? ParenClose ;

functionParameter:
    name=Identifier? typeRef=typeSpec;

codeBlock:
    CurlyOpen statements+=statement* CurlyClose;

statement:
    Var name=Identifier (Colon typeRef=typeSpec)? (Equals initializer=expression)? SemiColon # VariableStatement
    | Return value=expression? SemiColon                                                            # ReturnStatement
    | If cond=expression onTrue=codeBlock (Else onFalse=codeBlock)?                                 # ConditionalStatement
    | While expression codeBlock                                                                    # WhileStatement
    | label=Identifier Colon                                                                        # LabelStatement
    | Continue label=Identifier? SemiColon                                                          # ContinueStatement
    | Break label=Identifier? SemiColon                                                             # BreakStatement
    | Switch expr=expression CurlyOpen (cases+=caseStatement|defauls+=defaultStatement)* CurlyClose # SwitchStatement
    | expression op=postfixOp  SemiColon                                                            # UnaryStatement
    | left=expression Equals right=expression SemiColon                                             # AssignmentStatement
    | Ellipsis                                                                                      # NotYetImplementedStatement
    | expression SemiColon                                                                          # ExpressionStatement;

caseStatement:
    Case expr+=expression (Comma expr+=expression)* Colon (statements+=statement+|block=codeBlock)?;

defaultStatement:
    Default Colon (statements+=statement+|block=codeBlock)?;

typeDeclaration:
    attributes+=attribute* Type name=Identifier typeParams=typeParameters? spec=typeSpec constraints+=whereClause* SemiColon?;

typeSpec:
     Question element=typeSpec                                               # NullableTypeSpec
    | SquareOpen SquareClose element=typeSpec                                # SpanTypeSpec
    | refType=referencedType                                                 # ReferencedTypeSpec
    | primitive=primitiveType                                                # PrimitiveTypeSpec
    | Fn params=functionParameters returnType=typeSpec?                      # FunctionTypeSpec
    | Struct CurlyOpen fields+=field* CurlyClose                             # StructTypeSpec
    | Interface extends=extendsClause? CurlyOpen methods+=method* CurlyClose # InterfaceTypeSpec
    ;

referencedType:
    name=qualifiedName (LessThan typeArgs+=typeSpec (Comma typeArgs+=typeSpec)* GreaterThan)?;

extendsClause:
    Colon (refTypes+=referencedType (Comma refTypes+=referencedType)*)?;

method:
    attributes+=attribute* name=Identifier params=functionParameters returnType=typeSpec? constraints+=whereClause* SemiColon?;

whereClause:
    Where expr=expression;

typeParameters:
    LessThan params+=typeParameter (Comma params+=typeParameter)* GreaterThan;

typeParameter:
    name=Identifier;

primitiveType:
    tInt32=Int32|tBool=Bool|tStr=Str|tByte=Byte|tChar=Char; // TODO: float, double other ints etc.

field:
    attributes+=attribute* name=Identifier typeRef=typeSpec constraints+=whereClause* SemiColon?;

expression:
    Ellipsis                                                                                # NotYetImplementedExpression
    | ident=Identifier                                                                      # IdentifierExpression
    | target=expression Dot member=Identifier                                               # MemberAccessExpression
    | target=expression ParenOpen (args+=expression (Comma args+=expression)*)? ParenClose  # InvocationExpression
    | target=expression SquareOpen index=expression SquareClose                             # IndexExpression
    | predicate=expression Question onTrue=expression Colon onFalse=expression              # TernaryExpression
    | value=literal                                                                         # LiteralExpression
    | ParenOpen expression ParenClose                                                       # ParenthesisExpression
    | Value                                                                                 # ValueExpression
    | structExpr                                                                            # StructExpression
    | lambdaExpr                                                                            # LambdaExpression
    | op=prefixOp expression                                                                # UnaryExpression
    | left=expression op=binaryOp right=expression                                          # BinaryExpression
    ;

lambdaExpr:
     Fn typeParams=typeParameters? params=functionParameters returnType=typeSpec? constraints+=whereClause* body=codeBlock;

structExpr:
    typeRef=referencedType CurlyOpen assignments=fieldAssignments CurlyClose;

fieldAssignments:
    ((fields+=fieldAssignment Comma)* fields+=fieldAssignment)?;

fieldAssignment:
    name=Identifier Colon value=expression;

prefixOp:
    Not
    | Tilda;

postfixOp:
    MinusMinus
    | PlusPlus;

binaryOp:
    LessThan
    |LessThanOrEqualTo
    |GreaterThan
    |GreaterThanOrEqualTo
    |EqualsEquals
    |NotEquals
    |Plus
    |Minus
    |Asterisk
    |Slash
    |Modulus
    |Caret
    |OrOr
    |AndAnd
    ;

attribute:
    SquareOpen value=structExpr SquareClose;

qualifiedName: parts+=Identifier ('.' parts+=Identifier)*;

literal:
    intLit=IntegerLiteral
    | boolLit=BoolLiteral
    | stringLit=StringLiteral;

IntegerLiteral:
    '-'? Digit+;

BoolLiteral:
    True | False;

StringLiteral:
    '"' ('\\"'|~'"')* '"';

As:
    'as';

Bool:
    'bool';

Break:
    'break';

Byte:
    'byte';

Case:
    'case';

Char:
    'char';

Continue:
    'continue';

Default:
    'default';

Else:
    'else';

Extern:
    'extern';

False:
    'false';

Fn:
    'fn';

If:
    'if';

Import:
    'import';

Interface:
    'interface';

Int32:
    'int32';

Module:
    'module';

Public:
    'public';

Pure:
    'pure';

Return:
    'return';

Str:
    'string';

Struct:
    'struct';

Switch:
    'switch';

True:
    'true';

Type:
    'type';

Value:
    'value';

Var:
    'var';

Where:
    'where';

While:
    'while';

AndAnd:
    '&&';

Asterisk:
    '*';

At:
    '@';

Caret:
    '^';

Colon:
    ':';

Comma:
    ',';

CurlyOpen:
    '{';

CurlyClose:
    '}';

Dot: '.';

Ellipsis:
    '...';

Equals:
    '=';

EqualsEquals:
    '==';

GreaterThan:
    '>';

GreaterThanOrEqualTo:
    '>=';

LessThan:
    '<';

LessThanOrEqualTo:
    '<=';

Plus:
    '+';

Minus:
    '-';

MinusMinus:
    '--';

Modulus:
    '%';

Not:
    '!';

NotEquals:
    '!=';

OrOr:
    '||';

ParenOpen:
    '(';

ParenClose:
    ')';

PlusPlus:
    '++';

Question: '?';

SemiColon:
    ';';

Slash:
    '/';

SquareOpen:
    '[';

SquareClose:
    ']';

Tilda:
    '~';

Identifier: IdentifierSegment;

fragment IdentifierSegment: (Letter|Underscore) (LetterOrDigit|Underscore)*;

LineComment
    : '//' ~[\r\n]*  -> skip;

BlockComment
    :   '/*' .*? '*/' -> skip;

fragment Letter: [a-zA-Z]; // TODO: Full range

fragment Digit: [0-9];

fragment Underscore: '_';

fragment LetterOrDigit: Letter | Digit;

WHITESPACE:  [ \t\r\n\u000C]+ -> skip;
