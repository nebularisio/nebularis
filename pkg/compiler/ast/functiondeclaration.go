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

func (s *FunctionDeclaration) Write(w *text.Writer) {
	w.W(s.Attributes)
	if s.Extern {
		w.Fmt("extern ")
	}
	if s.Public {
		w.Fmt("public ")
	}
	if s.Pure {
		w.Fmt("pure ")
	}
	w.W("fn ")

	w.W(s.TypeParameters)
	if len(s.TypeParameters) != 0 {
		w.W(" ")
	}

	if s.Target != nil {
		w.W("(").W(s.Target).W(") ")
	}
	w.W(s.Name).W(s.Parameters).W(" ")
	if s.ReturnType != nil {
		w.W(s.ReturnType).W(" ")
	}
	if len(s.Constraints) > 0 {
		w.W(s.Constraints).W(" ")
	}
	w.W(s.Body).Ln()
}
