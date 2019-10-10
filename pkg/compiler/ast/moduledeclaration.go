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

type ModuleDeclaration struct {
	NodeBase
	Attributes Attributes
	Name       common.QualifiedName
}

var _ Node = &ModuleDeclaration{}

func (m *ModuleDeclaration) Write(w *text.Writer) {
	w.W(m.Attributes).W("module ").W(m.Name).W(";")
}
