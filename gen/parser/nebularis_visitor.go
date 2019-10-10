// Code generated from /Users/ozbenevren/dev/go/src/nebularis.io/nebularis/Nebularis.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Nebularis

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by NebularisParser.
type NebularisVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by NebularisParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by NebularisParser#moduleDeclaration.
	VisitModuleDeclaration(ctx *ModuleDeclarationContext) interface{}

	// Visit a parse tree produced by NebularisParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by NebularisParser#moduleReference.
	VisitModuleReference(ctx *ModuleReferenceContext) interface{}

	// Visit a parse tree produced by NebularisParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by NebularisParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by NebularisParser#functionParameters.
	VisitFunctionParameters(ctx *FunctionParametersContext) interface{}

	// Visit a parse tree produced by NebularisParser#functionParameter.
	VisitFunctionParameter(ctx *FunctionParameterContext) interface{}

	// Visit a parse tree produced by NebularisParser#codeBlock.
	VisitCodeBlock(ctx *CodeBlockContext) interface{}

	// Visit a parse tree produced by NebularisParser#VariableStatement.
	VisitVariableStatement(ctx *VariableStatementContext) interface{}

	// Visit a parse tree produced by NebularisParser#ReturnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by NebularisParser#ConditionalStatement.
	VisitConditionalStatement(ctx *ConditionalStatementContext) interface{}

	// Visit a parse tree produced by NebularisParser#WhileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by NebularisParser#LabelStatement.
	VisitLabelStatement(ctx *LabelStatementContext) interface{}

	// Visit a parse tree produced by NebularisParser#ContinueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by NebularisParser#BreakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by NebularisParser#SwitchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by NebularisParser#UnaryStatement.
	VisitUnaryStatement(ctx *UnaryStatementContext) interface{}

	// Visit a parse tree produced by NebularisParser#AssignmentStatement.
	VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{}

	// Visit a parse tree produced by NebularisParser#NotYetImplementedStatement.
	VisitNotYetImplementedStatement(ctx *NotYetImplementedStatementContext) interface{}

	// Visit a parse tree produced by NebularisParser#ExpressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by NebularisParser#caseStatement.
	VisitCaseStatement(ctx *CaseStatementContext) interface{}

	// Visit a parse tree produced by NebularisParser#defaultStatement.
	VisitDefaultStatement(ctx *DefaultStatementContext) interface{}

	// Visit a parse tree produced by NebularisParser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by NebularisParser#NullableTypeSpec.
	VisitNullableTypeSpec(ctx *NullableTypeSpecContext) interface{}

	// Visit a parse tree produced by NebularisParser#SpanTypeSpec.
	VisitSpanTypeSpec(ctx *SpanTypeSpecContext) interface{}

	// Visit a parse tree produced by NebularisParser#ReferencedTypeSpec.
	VisitReferencedTypeSpec(ctx *ReferencedTypeSpecContext) interface{}

	// Visit a parse tree produced by NebularisParser#PrimitiveTypeSpec.
	VisitPrimitiveTypeSpec(ctx *PrimitiveTypeSpecContext) interface{}

	// Visit a parse tree produced by NebularisParser#FunctionTypeSpec.
	VisitFunctionTypeSpec(ctx *FunctionTypeSpecContext) interface{}

	// Visit a parse tree produced by NebularisParser#StructTypeSpec.
	VisitStructTypeSpec(ctx *StructTypeSpecContext) interface{}

	// Visit a parse tree produced by NebularisParser#InterfaceTypeSpec.
	VisitInterfaceTypeSpec(ctx *InterfaceTypeSpecContext) interface{}

	// Visit a parse tree produced by NebularisParser#referencedType.
	VisitReferencedType(ctx *ReferencedTypeContext) interface{}

	// Visit a parse tree produced by NebularisParser#extendsClause.
	VisitExtendsClause(ctx *ExtendsClauseContext) interface{}

	// Visit a parse tree produced by NebularisParser#method.
	VisitMethod(ctx *MethodContext) interface{}

	// Visit a parse tree produced by NebularisParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by NebularisParser#typeParameters.
	VisitTypeParameters(ctx *TypeParametersContext) interface{}

	// Visit a parse tree produced by NebularisParser#typeParameter.
	VisitTypeParameter(ctx *TypeParameterContext) interface{}

	// Visit a parse tree produced by NebularisParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by NebularisParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by NebularisParser#TernaryExpression.
	VisitTernaryExpression(ctx *TernaryExpressionContext) interface{}

	// Visit a parse tree produced by NebularisParser#MemberAccessExpression.
	VisitMemberAccessExpression(ctx *MemberAccessExpressionContext) interface{}

	// Visit a parse tree produced by NebularisParser#ParenthesisExpression.
	VisitParenthesisExpression(ctx *ParenthesisExpressionContext) interface{}

	// Visit a parse tree produced by NebularisParser#LiteralExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by NebularisParser#IndexExpression.
	VisitIndexExpression(ctx *IndexExpressionContext) interface{}

	// Visit a parse tree produced by NebularisParser#UnaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by NebularisParser#InvocationExpression.
	VisitInvocationExpression(ctx *InvocationExpressionContext) interface{}

	// Visit a parse tree produced by NebularisParser#IdentifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by NebularisParser#StructExpression.
	VisitStructExpression(ctx *StructExpressionContext) interface{}

	// Visit a parse tree produced by NebularisParser#BinaryExpression.
	VisitBinaryExpression(ctx *BinaryExpressionContext) interface{}

	// Visit a parse tree produced by NebularisParser#ValueExpression.
	VisitValueExpression(ctx *ValueExpressionContext) interface{}

	// Visit a parse tree produced by NebularisParser#NotYetImplementedExpression.
	VisitNotYetImplementedExpression(ctx *NotYetImplementedExpressionContext) interface{}

	// Visit a parse tree produced by NebularisParser#LambdaExpression.
	VisitLambdaExpression(ctx *LambdaExpressionContext) interface{}

	// Visit a parse tree produced by NebularisParser#lambdaExpr.
	VisitLambdaExpr(ctx *LambdaExprContext) interface{}

	// Visit a parse tree produced by NebularisParser#structExpr.
	VisitStructExpr(ctx *StructExprContext) interface{}

	// Visit a parse tree produced by NebularisParser#fieldAssignments.
	VisitFieldAssignments(ctx *FieldAssignmentsContext) interface{}

	// Visit a parse tree produced by NebularisParser#fieldAssignment.
	VisitFieldAssignment(ctx *FieldAssignmentContext) interface{}

	// Visit a parse tree produced by NebularisParser#prefixOp.
	VisitPrefixOp(ctx *PrefixOpContext) interface{}

	// Visit a parse tree produced by NebularisParser#postfixOp.
	VisitPostfixOp(ctx *PostfixOpContext) interface{}

	// Visit a parse tree produced by NebularisParser#binaryOp.
	VisitBinaryOp(ctx *BinaryOpContext) interface{}

	// Visit a parse tree produced by NebularisParser#attribute.
	VisitAttribute(ctx *AttributeContext) interface{}

	// Visit a parse tree produced by NebularisParser#qualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) interface{}

	// Visit a parse tree produced by NebularisParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
