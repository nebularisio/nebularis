// Code generated from /Users/ozbenevren/dev/go/src/nebularis.io/nebularis/Nebularis.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Nebularis

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseNebularisVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseNebularisVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitModuleDeclaration(ctx *ModuleDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitModuleReference(ctx *ModuleReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitFunctionParameters(ctx *FunctionParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitFunctionParameter(ctx *FunctionParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitCodeBlock(ctx *CodeBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitVariableStatement(ctx *VariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitConditionalStatement(ctx *ConditionalStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitLabelStatement(ctx *LabelStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitUnaryStatement(ctx *UnaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitNotYetImplementedStatement(ctx *NotYetImplementedStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitCaseStatement(ctx *CaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitDefaultStatement(ctx *DefaultStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitNullableTypeSpec(ctx *NullableTypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitSpanTypeSpec(ctx *SpanTypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitReferencedTypeSpec(ctx *ReferencedTypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitPrimitiveTypeSpec(ctx *PrimitiveTypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitFunctionTypeSpec(ctx *FunctionTypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitStructTypeSpec(ctx *StructTypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitInterfaceTypeSpec(ctx *InterfaceTypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitReferencedType(ctx *ReferencedTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitExtendsClause(ctx *ExtendsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitMethod(ctx *MethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitTypeParameters(ctx *TypeParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitTernaryExpression(ctx *TernaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitMemberAccessExpression(ctx *MemberAccessExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitParenthesisExpression(ctx *ParenthesisExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitIndexExpression(ctx *IndexExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitInvocationExpression(ctx *InvocationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitStructExpression(ctx *StructExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitValueExpression(ctx *ValueExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitNotYetImplementedExpression(ctx *NotYetImplementedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitLambdaExpression(ctx *LambdaExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitLambdaExpr(ctx *LambdaExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitStructExpr(ctx *StructExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitFieldAssignments(ctx *FieldAssignmentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitFieldAssignment(ctx *FieldAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitPrefixOp(ctx *PrefixOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitPostfixOp(ctx *PostfixOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitBinaryOp(ctx *BinaryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitAttribute(ctx *AttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNebularisVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
