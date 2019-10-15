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

type StructType struct {
	NodeBase
	TypeParameters TypeParameters
	Fields         Fields
}

var _ TypeSpec = &StructType{}

func (s *StructType) Kind() TypeKind {
	return Struct
}

func (s *StructType) Write(w *text.Writer) {
	w.W("struct").W(s.TypeParameters).W(" {")

	if len(s.Fields) > 0 {
		w.Ln().Indent()

		for i, f := range s.Fields {
			if i > 0 {
				w.Ln()
			}

			w.W(f)
		}

		w.Dedent()
	}

	w.W("}")
}
