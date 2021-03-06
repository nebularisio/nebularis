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

import (
	"nebularis.io/nebularis/pkg/compiler/common"
	"nebularis.io/nebularis/pkg/text"
)

type VariableStatement struct {
	NodeBase
	Name        common.Identifier
	Type        TypeSpec
	Initializer Expression
}

var _ Statement = &VariableStatement{}

func (VariableStatement) Kind() StatementKind { return StVariable }

func (v *VariableStatement) Write(w *text.Writer) {
	w.W("var ").W(v.Name)

	if v.Type != nil {
		w.W(" : ").W(v.Type)
	}

	if v.Initializer != nil {
		w.W(" = ").W(v.Initializer)
	}

	w.W(";").Ln()
}
