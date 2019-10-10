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

package analyzer

import (
	"nebularis.io/nebularis/pkg/compiler/ast"
	"nebularis.io/nebularis/pkg/compiler/common"
	"nebularis.io/nebularis/pkg/compiler/diag/msg"
	"nebularis.io/nebularis/pkg/compiler/semantics"
)

func (c *context) create() {

	// First create modules
	c.createModules()

	// Don't create cross-module references. Fail if any of the references cannot be resolved.
	c.createAndResolveModuleReferences()
	if c.m.HasErrors() {
		return
	}

	// Create top-level declarations, before starting reference resolution and analysis
	c.createDeclarations()
	if c.m.HasErrors() {
		return
	}
}

func (c *context) createDeclarations() {
	for _, cu := range c.sources {
		m := c.moduleFor(cu.Module)

		for _, decl := range cu.Declarations {
			switch decl.Kind() {
			case ast.DeclFunction:
				c.createFunctionDeclaration(m, decl.(*ast.FunctionDeclaration))

			case ast.DeclType:
				c.createTypeDeclaration(m, decl.(*ast.TypeDeclaration))

			default:
				c.m.Append(msg.Internal(decl.DeclLocation(), "Unrecognized decl: %v", decl.Kind()))
			}
		}
	}
}

func (c *context) createFunctionDeclaration(m *semantics.Module, decl *ast.FunctionDeclaration) {
	if _, found := m.Functions[decl.LocalName()]; found {
		c.m.Append(msg.MemberAlreadyDeclared(decl.Loc(), decl.LocalName()))
		return
	}
	// TODO: extern, public, pure

	parameters := c.createParameters(decl.Parameters)

	fn := &semantics.Function{
		Attributes: c.createAttributes(decl.Attributes),
		Parameters: parameters,
		Body:       c.createCodeBlock(decl.Body),
	}
	m.Functions[decl.LocalName()] = fn
}

func (c *context) createTypeDeclaration(m *semantics.Module, decl *ast.TypeDeclaration) {
	if _, found := m.Types[decl.LocalName()]; found {
		c.m.Append(msg.MemberAlreadyDeclared(decl.Loc(), decl.LocalName()))
		return
	}

	base := c.createTypeBase(decl.Attributes, decl.TypeParameters, decl.Constraints)

	var t semantics.Type
	switch decl.Spec.Kind() {
	case ast.Primitive:
		t = c.createPrimitiveType(decl.Spec.(*ast.PrimitiveType), base)

	case ast.Nullable:
		t = c.createNullableType(decl.Spec.(*ast.NullableTypeSpec), base)

	case ast.Span:
		t = c.createSpanType(decl.Spec.(*ast.SpanTypeSpec), base)

	case ast.Struct:
		t = c.createStructType(decl.Spec.(*ast.StructType), base)

	case ast.Interface:
		t = c.createInterfaceType(decl.Spec.(*ast.InterfaceType), base)

	case ast.Function:
		t = c.createFunctionType(decl.Spec.(*ast.FunctionTypeSpec), base)

	case ast.TypeRef:
		t = c.createReferencedType(decl.Spec.(*ast.TypeReference), base)

	default:
		c.m.Append(msg.Internal(decl.Spec.Loc(), "Unrecognized type: %v", decl.Spec.Kind()))
	}
	m.Types[decl.LocalName()] = t
}

func (c *context) createPrimitiveType(p *ast.PrimitiveType, base semantics.TypeBase) *semantics.PrimitiveType {
	return &semantics.PrimitiveType{
		TypeBase:      base,
		PrimitiveKind: p.Primitive,
	}
}

func (c *context) createType(s ast.TypeSpec, a ast.Attributes, p ast.TypeParameters, w ast.WhereClauses) semantics.Type {
	base := c.createTypeBase(a, p, w)

	switch s.Kind() {
	case ast.Primitive:
		return c.createPrimitiveType(s.(*ast.PrimitiveType), base)

	case ast.Nullable:
		return c.createNullableType(s.(*ast.NullableTypeSpec), base)

	case ast.Span:
		return c.createSpanType(s.(*ast.SpanTypeSpec), base)

	case ast.Struct:
		return c.createStructType(s.(*ast.StructType), base)

	case ast.Interface:
		return c.createInterfaceType(s.(*ast.InterfaceType), base)

	case ast.Function:
		return c.createFunctionType(s.(*ast.FunctionTypeSpec), base)

	case ast.TypeRef:
		return c.createReferencedType(s.(*ast.TypeReference), base)

	default:
		c.m.Append(msg.Internal(s.Loc(), "Unrecognized type: %v", s.Kind()))
		return nil
	}
}

func (c *context) createNullableType(spec *ast.NullableTypeSpec, base semantics.TypeBase) *semantics.NullableType {
	return &semantics.NullableType{
		TypeBase:  base,
		InnerType: c.createType(spec.Element, nil, spec.TypeParameters, nil),
	}
}

func (c *context) createSpanType(spec *ast.SpanTypeSpec, base semantics.TypeBase) *semantics.SpanType {
	return &semantics.SpanType{
		TypeBase:  base,
		InnerType: c.createType(spec.Element, nil, spec.TypeParameters, nil),
	}
}

func (c *context) createStructType(spec *ast.StructType, base semantics.TypeBase) *semantics.StructType {
	fields := make(map[common.Identifier]*semantics.Field)

	for _, f := range spec.Fields {
		if _, found := fields[f.Name]; found {
			c.m.Append(msg.FieldAlreadyDeclared(f.Location, f.Name))
			continue
		}

		ft := c.createType(f.Type, nil, nil, f.Constraints)

		field := &semantics.Field{
			Name:       f.Name,
			Type:       ft,
			Attributes: c.createAttributes(f.Attributes),
		}

		fields[f.Name] = field
	}

	return &semantics.StructType{
		TypeBase: base,
		Fields:   fields,
	}
}

func (c *context) createInterfaceType(spec *ast.InterfaceType, base semantics.TypeBase) *semantics.InterfaceType {

	var baseTypes []*semantics.ReferencedType
	for _, baseSpec := range spec.Extends {
		baseType := c.createReferencedType(baseSpec, semantics.TypeBase{})
		baseTypes = append(baseTypes, baseType)
	}

	methods := make(map[common.Identifier]*semantics.Method)
	for _, m := range spec.Methods {

		if _, found := methods[m.Name]; found {
			c.m.Append(msg.MethodAlreadyDeclared(m.Location, m.Name))
			continue
		}

		parameters := c.createParameters(m.Parameters)

		b := c.createTypeBase(m.Attributes, nil, m.Constraints)
		t := c.createFunctionType(&m.FunctionTypeSpec, b)

		method := &semantics.Method{
			Name:       m.Name,
			Parameters: parameters,
			Type:       t,
		}

		methods[m.Name] = method
	}

	return &semantics.InterfaceType{
		TypeBase:  base,
		BaseTypes: baseTypes,
		Methods:   methods,
	}
}

func (c *context) createFunctionType(spec *ast.FunctionTypeSpec, base semantics.TypeBase) *semantics.FunctionType {
	var types semantics.Types
	for _, p := range spec.Parameters {
		pt := c.createType(p.Type, nil, nil, nil)
		types = append(types, pt)
	}

	var returnType semantics.Type = &semantics.VoidType{}
	if spec.ReturnType != nil {
		returnType = c.createType(spec.ReturnType, nil, nil, nil)
	}

	var target semantics.Type
	if spec.Target != nil {
		target = c.createReferencedType(spec.Target, semantics.TypeBase{})
	}

	return &semantics.FunctionType{
		TypeBase:       base,
		Target:         target,
		ParameterTypes: types,
		ReturnType:     returnType,
	}
}

func (c *context) createReferencedType(spec *ast.TypeReference, base semantics.TypeBase) *semantics.ReferencedType {
	var args semantics.Types
	for _, a := range spec.TypeArguments {
		t := c.createType(a, nil, nil, nil)
		args = append(args, t)
	}

	return &semantics.ReferencedType{
		TypeBase:      base,
		Name:          spec.Name,
		TypeArguments: args,
	}
}

func (c *context) createParameters(ps ast.FunctionParameters) semantics.Parameters {
	var result semantics.Parameters
	for _, p := range ps {
		parameter := &semantics.Parameter{
			Name: identOrEmpty(p.Name),
			Type: c.createType(p.Type, nil, nil, nil),
		}
		result = append(result, parameter)
	}
	return result
}

func (c *context) createTypeBase(a ast.Attributes, t ast.TypeParameters, w ast.WhereClauses) semantics.TypeBase {
	return semantics.TypeBase{
		Attr:    c.createAttributes(a),
		Constr:  c.createConstraints(w),
		TypePar: c.createTypeParameters(t),
	}
}

func (c *context) createConstraints(cls ast.WhereClauses) semantics.WhereClauses {
	var result semantics.WhereClauses
	for _, cl := range cls {
		constraint := &semantics.WhereClause{
			Expression: c.createExpression(cl.Expression),
		}

		result = append(result, constraint)
	}
	return result
}

func (c *context) createTypeParameters(tps ast.TypeParameters) semantics.TypeParameters {
	var result semantics.TypeParameters
	for _, tp := range tps {
		param := &semantics.TypeParameter{
			Name: tp.Name,
		}
		result = append(result, param)
	}
	return result
}

func (c *context) createAttributes(attrs ast.Attributes) semantics.Attributes {
	var result semantics.Attributes
	for _, attr := range attrs {
		attribute := &semantics.Attribute{
			Value: c.createStructExpression(attr.Value),
		}
		result = append(result, attribute)
	}
	return result
}

func (c *context) createCodeBlock(s *ast.CodeBlock) semantics.CodeBlock {
	var result semantics.CodeBlock

	for _, s := range s.Statements {
		result = append(result, c.createStatement(s))
	}

	return result
}

func (c *context) createStatement(s ast.Statement) semantics.Statement {
	switch s.Kind() {
	case ast.StAssigment:
		return c.createAssignmentStatement(s.(*ast.AssignmentStatement))
	case ast.StBreak:
		return c.createBreakStatement(s.(*ast.BreakStatement))
	case ast.StContinue:
		return c.createContinueStatement(s.(*ast.ContinueStatement))
	case ast.StExpression:
		return c.createExpressionStatement(s.(*ast.ExpressionStatement))
	case ast.StIf:
		return c.createConditionalStatement(s.(*ast.ConditionalStatement))
	case ast.StLabel:
		return c.createLabelStatement(s.(*ast.LabelStatement))
	case ast.StNotYetImplemented:
		return c.createNotYetImplementedStatement(s.(*ast.NotYetImplementedStatement))
	case ast.StReturn:
		return c.createReturnStatement(s.(*ast.ReturnStatement))
	case ast.StSwitch:
		return c.createSwitchStatement(s.(*ast.SwitchStatement))
	case ast.StUnary:
		return c.createUnaryStatement(s.(*ast.UnaryStatement))
	case ast.StVariable:
		return c.createVariableStatement(s.(*ast.VariableStatement))
	case ast.StWhile:
		return c.createWhileStatement(s.(*ast.WhileStatement))

	default:
		c.m.Append(msg.Internal(s.Loc(), "Unrecognized statement: %v", s.Kind()))
		return nil
	}
}

func (c *context) createAssignmentStatement(statement *ast.AssignmentStatement) *semantics.AssignmentStatement {
	return &semantics.AssignmentStatement{
		Left:  c.createExpression(statement.Left),
		Right: c.createExpression(statement.Right),
	}
}

func (c *context) createBreakStatement(statement *ast.BreakStatement) *semantics.BreakStatement {
	return &semantics.BreakStatement{
		Label: identOrEmpty(statement.Label),
	}
}

func (c *context) createContinueStatement(statement *ast.ContinueStatement) *semantics.ContinueStatement {
	return &semantics.ContinueStatement{
		Label: identOrEmpty(statement.Label),
	}
}

func (c *context) createExpressionStatement(statement *ast.ExpressionStatement) *semantics.ExpressionStatement {
	return &semantics.ExpressionStatement{
		Expression: c.createExpression(statement.Expression),
	}
}

func (c *context) createConditionalStatement(statement *ast.ConditionalStatement) *semantics.ConditionalStatement {
	return &semantics.ConditionalStatement{
		Condition: c.createExpression(statement.Condition),
		OnTrue:    c.createCodeBlock(statement.OnTrue),
		OnFalse:   c.createCodeBlock(statement.OnFalse),
	}
}

func (c *context) createLabelStatement(statement *ast.LabelStatement) *semantics.LabelStatement {
	return &semantics.LabelStatement{
		Label: statement.Label,
	}
}

func (c *context) createNotYetImplementedStatement(statement *ast.NotYetImplementedStatement) *semantics.NotYetImplementedStatement {
	return &semantics.NotYetImplementedStatement{}
}

func (c *context) createReturnStatement(statement *ast.ReturnStatement) *semantics.ReturnStatement {
	return &semantics.ReturnStatement{
		Value: c.createExpression(statement.Value),
	}
}

func (c *context) createSwitchStatement(statement *ast.SwitchStatement) *semantics.SwitchStatement {
	var caseBlocks semantics.CaseBlocks
	for _, b := range statement.Cases {
		block := &semantics.CaseBlock{
			Values: c.createExpressions(b.Expressions),
			Body:   c.createCodeBlock(b.Body),
		}
		caseBlocks = append(caseBlocks, block)
	}

	var defaultBlock *semantics.DefaultBlock
	if statement.Default != nil {
		defaultBlock = &semantics.DefaultBlock{
			Body: c.createCodeBlock(statement.Default.Body),
		}
	}

	return &semantics.SwitchStatement{
		Expression:   c.createExpression(statement.Expression),
		CaseBlocks:   caseBlocks,
		DefaultBlock: defaultBlock,
	}
}

func (c *context) createUnaryStatement(statement *ast.UnaryStatement) *semantics.UnaryStatement {
	return &semantics.UnaryStatement{
		UnaryPostFixOp: statement.Op,
		Expression:     c.createExpression(statement.Expression),
	}
}

func (c *context) createVariableStatement(statement *ast.VariableStatement) *semantics.VariableStatement {
	return &semantics.VariableStatement{}
}

func (c *context) createWhileStatement(statement *ast.WhileStatement) *semantics.WhileStatement {
	return &semantics.WhileStatement{
		Condition: c.createExpression(statement.Condition),
		Body:      c.createCodeBlock(statement.Body),
	}
}

func (c *context) createExpression(e ast.Expression) semantics.Expression {
	switch e.Kind() {
	case ast.ExprBinary:
		return c.createBinaryExpression(e.(*ast.BinaryExpression))
	case ast.ExprIdentifier:
		return c.createIdentifierExpression(e.(*ast.IdentifierExpression))
	case ast.ExprInvocation:
		return c.createInvocationExpression(e.(*ast.InvocationExpression))
	case ast.ExprLambda:
		return c.createLambdaExpression(e.(*ast.LambdaExpression))
	case ast.ExprLiteral:
		return c.createLiteralExpression(e.(*ast.LiteralExpression))
	case ast.ExprMemberAccess:
		return c.createMemberAccessExpression(e.(*ast.MemberAccessExpression))
	case ast.ExprParenthesis:
		return c.createParenthesisExpression(e.(*ast.ParenthesisExpression))
	case ast.ExprNotYetImplemented:
		return &semantics.NotYetImplementedExpression{}
	case ast.ExprStruct:
		return c.createStructExpression(e.(*ast.StructExpression))
	case ast.ExprTernary:
		return c.createTernaryExpression(e.(*ast.TernaryExpression))
	case ast.ExprUnary:
		return c.createUnaryExpression(e.(*ast.UnaryExpression))
	case ast.ExprValue:
		return c.createValueExpression(e.(*ast.ValueExpression))

	default:
		c.m.Append(msg.Internal(e.Loc(), "Unrecognized expression: %v", e.Kind()))
		return nil
	}
}

func (c *context) createExpressions(es ast.Expressions) semantics.Expressions {
	var result semantics.Expressions
	for _, e := range es {
		result = append(result, c.createExpression(e))
	}
	return result
}

func (c *context) createBinaryExpression(e *ast.BinaryExpression) *semantics.BinaryExpression {
	return &semantics.BinaryExpression{
		Op:    e.Op,
		Left:  c.createExpression(e.Left),
		Right: c.createExpression(e.Right),
	}
}

func (c *context) createIdentifierExpression(e *ast.IdentifierExpression) *semantics.IdentifierExpression {
	return &semantics.IdentifierExpression{
		Identifier: e.Identifier,
	}
}

func (c *context) createInvocationExpression(e *ast.InvocationExpression) *semantics.InvocationExpression {
	return &semantics.InvocationExpression{
		Target:    c.createExpression(e.Target),
		Arguments: c.createExpressions(e.Arguments),
	}
}

func (c *context) createLambdaExpression(e *ast.LambdaExpression) *semantics.LambdaExpression {
	return &semantics.LambdaExpression{
		TypeParameters: c.createTypeParameters(e.TypeParameters),
		Parameters:     c.createParameters(e.Parameters),
		ReturnType:     c.createType(e.ReturnType, nil, nil, nil),
		Constraints:    c.createConstraints(e.Constraints),
		Body:           c.createCodeBlock(e.Body),
	}
}

func (c *context) createLiteralExpression(e *ast.LiteralExpression) *semantics.LiteralExpression {
	return &semantics.LiteralExpression{
		Literal: e.Value,
	}
}

func (c *context) createMemberAccessExpression(e *ast.MemberAccessExpression) *semantics.MemberAccessExpression {
	return &semantics.MemberAccessExpression{
		Target: c.createExpression(e.Target),
		Member: e.Member,
	}
}

func (c *context) createParenthesisExpression(e *ast.ParenthesisExpression) *semantics.ParenthesisExpression {
	return &semantics.ParenthesisExpression{
		Expression: c.createExpression(e.Expression),
	}
}

func (c *context) createStructExpression(e *ast.StructExpression) *semantics.StructExpression {
	var assignments semantics.FieldAssignments
	for _, a := range e.Assignments {
		assignment := &semantics.FieldAssignment{
			Name:  a.Name,
			Value: c.createExpression(a.Value),
		}
		assignments = append(assignments, assignment)
	}

	return &semantics.StructExpression{
		Assignments: assignments,
		Type:        c.createReferencedType(e.Type, semantics.TypeBase{}), // TODO: Generics support
	}
}

func (c *context) createTernaryExpression(e *ast.TernaryExpression) *semantics.TernaryExpression {
	return &semantics.TernaryExpression{
		Condition: c.createExpression(e.Condition),
		OnTrue:    c.createExpression(e.OnTrue),
		OnFalse:   c.createExpression(e.OnFalse),
	}
}

func (c *context) createUnaryExpression(e *ast.UnaryExpression) *semantics.UnaryExpression {
	return &semantics.UnaryExpression{
		Op:    e.Op,
		Value: c.createExpression(e.Expression),
	}
}

func (c *context) createValueExpression(e *ast.ValueExpression) *semantics.ValueExpression {
	return &semantics.ValueExpression{}
}

func (c *context) createAndResolveModuleReferences() {
	for _, source := range c.sources {
		for _, imports := range source.Imports {
			im := c.modules[imports.ModuleReference]
			if im == nil {
				c.m.Append(msg.UnresolvedImport(imports.Location, imports.ModuleReference))
				continue
			}
		}
	}
}

func (c *context) createModules() {
	for _, source := range c.sources {
		id := semantics.DefaultModuleName
		if source.Module != nil {
			id = source.Module.Name
		}

		m := c.s.Modules[id]
		if m == nil {
			m = semantics.NewModule(id)
			c.s.Modules[id] = m
		}

		c.cuToModules[source] = m
		c.modules[id] = m
	}
}
