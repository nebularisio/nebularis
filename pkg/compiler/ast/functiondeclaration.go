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

type FunctionDeclaration struct {
	FunctionTypeSpec
	Attributes     Attributes
	Extern         bool
	Public         bool
	Pure           bool
	Target         *TypeReference
	Name           common.Identifier
	TypeParameters TypeParameters
	Constraints    WhereClauses
	Body           *CodeBlock
}

var _ Declaration = &FunctionDeclaration{}

func (t *FunctionDeclaration) Kind() DeclarationKind         { return DeclFunction }
func (t *FunctionDeclaration) LocalName() common.Identifier  { return t.Name }
func (t *FunctionDeclaration) DeclLocation() common.Location { return t.Location }

func (t *FunctionDeclaration) Write(w *text.Writer) {
	w.W(t.Attributes)

	if t.Extern {
		w.Fmt("extern ")
	}

	if t.Public {
		w.Fmt("public ")
	}

	if t.Pure {
		w.Fmt("pure ")
	}

	w.W("fn ")

	w.W(t.TypeParameters)

	if len(t.TypeParameters) != 0 {
		w.W(" ")
	}

	if t.Target != nil {
		w.W("(").W(t.Target).W(") ")
	}

	w.W(t.Name).W(t.Parameters).W(" ")

	if t.ReturnType != nil {
		w.W(t.ReturnType).W(" ")
	}

	if len(t.Constraints) > 0 {
		w.W(t.Constraints).W(" ")
	}

	w.W(t.Body).Ln()
}
