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

import "nebularis.io/nebularis/pkg/text"

type UnaryPrefixOp string

const (
	UnaryNot   = UnaryPrefixOp("!")
	UnaryTilda = UnaryPrefixOp("~")
)

type UnaryExpression struct {
	NodeBase
	Op         UnaryPrefixOp
	Expression Expression
}

var _ Expression = &UnaryExpression{}

func (e *UnaryExpression) Kind() ExpressionKind { return ExprUnary }

func (e *UnaryExpression) Write(w *text.Writer) {
	w.W(e.Op).W(e.Expression)
}
