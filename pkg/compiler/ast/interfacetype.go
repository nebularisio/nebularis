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
	"nebularis.io/nebularis/pkg/text"
)

type InterfaceType struct {
	NodeBase
	TypeParameters TypeParameters
	Extends        TypeReferences
	Methods        Methods
}

var _ TypeSpec = &InterfaceType{}

func (t *InterfaceType) Kind() TypeKind { return Interface }

func (t *InterfaceType) Write(w *text.Writer) {
	w.W("interface").W(t.TypeParameters).W(" ")

	if len(t.Extends) > 0 {
		w.W(": ")

		for i, e := range t.Extends {
			if i > 0 {
				w.W(", ")
			}

			w.W(e)
		}

		w.W(" ")
	}

	w.W("{")

	if len(t.Methods) > 0 {
		w.Ln().Indent()
		w.W(t.Methods)
		w.Dedent()
	}

	w.W("}")
}
