// Code generated from /Users/ozbenevren/dev/go/src/nebularis.io/nebularis/Nebularis.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Nebularis

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNebularisListener is a complete listener for a parse tree produced by NebularisParser.
type BaseNebularisListener struct{}

var _ NebularisListener = &BaseNebularisListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNebularisListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNebularisListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNebularisListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNebularisListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseNebularisListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseNebularisListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterModuleDeclaration is called when production moduleDeclaration is entered.
func (s *BaseNebularisListener) EnterModuleDeclaration(ctx *ModuleDeclarationContext) {}

// ExitModuleDeclaration is called when production moduleDeclaration is exited.
func (s *BaseNebularisListener) ExitModuleDeclaration(ctx *ModuleDeclarationContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BaseNebularisListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BaseNebularisListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterModuleReference is called when production moduleReference is entered.
func (s *BaseNebularisListener) EnterModuleReference(ctx *ModuleReferenceContext) {}

// ExitModuleReference is called when production moduleReference is exited.
func (s *BaseNebularisListener) ExitModuleReference(ctx *ModuleReferenceContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseNebularisListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseNebularisListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseNebularisListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseNebularisListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFunctionParameters is called when production functionParameters is entered.
func (s *BaseNebularisListener) EnterFunctionParameters(ctx *FunctionParametersContext) {}

// ExitFunctionParameters is called when production functionParameters is exited.
func (s *BaseNebularisListener) ExitFunctionParameters(ctx *FunctionParametersContext) {}

// EnterFunctionParameter is called when production functionParameter is entered.
func (s *BaseNebularisListener) EnterFunctionParameter(ctx *FunctionParameterContext) {}

// ExitFunctionParameter is called when production functionParameter is exited.
func (s *BaseNebularisListener) ExitFunctionParameter(ctx *FunctionParameterContext) {}

// EnterCodeBlock is called when production codeBlock is entered.
func (s *BaseNebularisListener) EnterCodeBlock(ctx *CodeBlockContext) {}

// ExitCodeBlock is called when production codeBlock is exited.
func (s *BaseNebularisListener) ExitCodeBlock(ctx *CodeBlockContext) {}

// EnterVariableStatement is called when production VariableStatement is entered.
func (s *BaseNebularisListener) EnterVariableStatement(ctx *VariableStatementContext) {}

// ExitVariableStatement is called when production VariableStatement is exited.
func (s *BaseNebularisListener) ExitVariableStatement(ctx *VariableStatementContext) {}

// EnterReturnStatement is called when production ReturnStatement is entered.
func (s *BaseNebularisListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production ReturnStatement is exited.
func (s *BaseNebularisListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterConditionalStatement is called when production ConditionalStatement is entered.
func (s *BaseNebularisListener) EnterConditionalStatement(ctx *ConditionalStatementContext) {}

// ExitConditionalStatement is called when production ConditionalStatement is exited.
func (s *BaseNebularisListener) ExitConditionalStatement(ctx *ConditionalStatementContext) {}

// EnterWhileStatement is called when production WhileStatement is entered.
func (s *BaseNebularisListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production WhileStatement is exited.
func (s *BaseNebularisListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterLabelStatement is called when production LabelStatement is entered.
func (s *BaseNebularisListener) EnterLabelStatement(ctx *LabelStatementContext) {}

// ExitLabelStatement is called when production LabelStatement is exited.
func (s *BaseNebularisListener) ExitLabelStatement(ctx *LabelStatementContext) {}

// EnterContinueStatement is called when production ContinueStatement is entered.
func (s *BaseNebularisListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production ContinueStatement is exited.
func (s *BaseNebularisListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterBreakStatement is called when production BreakStatement is entered.
func (s *BaseNebularisListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production BreakStatement is exited.
func (s *BaseNebularisListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterSwitchStatement is called when production SwitchStatement is entered.
func (s *BaseNebularisListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production SwitchStatement is exited.
func (s *BaseNebularisListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterUnaryStatement is called when production UnaryStatement is entered.
func (s *BaseNebularisListener) EnterUnaryStatement(ctx *UnaryStatementContext) {}

// ExitUnaryStatement is called when production UnaryStatement is exited.
func (s *BaseNebularisListener) ExitUnaryStatement(ctx *UnaryStatementContext) {}

// EnterAssignmentStatement is called when production AssignmentStatement is entered.
func (s *BaseNebularisListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production AssignmentStatement is exited.
func (s *BaseNebularisListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterNotYetImplementedStatement is called when production NotYetImplementedStatement is entered.
func (s *BaseNebularisListener) EnterNotYetImplementedStatement(ctx *NotYetImplementedStatementContext) {
}

// ExitNotYetImplementedStatement is called when production NotYetImplementedStatement is exited.
func (s *BaseNebularisListener) ExitNotYetImplementedStatement(ctx *NotYetImplementedStatementContext) {
}

// EnterExpressionStatement is called when production ExpressionStatement is entered.
func (s *BaseNebularisListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production ExpressionStatement is exited.
func (s *BaseNebularisListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterCaseStatement is called when production caseStatement is entered.
func (s *BaseNebularisListener) EnterCaseStatement(ctx *CaseStatementContext) {}

// ExitCaseStatement is called when production caseStatement is exited.
func (s *BaseNebularisListener) ExitCaseStatement(ctx *CaseStatementContext) {}

// EnterDefaultStatement is called when production defaultStatement is entered.
func (s *BaseNebularisListener) EnterDefaultStatement(ctx *DefaultStatementContext) {}

// ExitDefaultStatement is called when production defaultStatement is exited.
func (s *BaseNebularisListener) ExitDefaultStatement(ctx *DefaultStatementContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseNebularisListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseNebularisListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterNullableTypeSpec is called when production NullableTypeSpec is entered.
func (s *BaseNebularisListener) EnterNullableTypeSpec(ctx *NullableTypeSpecContext) {}

// ExitNullableTypeSpec is called when production NullableTypeSpec is exited.
func (s *BaseNebularisListener) ExitNullableTypeSpec(ctx *NullableTypeSpecContext) {}

// EnterSpanTypeSpec is called when production SpanTypeSpec is entered.
func (s *BaseNebularisListener) EnterSpanTypeSpec(ctx *SpanTypeSpecContext) {}

// ExitSpanTypeSpec is called when production SpanTypeSpec is exited.
func (s *BaseNebularisListener) ExitSpanTypeSpec(ctx *SpanTypeSpecContext) {}

// EnterReferencedTypeSpec is called when production ReferencedTypeSpec is entered.
func (s *BaseNebularisListener) EnterReferencedTypeSpec(ctx *ReferencedTypeSpecContext) {}

// ExitReferencedTypeSpec is called when production ReferencedTypeSpec is exited.
func (s *BaseNebularisListener) ExitReferencedTypeSpec(ctx *ReferencedTypeSpecContext) {}

// EnterPrimitiveTypeSpec is called when production PrimitiveTypeSpec is entered.
func (s *BaseNebularisListener) EnterPrimitiveTypeSpec(ctx *PrimitiveTypeSpecContext) {}

// ExitPrimitiveTypeSpec is called when production PrimitiveTypeSpec is exited.
func (s *BaseNebularisListener) ExitPrimitiveTypeSpec(ctx *PrimitiveTypeSpecContext) {}

// EnterFunctionTypeSpec is called when production FunctionTypeSpec is entered.
func (s *BaseNebularisListener) EnterFunctionTypeSpec(ctx *FunctionTypeSpecContext) {}

// ExitFunctionTypeSpec is called when production FunctionTypeSpec is exited.
func (s *BaseNebularisListener) ExitFunctionTypeSpec(ctx *FunctionTypeSpecContext) {}

// EnterStructTypeSpec is called when production StructTypeSpec is entered.
func (s *BaseNebularisListener) EnterStructTypeSpec(ctx *StructTypeSpecContext) {}

// ExitStructTypeSpec is called when production StructTypeSpec is exited.
func (s *BaseNebularisListener) ExitStructTypeSpec(ctx *StructTypeSpecContext) {}

// EnterInterfaceTypeSpec is called when production InterfaceTypeSpec is entered.
func (s *BaseNebularisListener) EnterInterfaceTypeSpec(ctx *InterfaceTypeSpecContext) {}

// ExitInterfaceTypeSpec is called when production InterfaceTypeSpec is exited.
func (s *BaseNebularisListener) ExitInterfaceTypeSpec(ctx *InterfaceTypeSpecContext) {}

// EnterReferencedType is called when production referencedType is entered.
func (s *BaseNebularisListener) EnterReferencedType(ctx *ReferencedTypeContext) {}

// ExitReferencedType is called when production referencedType is exited.
func (s *BaseNebularisListener) ExitReferencedType(ctx *ReferencedTypeContext) {}

// EnterExtendsClause is called when production extendsClause is entered.
func (s *BaseNebularisListener) EnterExtendsClause(ctx *ExtendsClauseContext) {}

// ExitExtendsClause is called when production extendsClause is exited.
func (s *BaseNebularisListener) ExitExtendsClause(ctx *ExtendsClauseContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseNebularisListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseNebularisListener) ExitMethod(ctx *MethodContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseNebularisListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseNebularisListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterTypeParameters is called when production typeParameters is entered.
func (s *BaseNebularisListener) EnterTypeParameters(ctx *TypeParametersContext) {}

// ExitTypeParameters is called when production typeParameters is exited.
func (s *BaseNebularisListener) ExitTypeParameters(ctx *TypeParametersContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseNebularisListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseNebularisListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseNebularisListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseNebularisListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterField is called when production field is entered.
func (s *BaseNebularisListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseNebularisListener) ExitField(ctx *FieldContext) {}

// EnterTernaryExpression is called when production TernaryExpression is entered.
func (s *BaseNebularisListener) EnterTernaryExpression(ctx *TernaryExpressionContext) {}

// ExitTernaryExpression is called when production TernaryExpression is exited.
func (s *BaseNebularisListener) ExitTernaryExpression(ctx *TernaryExpressionContext) {}

// EnterMemberAccessExpression is called when production MemberAccessExpression is entered.
func (s *BaseNebularisListener) EnterMemberAccessExpression(ctx *MemberAccessExpressionContext) {}

// ExitMemberAccessExpression is called when production MemberAccessExpression is exited.
func (s *BaseNebularisListener) ExitMemberAccessExpression(ctx *MemberAccessExpressionContext) {}

// EnterParenthesisExpression is called when production ParenthesisExpression is entered.
func (s *BaseNebularisListener) EnterParenthesisExpression(ctx *ParenthesisExpressionContext) {}

// ExitParenthesisExpression is called when production ParenthesisExpression is exited.
func (s *BaseNebularisListener) ExitParenthesisExpression(ctx *ParenthesisExpressionContext) {}

// EnterLiteralExpression is called when production LiteralExpression is entered.
func (s *BaseNebularisListener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production LiteralExpression is exited.
func (s *BaseNebularisListener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterIndexExpression is called when production IndexExpression is entered.
func (s *BaseNebularisListener) EnterIndexExpression(ctx *IndexExpressionContext) {}

// ExitIndexExpression is called when production IndexExpression is exited.
func (s *BaseNebularisListener) ExitIndexExpression(ctx *IndexExpressionContext) {}

// EnterUnaryExpression is called when production UnaryExpression is entered.
func (s *BaseNebularisListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production UnaryExpression is exited.
func (s *BaseNebularisListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterInvocationExpression is called when production InvocationExpression is entered.
func (s *BaseNebularisListener) EnterInvocationExpression(ctx *InvocationExpressionContext) {}

// ExitInvocationExpression is called when production InvocationExpression is exited.
func (s *BaseNebularisListener) ExitInvocationExpression(ctx *InvocationExpressionContext) {}

// EnterIdentifierExpression is called when production IdentifierExpression is entered.
func (s *BaseNebularisListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production IdentifierExpression is exited.
func (s *BaseNebularisListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterStructExpression is called when production StructExpression is entered.
func (s *BaseNebularisListener) EnterStructExpression(ctx *StructExpressionContext) {}

// ExitStructExpression is called when production StructExpression is exited.
func (s *BaseNebularisListener) ExitStructExpression(ctx *StructExpressionContext) {}

// EnterBinaryExpression is called when production BinaryExpression is entered.
func (s *BaseNebularisListener) EnterBinaryExpression(ctx *BinaryExpressionContext) {}

// ExitBinaryExpression is called when production BinaryExpression is exited.
func (s *BaseNebularisListener) ExitBinaryExpression(ctx *BinaryExpressionContext) {}

// EnterValueExpression is called when production ValueExpression is entered.
func (s *BaseNebularisListener) EnterValueExpression(ctx *ValueExpressionContext) {}

// ExitValueExpression is called when production ValueExpression is exited.
func (s *BaseNebularisListener) ExitValueExpression(ctx *ValueExpressionContext) {}

// EnterNotYetImplementedExpression is called when production NotYetImplementedExpression is entered.
func (s *BaseNebularisListener) EnterNotYetImplementedExpression(ctx *NotYetImplementedExpressionContext) {
}

// ExitNotYetImplementedExpression is called when production NotYetImplementedExpression is exited.
func (s *BaseNebularisListener) ExitNotYetImplementedExpression(ctx *NotYetImplementedExpressionContext) {
}

// EnterLambdaExpression is called when production LambdaExpression is entered.
func (s *BaseNebularisListener) EnterLambdaExpression(ctx *LambdaExpressionContext) {}

// ExitLambdaExpression is called when production LambdaExpression is exited.
func (s *BaseNebularisListener) ExitLambdaExpression(ctx *LambdaExpressionContext) {}

// EnterLambdaExpr is called when production lambdaExpr is entered.
func (s *BaseNebularisListener) EnterLambdaExpr(ctx *LambdaExprContext) {}

// ExitLambdaExpr is called when production lambdaExpr is exited.
func (s *BaseNebularisListener) ExitLambdaExpr(ctx *LambdaExprContext) {}

// EnterStructExpr is called when production structExpr is entered.
func (s *BaseNebularisListener) EnterStructExpr(ctx *StructExprContext) {}

// ExitStructExpr is called when production structExpr is exited.
func (s *BaseNebularisListener) ExitStructExpr(ctx *StructExprContext) {}

// EnterFieldAssignments is called when production fieldAssignments is entered.
func (s *BaseNebularisListener) EnterFieldAssignments(ctx *FieldAssignmentsContext) {}

// ExitFieldAssignments is called when production fieldAssignments is exited.
func (s *BaseNebularisListener) ExitFieldAssignments(ctx *FieldAssignmentsContext) {}

// EnterFieldAssignment is called when production fieldAssignment is entered.
func (s *BaseNebularisListener) EnterFieldAssignment(ctx *FieldAssignmentContext) {}

// ExitFieldAssignment is called when production fieldAssignment is exited.
func (s *BaseNebularisListener) ExitFieldAssignment(ctx *FieldAssignmentContext) {}

// EnterPrefixOp is called when production prefixOp is entered.
func (s *BaseNebularisListener) EnterPrefixOp(ctx *PrefixOpContext) {}

// ExitPrefixOp is called when production prefixOp is exited.
func (s *BaseNebularisListener) ExitPrefixOp(ctx *PrefixOpContext) {}

// EnterPostfixOp is called when production postfixOp is entered.
func (s *BaseNebularisListener) EnterPostfixOp(ctx *PostfixOpContext) {}

// ExitPostfixOp is called when production postfixOp is exited.
func (s *BaseNebularisListener) ExitPostfixOp(ctx *PostfixOpContext) {}

// EnterBinaryOp is called when production binaryOp is entered.
func (s *BaseNebularisListener) EnterBinaryOp(ctx *BinaryOpContext) {}

// ExitBinaryOp is called when production binaryOp is exited.
func (s *BaseNebularisListener) ExitBinaryOp(ctx *BinaryOpContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseNebularisListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseNebularisListener) ExitAttribute(ctx *AttributeContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseNebularisListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseNebularisListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseNebularisListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseNebularisListener) ExitLiteral(ctx *LiteralContext) {}
