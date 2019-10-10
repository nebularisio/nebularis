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

type TypeDeclaration struct {
	NodeBase
	Attributes     Attributes
	TypeParameters TypeParameters
	Name           common.Identifier
	Spec           TypeSpec
	Constraints    WhereClauses
}

var _ Declaration = &TypeDeclaration{}

func (t *TypeDeclaration) Kind() DeclarationKind         { return DeclType }
func (t *TypeDeclaration) LocalName() common.Identifier  { return t.Name }
func (t *TypeDeclaration) DeclLocation() common.Location { return t.Location }

func (i *TypeDeclaration) Write(w *text.Writer) {
	w.W(i.Attributes).W("type ").W(i.Name).W(i.TypeParameters)
	w.W(" ").W(i.Spec)
	if len(i.Constraints) > 0 {
		w.W(" ").W(i.Constraints)
	}
	w.W(";").Ln()
}
