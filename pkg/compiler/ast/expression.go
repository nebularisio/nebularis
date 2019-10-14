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

package ast

type ExpressionKind string

const (
	ExprNotYetImplemented = "NotYetImplemented"
	ExprIdentifier        = "Identifier"
	ExprMemberAccess      = "MemberAccess"
	ExprInvocation        = "Invocation"
	ExprIndex             = "Index"
	ExprTernary           = "Ternary"
	ExprLiteral           = "Literal"
	ExprParenthesis       = "Parenthesis"
	ExprValue             = "Value"
	ExprStruct            = "Target"
	ExprLambda            = "Lambda"
	ExprUnary             = "Unary"
	ExprBinary            = "Binary"
)

type Expression interface {
	Node
	Kind() ExpressionKind
}

type Expressions []Expression
