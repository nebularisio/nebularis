// Code generated from /Users/ozbenevren/dev/go/src/nebularis.io/nebularis/Nebularis.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Nebularis

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NebularisListener is a complete listener for a parse tree produced by NebularisParser.
type NebularisListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterModuleDeclaration is called when entering the moduleDeclaration production.
	EnterModuleDeclaration(c *ModuleDeclarationContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterModuleReference is called when entering the moduleReference production.
	EnterModuleReference(c *ModuleReferenceContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionParameters is called when entering the functionParameters production.
	EnterFunctionParameters(c *FunctionParametersContext)

	// EnterFunctionParameter is called when entering the functionParameter production.
	EnterFunctionParameter(c *FunctionParameterContext)

	// EnterCodeBlock is called when entering the codeBlock production.
	EnterCodeBlock(c *CodeBlockContext)

	// EnterVariableStatement is called when entering the VariableStatement production.
	EnterVariableStatement(c *VariableStatementContext)

	// EnterReturnStatement is called when entering the ReturnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterConditionalStatement is called when entering the ConditionalStatement production.
	EnterConditionalStatement(c *ConditionalStatementContext)

	// EnterWhileStatement is called when entering the WhileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterLabelStatement is called when entering the LabelStatement production.
	EnterLabelStatement(c *LabelStatementContext)

	// EnterContinueStatement is called when entering the ContinueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterBreakStatement is called when entering the BreakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterSwitchStatement is called when entering the SwitchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterUnaryStatement is called when entering the UnaryStatement production.
	EnterUnaryStatement(c *UnaryStatementContext)

	// EnterAssignmentStatement is called when entering the AssignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterNotYetImplementedStatement is called when entering the NotYetImplementedStatement production.
	EnterNotYetImplementedStatement(c *NotYetImplementedStatementContext)

	// EnterExpressionStatement is called when entering the ExpressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterCaseStatement is called when entering the caseStatement production.
	EnterCaseStatement(c *CaseStatementContext)

	// EnterDefaultStatement is called when entering the defaultStatement production.
	EnterDefaultStatement(c *DefaultStatementContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterNullableTypeSpec is called when entering the NullableTypeSpec production.
	EnterNullableTypeSpec(c *NullableTypeSpecContext)

	// EnterSpanTypeSpec is called when entering the SpanTypeSpec production.
	EnterSpanTypeSpec(c *SpanTypeSpecContext)

	// EnterReferencedTypeSpec is called when entering the ReferencedTypeSpec production.
	EnterReferencedTypeSpec(c *ReferencedTypeSpecContext)

	// EnterPrimitiveTypeSpec is called when entering the PrimitiveTypeSpec production.
	EnterPrimitiveTypeSpec(c *PrimitiveTypeSpecContext)

	// EnterFunctionTypeSpec is called when entering the FunctionTypeSpec production.
	EnterFunctionTypeSpec(c *FunctionTypeSpecContext)

	// EnterStructTypeSpec is called when entering the StructTypeSpec production.
	EnterStructTypeSpec(c *StructTypeSpecContext)

	// EnterInterfaceTypeSpec is called when entering the InterfaceTypeSpec production.
	EnterInterfaceTypeSpec(c *InterfaceTypeSpecContext)

	// EnterReferencedType is called when entering the referencedType production.
	EnterReferencedType(c *ReferencedTypeContext)

	// EnterExtendsClause is called when entering the extendsClause production.
	EnterExtendsClause(c *ExtendsClauseContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterTypeParameters is called when entering the typeParameters production.
	EnterTypeParameters(c *TypeParametersContext)

	// EnterTypeParameter is called when entering the typeParameter production.
	EnterTypeParameter(c *TypeParameterContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterTernaryExpression is called when entering the TernaryExpression production.
	EnterTernaryExpression(c *TernaryExpressionContext)

	// EnterMemberAccessExpression is called when entering the MemberAccessExpression production.
	EnterMemberAccessExpression(c *MemberAccessExpressionContext)

	// EnterParenthesisExpression is called when entering the ParenthesisExpression production.
	EnterParenthesisExpression(c *ParenthesisExpressionContext)

	// EnterLiteralExpression is called when entering the LiteralExpression production.
	EnterLiteralExpression(c *LiteralExpressionContext)

	// EnterIndexExpression is called when entering the IndexExpression production.
	EnterIndexExpression(c *IndexExpressionContext)

	// EnterUnaryExpression is called when entering the UnaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterInvocationExpression is called when entering the InvocationExpression production.
	EnterInvocationExpression(c *InvocationExpressionContext)

	// EnterIdentifierExpression is called when entering the IdentifierExpression production.
	EnterIdentifierExpression(c *IdentifierExpressionContext)

	// EnterStructExpression is called when entering the StructExpression production.
	EnterStructExpression(c *StructExpressionContext)

	// EnterBinaryExpression is called when entering the BinaryExpression production.
	EnterBinaryExpression(c *BinaryExpressionContext)

	// EnterValueExpression is called when entering the ValueExpression production.
	EnterValueExpression(c *ValueExpressionContext)

	// EnterNotYetImplementedExpression is called when entering the NotYetImplementedExpression production.
	EnterNotYetImplementedExpression(c *NotYetImplementedExpressionContext)

	// EnterLambdaExpression is called when entering the LambdaExpression production.
	EnterLambdaExpression(c *LambdaExpressionContext)

	// EnterLambdaExpr is called when entering the lambdaExpr production.
	EnterLambdaExpr(c *LambdaExprContext)

	// EnterStructExpr is called when entering the structExpr production.
	EnterStructExpr(c *StructExprContext)

	// EnterFieldAssignments is called when entering the fieldAssignments production.
	EnterFieldAssignments(c *FieldAssignmentsContext)

	// EnterFieldAssignment is called when entering the fieldAssignment production.
	EnterFieldAssignment(c *FieldAssignmentContext)

	// EnterPrefixOp is called when entering the prefixOp production.
	EnterPrefixOp(c *PrefixOpContext)

	// EnterPostfixOp is called when entering the postfixOp production.
	EnterPostfixOp(c *PostfixOpContext)

	// EnterBinaryOp is called when entering the binaryOp production.
	EnterBinaryOp(c *BinaryOpContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterQualifiedName is called when entering the qualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitModuleDeclaration is called when exiting the moduleDeclaration production.
	ExitModuleDeclaration(c *ModuleDeclarationContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitModuleReference is called when exiting the moduleReference production.
	ExitModuleReference(c *ModuleReferenceContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionParameters is called when exiting the functionParameters production.
	ExitFunctionParameters(c *FunctionParametersContext)

	// ExitFunctionParameter is called when exiting the functionParameter production.
	ExitFunctionParameter(c *FunctionParameterContext)

	// ExitCodeBlock is called when exiting the codeBlock production.
	ExitCodeBlock(c *CodeBlockContext)

	// ExitVariableStatement is called when exiting the VariableStatement production.
	ExitVariableStatement(c *VariableStatementContext)

	// ExitReturnStatement is called when exiting the ReturnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitConditionalStatement is called when exiting the ConditionalStatement production.
	ExitConditionalStatement(c *ConditionalStatementContext)

	// ExitWhileStatement is called when exiting the WhileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitLabelStatement is called when exiting the LabelStatement production.
	ExitLabelStatement(c *LabelStatementContext)

	// ExitContinueStatement is called when exiting the ContinueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitBreakStatement is called when exiting the BreakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitSwitchStatement is called when exiting the SwitchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitUnaryStatement is called when exiting the UnaryStatement production.
	ExitUnaryStatement(c *UnaryStatementContext)

	// ExitAssignmentStatement is called when exiting the AssignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitNotYetImplementedStatement is called when exiting the NotYetImplementedStatement production.
	ExitNotYetImplementedStatement(c *NotYetImplementedStatementContext)

	// ExitExpressionStatement is called when exiting the ExpressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitCaseStatement is called when exiting the caseStatement production.
	ExitCaseStatement(c *CaseStatementContext)

	// ExitDefaultStatement is called when exiting the defaultStatement production.
	ExitDefaultStatement(c *DefaultStatementContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitNullableTypeSpec is called when exiting the NullableTypeSpec production.
	ExitNullableTypeSpec(c *NullableTypeSpecContext)

	// ExitSpanTypeSpec is called when exiting the SpanTypeSpec production.
	ExitSpanTypeSpec(c *SpanTypeSpecContext)

	// ExitReferencedTypeSpec is called when exiting the ReferencedTypeSpec production.
	ExitReferencedTypeSpec(c *ReferencedTypeSpecContext)

	// ExitPrimitiveTypeSpec is called when exiting the PrimitiveTypeSpec production.
	ExitPrimitiveTypeSpec(c *PrimitiveTypeSpecContext)

	// ExitFunctionTypeSpec is called when exiting the FunctionTypeSpec production.
	ExitFunctionTypeSpec(c *FunctionTypeSpecContext)

	// ExitStructTypeSpec is called when exiting the StructTypeSpec production.
	ExitStructTypeSpec(c *StructTypeSpecContext)

	// ExitInterfaceTypeSpec is called when exiting the InterfaceTypeSpec production.
	ExitInterfaceTypeSpec(c *InterfaceTypeSpecContext)

	// ExitReferencedType is called when exiting the referencedType production.
	ExitReferencedType(c *ReferencedTypeContext)

	// ExitExtendsClause is called when exiting the extendsClause production.
	ExitExtendsClause(c *ExtendsClauseContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitTypeParameters is called when exiting the typeParameters production.
	ExitTypeParameters(c *TypeParametersContext)

	// ExitTypeParameter is called when exiting the typeParameter production.
	ExitTypeParameter(c *TypeParameterContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitTernaryExpression is called when exiting the TernaryExpression production.
	ExitTernaryExpression(c *TernaryExpressionContext)

	// ExitMemberAccessExpression is called when exiting the MemberAccessExpression production.
	ExitMemberAccessExpression(c *MemberAccessExpressionContext)

	// ExitParenthesisExpression is called when exiting the ParenthesisExpression production.
	ExitParenthesisExpression(c *ParenthesisExpressionContext)

	// ExitLiteralExpression is called when exiting the LiteralExpression production.
	ExitLiteralExpression(c *LiteralExpressionContext)

	// ExitIndexExpression is called when exiting the IndexExpression production.
	ExitIndexExpression(c *IndexExpressionContext)

	// ExitUnaryExpression is called when exiting the UnaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitInvocationExpression is called when exiting the InvocationExpression production.
	ExitInvocationExpression(c *InvocationExpressionContext)

	// ExitIdentifierExpression is called when exiting the IdentifierExpression production.
	ExitIdentifierExpression(c *IdentifierExpressionContext)

	// ExitStructExpression is called when exiting the StructExpression production.
	ExitStructExpression(c *StructExpressionContext)

	// ExitBinaryExpression is called when exiting the BinaryExpression production.
	ExitBinaryExpression(c *BinaryExpressionContext)

	// ExitValueExpression is called when exiting the ValueExpression production.
	ExitValueExpression(c *ValueExpressionContext)

	// ExitNotYetImplementedExpression is called when exiting the NotYetImplementedExpression production.
	ExitNotYetImplementedExpression(c *NotYetImplementedExpressionContext)

	// ExitLambdaExpression is called when exiting the LambdaExpression production.
	ExitLambdaExpression(c *LambdaExpressionContext)

	// ExitLambdaExpr is called when exiting the lambdaExpr production.
	ExitLambdaExpr(c *LambdaExprContext)

	// ExitStructExpr is called when exiting the structExpr production.
	ExitStructExpr(c *StructExprContext)

	// ExitFieldAssignments is called when exiting the fieldAssignments production.
	ExitFieldAssignments(c *FieldAssignmentsContext)

	// ExitFieldAssignment is called when exiting the fieldAssignment production.
	ExitFieldAssignment(c *FieldAssignmentContext)

	// ExitPrefixOp is called when exiting the prefixOp production.
	ExitPrefixOp(c *PrefixOpContext)

	// ExitPostfixOp is called when exiting the postfixOp production.
	ExitPostfixOp(c *PostfixOpContext)

	// ExitBinaryOp is called when exiting the binaryOp production.
	ExitBinaryOp(c *BinaryOpContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitQualifiedName is called when exiting the qualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
