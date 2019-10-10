// Copyright 2019 Nebularis Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package semantics

import "nebularis.io/nebularis/pkg/compiler/ast"

type LiteralExpression struct {
	Literal ast.Literal
}

var _ Expression = &LiteralExpression{}

func (e *LiteralExpression) ExpressionType() Type {
	switch e.Literal.Kind() {
	case ast.LitBool:
		return &PrimitiveType{PrimitiveKind: ast.Bool}
	case ast.LitInt:
		return &PrimitiveType{PrimitiveKind: ast.Int32}
	case ast.LitString:
		return &PrimitiveType{PrimitiveKind: ast.String}
	default:
		panic("NYI")
	}
}
