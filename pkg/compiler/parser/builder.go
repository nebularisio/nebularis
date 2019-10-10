//  Copyright 2019 Nebularis Authors.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package parser

import (
	"reflect"
	"strconv"

	"nebularis.io/nebularis/gen/parser"
	"nebularis.io/nebularis/pkg/compiler/ast"
	"nebularis.io/nebularis/pkg/compiler/common"
	"nebularis.io/nebularis/pkg/compiler/diag"
	"nebularis.io/nebularis/pkg/compiler/diag/msg"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type builder struct {
	ms *diag.Messages
	cu *ast.CompilationUnit
}

func (b *builder) compilationUnit(ctx parser.ICompilationUnitContext) {
	b.cu = &ast.CompilationUnit{}

	if ctx.GetModule() != nil {
		b.cu.Module = b.moduleDeclaration(ctx.GetModule())
	}

	for _, imp := range ctx.GetImports() {
		impDecl := b.importStatement(imp)
		b.cu.Imports = append(b.cu.Imports, impDecl)
	}

	for _, decl := range ctx.GetDeclarations() {
		if decl.GetTypeDecl() != nil {
			d := b.typeDeclaration(decl.GetTypeDecl())
			b.cu.Declarations = append(b.cu.Declarations, d)
		} else if decl.GetFnDecl() != nil {
			d := b.functionDeclaration(decl.GetFnDecl())
			b.cu.Declarations = append(b.cu.Declarations, d)
		} else {
			b.ms.Append(msg.Internal(b.loc(decl), "Unrecognized declaration in module: %v", reflect.TypeOf(decl)))
			return
		}
	}
}

func (b *builder) moduleDeclaration(ctx parser.IModuleDeclarationContext) *ast.ModuleDeclaration {
	return &ast.ModuleDeclaration{
		NodeBase:   ast.NodeBase{Location: b.loc(ctx)},
		Attributes: b.attributes(ctx.GetAttributes()),
		Name:       b.qualifiedName(ctx.GetName()),
	}
}

func (b *builder) attributes(ctxs []parser.IAttributeContext) ast.Attributes {
	var result ast.Attributes
	for _, ctx := range ctxs {
		result = append(result, b.attribute(ctx))
	}
	return result
}

func (b *builder) attribute(ctx parser.IAttributeContext) *ast.Attribute {
	return &ast.Attribute{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Value:    b.structExpression(ctx.GetValue()),
	}
}

func (b *builder) structExpression(ctx parser.IStructExprContext) *ast.StructExpression {
	return &ast.StructExpression{
		NodeBase:    ast.NodeBase{Location: b.loc(ctx)},
		Type:        b.typeReference(ctx.GetTypeRef()),
		Assignments: b.fieldAssignments(ctx.GetAssignments().GetFields()),
	}
}

func (b *builder) typeReference(ctx parser.IReferencedTypeContext) *ast.TypeReference {
	if ctx == nil {
		return nil
	}
	return &ast.TypeReference{
		NodeBase:      ast.NodeBase{Location: b.loc(ctx)},
		Name:          b.qualifiedName(ctx.GetName()),
		TypeArguments: b.typeSpecs(ctx.GetTypeArgs()),
	}
}

func (b *builder) typeSpecs(ctxs []parser.ITypeSpecContext) ast.TypeSpecs {
	var result ast.TypeSpecs
	for _, ctx := range ctxs {
		result = append(result, b.typeSpec(ctx))
	}
	return result
}

func (b *builder) typeSpec(ctx parser.ITypeSpecContext) ast.TypeSpec {
	if ctx == nil {
		return nil
	}
	switch v := ctx.(type) {
	case *parser.PrimitiveTypeSpecContext:
		return b.primitiveTypeSpec(v)
	case *parser.ReferencedTypeSpecContext:
		return b.referencedTypeSpec(v)
	case *parser.InterfaceTypeSpecContext:
		return b.interfaceTypeSpec(v)
	case *parser.NullableTypeSpecContext:
		return b.nullableTypSpec(v)
	case *parser.SpanTypeSpecContext:
		return b.spanTypeSpec(v)
	case *parser.StructTypeSpecContext:
		return b.structTypeSpec(v)
	case *parser.FunctionTypeSpecContext:
		return b.functionTypeSpec(v)
	default:
		b.ms.Append(msg.Internal(b.loc(v), "Unrecognized type: %v", reflect.TypeOf(v)))
		return nil
	}
}

func (b *builder) functionTypeSpec(ctx *parser.FunctionTypeSpecContext) *ast.FunctionTypeSpec {
	return &ast.FunctionTypeSpec{
		NodeBase:   ast.NodeBase{Location: b.loc(ctx)},
		Parameters: b.parameters(ctx.GetParams()),
		ReturnType: b.typeSpec(ctx.GetReturnType()),
	}
}

func (b *builder) structTypeSpec(ctx *parser.StructTypeSpecContext) *ast.StructType {
	return &ast.StructType{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Fields:   b.fields(ctx.GetFields()),
	}
}

func (b *builder) fields(ctxs []parser.IFieldContext) ast.Fields {
	var result ast.Fields
	for _, ctx := range ctxs {
		result = append(result, b.field(ctx))
	}
	return result
}

func (b *builder) field(ctx parser.IFieldContext) *ast.Field {
	return &ast.Field{
		Attributes:  b.attributes(ctx.GetAttributes()),
		Type:        b.typeSpec(ctx.GetTypeRef()),
		Name:        b.identifier(ctx.GetName()),
		Constraints: b.whereClauses(ctx.GetConstraints()),
	}
}

func (b *builder) spanTypeSpec(ctx *parser.SpanTypeSpecContext) *ast.SpanTypeSpec {
	return &ast.SpanTypeSpec{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Element:  b.typeSpec(ctx.TypeSpec()),
	}
}

func (b *builder) nullableTypSpec(ctx *parser.NullableTypeSpecContext) *ast.NullableTypeSpec {
	return &ast.NullableTypeSpec{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Element:  b.typeSpec(ctx.TypeSpec()),
	}
}

func (b *builder) primitiveTypeSpec(ctx *parser.PrimitiveTypeSpecContext) *ast.PrimitiveType {
	var p ast.PrimitiveKind

	if ctx.PrimitiveType().GetTBool() != nil {
		p = ast.Bool
	} else if ctx.PrimitiveType().GetTByte() != nil {
		p = ast.Byte
	} else if ctx.PrimitiveType().GetTChar() != nil {
		p = ast.Char
	} else if ctx.PrimitiveType().GetTInt32() != nil {
		p = ast.Int32
	} else if ctx.PrimitiveType().GetTStr() != nil {
		p = ast.String
	} else {
		b.ms.Append(msg.Internal(b.loc(ctx), "Unrecognized primitive: %v", ctx.GetText()))
	}

	return &ast.PrimitiveType{
		NodeBase:  ast.NodeBase{Location: b.loc(ctx)},
		Primitive: p,
	}
}

func (b *builder) interfaceTypeSpec(ctx *parser.InterfaceTypeSpecContext) *ast.InterfaceType {
	return &ast.InterfaceType{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Extends:  b.extends(ctx.GetExtends()),
		Methods:  b.methods(ctx.GetMethods()),
	}
}

func (b *builder) extends(ctxs parser.IExtendsClauseContext) ast.TypeReferences {
	var result ast.TypeReferences
	if ctxs != nil {
		for _, ctx := range ctxs.GetRefTypes() {
			result = append(result, b.typeReference(ctx))
		}
	}
	return result
}

func (b *builder) methods(ctxs []parser.IMethodContext) ast.Methods {
	var result ast.Methods
	for _, ctx := range ctxs {
		result = append(result, b.method(ctx))
	}
	return result
}

func (b *builder) method(ctx parser.IMethodContext) *ast.Method {
	return &ast.Method{
		FunctionTypeSpec: ast.FunctionTypeSpec{
			NodeBase:   ast.NodeBase{Location: b.loc(ctx)},
			Parameters: b.parameters(ctx.GetParams()),
			ReturnType: b.typeSpec(ctx.GetReturnType()),
		},
		Attributes:  b.attributes(ctx.GetAttributes()),
		Name:        b.identifier(ctx.GetName()),
		Constraints: b.whereClauses(ctx.GetConstraints()),
	}
}

func (b *builder) parameters(ctx parser.IFunctionParametersContext) ast.FunctionParameters {
	var result ast.FunctionParameters
	for _, c := range ctx.GetParams() {
		result = append(result, b.parameter(c))
	}
	return result
}

func (b *builder) parameter(ctx parser.IFunctionParameterContext) *ast.FunctionParameter {
	return &ast.FunctionParameter{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Name:     b.optIdentifer(ctx.GetName()),
		Type:     b.typeSpec(ctx.GetTypeRef()),
	}
}

func (b *builder) typeParameters(ctxs parser.ITypeParametersContext) ast.TypeParameters {
	var result ast.TypeParameters
	if ctxs != nil && ctxs.GetParams() != nil {
		for _, ctx := range ctxs.GetParams() {
			result = append(result, b.typeParameter(ctx))
		}
	}
	return result
}

func (b *builder) typeParameter(ctx parser.ITypeParameterContext) *ast.TypeParameter {
	return &ast.TypeParameter{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Name:     b.identifier(ctx.GetName()),
	}
}

func (b *builder) referencedTypeSpec(ctx *parser.ReferencedTypeSpecContext) *ast.TypeReference {
	return b.typeReference(ctx.GetRefType())
}

func (b *builder) fieldAssignments(ctxs []parser.IFieldAssignmentContext) ast.FieldAssignments {
	var result ast.FieldAssignments
	for _, ctx := range ctxs {
		result = append(result, b.fieldAssignment(ctx))
	}
	return result
}

func (b *builder) fieldAssignment(ctx parser.IFieldAssignmentContext) *ast.FieldAssignment {
	return &ast.FieldAssignment{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Name:     b.identifier(ctx.GetName()),
		Value:    b.expression(ctx.GetValue()),
	}
}

func (b *builder) expressions(ctxs []parser.IExpressionContext) ast.Expressions {
	var result ast.Expressions
	for _, ctx := range ctxs {
		result = append(result, b.expression(ctx))
	}
	return result
}

func (b *builder) expression(ctx parser.IExpressionContext) ast.Expression {
	if ctx == nil {
		return nil
	}

	switch v := ctx.(type) {
	case *parser.LiteralExpressionContext:
		return b.literalExpression(v)
	case *parser.InvocationExpressionContext:
		return b.invocationExpression(v)
	case *parser.MemberAccessExpressionContext:
		return b.memberAccessExpression(v)
	case *parser.IdentifierExpressionContext:
		return b.identifierExpression(v)
	case *parser.NotYetImplementedExpressionContext:
		return b.notYetImplementedExpression(v)
	case *parser.BinaryExpressionContext:
		return b.binaryExpression(v)
	case *parser.ValueExpressionContext:
		return b.valueExpression(v)
	case *parser.IndexExpressionContext:
		return b.indexExpression(v)
	case *parser.TernaryExpressionContext:
		return b.ternaryExpression(v)
	case *parser.ParenthesisExpressionContext:
		return b.parenthesisExpression(v)
	case *parser.UnaryExpressionContext:
		return b.unaryExpression(v)
	case *parser.StructExpressionContext:
		return b.structExpression(v.StructExpr())
	case *parser.LambdaExpressionContext:
		return b.lambdaExpression(v.LambdaExpr())
	default:
		b.ms.Append(msg.Internal(b.loc(ctx), "Unrecognized expression: %v", ctx.GetText()))
		return nil
	}
}

func (b *builder) lambdaExpression(ctx parser.ILambdaExprContext) *ast.LambdaExpression {
	return &ast.LambdaExpression{
		NodeBase:       ast.NodeBase{Location: b.loc(ctx)},
		TypeParameters: b.typeParameters(ctx.GetTypeParams()),
		Parameters:     b.parameters(ctx.GetParams()),
		ReturnType:     b.typeSpec(ctx.GetReturnType()),
		Constraints:    b.whereClauses(ctx.GetConstraints()),
		Body:           b.codeBlock(ctx.GetBody()),
	}
}

func (b *builder) unaryExpression(ctx *parser.UnaryExpressionContext) *ast.UnaryExpression {
	return &ast.UnaryExpression{
		NodeBase:   ast.NodeBase{Location: b.loc(ctx)},
		Op:         b.unaryPrefixOp(ctx.GetOp()),
		Expression: b.expression(ctx.Expression()),
	}
}

func (b *builder) unaryPrefixOp(op parser.IPrefixOpContext) ast.UnaryPrefixOp {
	return ast.UnaryPrefixOp(op.GetText())
}

func (b *builder) parenthesisExpression(ctx *parser.ParenthesisExpressionContext) *ast.ParenthesisExpression {
	return &ast.ParenthesisExpression{
		NodeBase:   ast.NodeBase{Location: b.loc(ctx)},
		Expression: b.expression(ctx.Expression()),
	}
}

func (b *builder) ternaryExpression(ctx *parser.TernaryExpressionContext) *ast.TernaryExpression {
	return &ast.TernaryExpression{
		NodeBase:  ast.NodeBase{Location: b.loc(ctx)},
		Condition: b.expression(ctx.GetPredicate()),
		OnTrue:    b.expression(ctx.GetOnTrue()),
		OnFalse:   b.expression(ctx.GetOnFalse()),
	}
}

func (b *builder) indexExpression(ctx *parser.IndexExpressionContext) *ast.IndexExpression {
	return &ast.IndexExpression{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Target:   b.expression(ctx.GetTarget()),
		Index:    b.expression(ctx.GetIndex()),
	}
}

func (b *builder) valueExpression(ctx *parser.ValueExpressionContext) *ast.ValueExpression {
	return &ast.ValueExpression{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
	}
}

func (b *builder) notYetImplementedExpression(ctx *parser.NotYetImplementedExpressionContext) *ast.NotYetImplementedExpression {
	return &ast.NotYetImplementedExpression{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
	}
}

func (b *builder) binaryExpression(ctx *parser.BinaryExpressionContext) *ast.BinaryExpression {
	return &ast.BinaryExpression{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Op:       b.binaryOp(ctx.BinaryOp()),
		Left:     b.expression(ctx.GetLeft()),
		Right:    b.expression(ctx.GetRight()),
	}
}

func (b *builder) binaryOp(ctx parser.IBinaryOpContext) ast.BinaryOp {
	return ast.BinaryOp(ctx.GetText())
}

func (b *builder) identifierExpression(ctx *parser.IdentifierExpressionContext) *ast.IdentifierExpression {
	return &ast.IdentifierExpression{
		NodeBase:   ast.NodeBase{Location: b.loc(ctx)},
		Identifier: b.identifier(ctx.GetIdent()),
	}
}

func (b *builder) memberAccessExpression(ctx *parser.MemberAccessExpressionContext) *ast.MemberAccessExpression {
	return &ast.MemberAccessExpression{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Target:   b.expression(ctx.GetTarget()),
		Member:   b.identifier(ctx.GetMember()),
	}
}

func (b *builder) invocationExpression(ctx *parser.InvocationExpressionContext) *ast.InvocationExpression {
	return &ast.InvocationExpression{
		NodeBase:  ast.NodeBase{Location: b.loc(ctx)},
		Target:    b.expression(ctx.GetTarget()),
		Arguments: b.expressions(ctx.GetArgs()),
	}
}

func (b *builder) literalExpression(ctx *parser.LiteralExpressionContext) *ast.LiteralExpression {
	var v ast.Literal
	if ctx.GetValue().GetBoolLit() != nil {
		v = &ast.BoolLiteral{
			Value: ctx.GetValue().GetBoolLit().GetText() == "true",
		}
	} else if ctx.GetValue().GetIntLit() != nil {
		i, err := strconv.Atoi(ctx.GetValue().GetIntLit().GetText())
		if err != nil {
			b.ms.Append(msg.Internal(b.loc(ctx), "Unparseable literal %q: %v", ctx.GetText(), err))
		}
		v = &ast.IntLiteral{Value: i}
	} else if ctx.GetValue().GetStringLit() != nil {
		v = &ast.StringLiteral{Value: ctx.GetValue().GetStringLit().GetText()}
	} else {
		b.ms.Append(msg.Internal(b.loc(ctx), "Unrecognized literal: %v", ctx.GetText()))
	}

	return &ast.LiteralExpression{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Value:    v,
	}
}

func (b *builder) importStatement(ctx parser.IImportStatementContext) *ast.ImportStatement {
	return &ast.ImportStatement{
		NodeBase:        ast.NodeBase{Location: b.loc(ctx)},
		Attributes:      b.attributes(ctx.GetAttributes()),
		ModuleReference: b.qualifiedName(ctx.GetModule().GetName()),
		Alias:           b.optIdentifer(ctx.GetAlias()),
	}
}

func (b *builder) typeDeclaration(ctx parser.ITypeDeclarationContext) *ast.TypeDeclaration {
	return &ast.TypeDeclaration{
		NodeBase:       ast.NodeBase{Location: b.loc(ctx)},
		Attributes:     b.attributes(ctx.GetAttributes()),
		Name:           b.identifier(ctx.GetName()),
		TypeParameters: b.typeParameters(ctx.GetTypeParams()),
		Spec:           b.typeSpec(ctx.GetSpec()),
		Constraints:    b.whereClauses(ctx.GetConstraints()),
	}
}

func (b *builder) functionDeclaration(ctx parser.IFunctionDeclarationContext) *ast.FunctionDeclaration {
	return &ast.FunctionDeclaration{
		FunctionTypeSpec: ast.FunctionTypeSpec{
			NodeBase: ast.NodeBase{
				Location: b.loc(ctx),
			},
			Parameters: b.parameters(ctx.GetParams()),
			ReturnType: b.typeSpec(ctx.GetReturnType()),
		},
		Attributes:     b.attributes(ctx.GetAttributes()),
		Public:         ctx.GetPub() != nil,
		Pure:           ctx.GetPure() != nil,
		Extern:         ctx.GetExtern() != nil,
		Target:         b.typeReference(ctx.GetTarget()),
		Name:           b.identifier(ctx.GetName()),
		TypeParameters: b.typeParameters(ctx.GetTypeParams()),
		Constraints:    b.whereClauses(ctx.GetConstraints()),
		Body:           b.codeBlock(ctx.GetBody()),
	}
}

func (b *builder) codeBlock(ctx parser.ICodeBlockContext) *ast.CodeBlock {
	if ctx == nil {
		return nil
	}
	return &ast.CodeBlock{
		NodeBase:   ast.NodeBase{Location: b.loc(ctx)},
		Statements: b.statements(ctx.GetStatements()),
	}
}

func (b *builder) statements(ctxs []parser.IStatementContext) ast.Statements {
	var result ast.Statements
	for _, ctx := range ctxs {
		result = append(result, b.statement(ctx))
	}
	return result
}

func (b *builder) statement(ctx parser.IStatementContext) ast.Statement {
	switch v := ctx.(type) {
	case *parser.VariableStatementContext:
		return b.variableStatement(v)
	case *parser.ReturnStatementContext:
		return b.returnStatement(v)
	case *parser.ConditionalStatementContext:
		return b.conditionalStatement(v)
	case *parser.WhileStatementContext:
		return b.whileStatement(v)
	case *parser.ExpressionStatementContext:
		return b.expressionStatement(v)
	case *parser.NotYetImplementedStatementContext:
		return b.notYetImplementedStatement(v)
	case *parser.UnaryStatementContext:
		return b.unaryStatement(v)
	case *parser.BreakStatementContext:
		return b.breakStatement(v)
	case *parser.LabelStatementContext:
		return b.labelStatement(v)
	case *parser.ContinueStatementContext:
		return b.continueStatement(v)
	case *parser.AssignmentStatementContext:
		return b.assignmentStatement(v)
	case *parser.SwitchStatementContext:
		return b.switchStatement(v)
	default:
		b.ms.Append(msg.Internal(b.loc(ctx), "Unrecognized statement: %v", ctx.GetText()))
		return nil
	}
}

func (b *builder) switchStatement(ctx *parser.SwitchStatementContext) *ast.SwitchStatement {
	return &ast.SwitchStatement{
		NodeBase:   ast.NodeBase{Location: b.loc(ctx)},
		Expression: b.expression(ctx.Expression()),
		Cases:      b.caseBlocks(ctx.GetCases()),
		Default:    b.defaultBlock(ctx.GetDefauls()),
	}
}

func (b *builder) caseBlocks(ctxs []parser.ICaseStatementContext) ast.CaseBlocks {
	var result ast.CaseBlocks
	for _, ctx := range ctxs {
		result = append(result, b.caseBlock(ctx))
	}
	return result
}

func (b *builder) caseBlock(ctx parser.ICaseStatementContext) *ast.CaseBlock {
	body := b.codeBlock(ctx.GetBlock())
	if body == nil {
		body = &ast.CodeBlock{
			NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		}
		for _, s := range ctx.GetStatements() {
			body.Statements = append(body.Statements, b.statement(s))
		}
		body.Inline = true
	}

	return &ast.CaseBlock{
		NodeBase:    ast.NodeBase{Location: b.loc(ctx)},
		Expressions: b.expressions(ctx.GetExpr()),
		Body:        body,
	}
}

func (b *builder) defaultBlock(ctxs []parser.IDefaultStatementContext) *ast.DefaultBlock {
	switch len(ctxs) {
	case 0:
		return nil
	case 1:
		body := b.codeBlock(ctxs[0].GetBlock())
		if body == nil {
			body = &ast.CodeBlock{
				NodeBase: ast.NodeBase{Location: b.loc(ctxs[0])},
			}
			for _, s := range ctxs[0].GetStatements() {
				body.Statements = append(body.Statements, b.statement(s))
			}
			body.Inline = true
		}

		return &ast.DefaultBlock{
			NodeBase: ast.NodeBase{Location: b.loc(ctxs[0])},
			Body:     body,
		}
	default:
		b.ms.Append(msg.MultipleDefaults(b.loc(ctxs[1])))
		return nil
	}
}

func (b *builder) assignmentStatement(ctx *parser.AssignmentStatementContext) *ast.AssignmentStatement {
	return &ast.AssignmentStatement{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Left:     b.expression(ctx.GetLeft()),
		Right:    b.expression(ctx.GetRight()),
	}
}

func (b *builder) continueStatement(ctx *parser.ContinueStatementContext) *ast.ContinueStatement {
	return &ast.ContinueStatement{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Label:    b.optIdentifer(ctx.GetLabel()),
	}
}

func (b *builder) labelStatement(ctx *parser.LabelStatementContext) *ast.LabelStatement {
	return &ast.LabelStatement{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Label:    b.identifier(ctx.GetLabel()),
	}
}

func (b *builder) breakStatement(ctx *parser.BreakStatementContext) *ast.BreakStatement {
	return &ast.BreakStatement{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Label:    b.optIdentifer(ctx.GetLabel()),
	}
}

func (b *builder) unaryStatement(ctx *parser.UnaryStatementContext) *ast.UnaryStatement {
	return &ast.UnaryStatement{
		NodeBase:   ast.NodeBase{Location: b.loc(ctx)},
		Expression: b.expression(ctx.Expression()),
		Op:         b.unaryPostfix(ctx.GetOp()),
	}
}

func (b *builder) unaryPostfix(op parser.IPostfixOpContext) ast.UnaryPostFixOp {
	return ast.UnaryPostFixOp(op.GetText())
}

func (b *builder) whileStatement(ctx *parser.WhileStatementContext) *ast.WhileStatement {
	return &ast.WhileStatement{
		NodeBase:  ast.NodeBase{Location: b.loc(ctx)},
		Condition: b.expression(ctx.Expression()),
		Body:      b.codeBlock(ctx.CodeBlock()),
	}
}

func (b *builder) conditionalStatement(ctx *parser.ConditionalStatementContext) *ast.ConditionalStatement {
	return &ast.ConditionalStatement{
		NodeBase:  ast.NodeBase{Location: b.loc(ctx)},
		Condition: b.expression(ctx.Expression()),
		OnTrue:    b.codeBlock(ctx.GetOnTrue()),
		OnFalse:   b.codeBlock(ctx.GetOnFalse()),
	}
}

func (b *builder) variableStatement(ctx *parser.VariableStatementContext) *ast.VariableStatement {
	return &ast.VariableStatement{
		NodeBase:    ast.NodeBase{Location: b.loc(ctx)},
		Name:        b.identifier(ctx.GetName()),
		Type:        b.typeSpec(ctx.GetTypeRef()),
		Initializer: b.expression(ctx.GetInitializer()),
	}
}

func (b *builder) returnStatement(ctx *parser.ReturnStatementContext) *ast.ReturnStatement {
	return &ast.ReturnStatement{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
		Value:    b.expression(ctx.GetValue()),
	}
}

func (b *builder) notYetImplementedStatement(ctx *parser.NotYetImplementedStatementContext) *ast.NotYetImplementedStatement {
	return &ast.NotYetImplementedStatement{
		NodeBase: ast.NodeBase{Location: b.loc(ctx)},
	}
}

func (b *builder) expressionStatement(ctx *parser.ExpressionStatementContext) *ast.ExpressionStatement {
	return &ast.ExpressionStatement{
		NodeBase:   ast.NodeBase{Location: b.loc(ctx)},
		Expression: b.expression(ctx.Expression()),
	}
}

func (b *builder) whereClauses(ctxs []parser.IWhereClauseContext) ast.WhereClauses {
	var result ast.WhereClauses
	for _, ctx := range ctxs {
		result = append(result, b.whereClause(ctx))
	}

	return result
}

func (b *builder) whereClause(ctx parser.IWhereClauseContext) *ast.WhereClause {
	return &ast.WhereClause{
		NodeBase:   ast.NodeBase{Location: b.loc(ctx)},
		Expression: b.expression(ctx.GetExpr()),
	}
}

func (b *builder) qualifiedName(ctx parser.IQualifiedNameContext) common.QualifiedName {
	return common.QualifiedName(ctx.GetText())
}

func (b *builder) optIdentifer(token antlr.Token) *common.Identifier {
	if token == nil {
		return nil
	}
	i := b.identifier(token)
	return &i
}

func (b *builder) identifier(token antlr.Token) common.Identifier {
	return common.Identifier(token.GetText())
}

func (b *builder) loc(ctx antlr.ParserRuleContext) common.Location {
	return b.locToken(ctx.GetStart())
}

func (b *builder) locToken(t antlr.Token) common.Location {
	l := t.GetLine()
	c := t.GetColumn()

	return common.Location{Column: int32(c), Line: int32(l)}
}
