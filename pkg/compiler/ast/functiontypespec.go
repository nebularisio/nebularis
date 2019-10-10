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

type FunctionTypeSpec struct {
	NodeBase
	Target     *TypeReference
	Parameters FunctionParameters
	ReturnType TypeSpec
}

var _ TypeSpec = &FunctionTypeSpec{}

func (f *FunctionTypeSpec) Kind() TypeKind {
	return Function
}

func (s *FunctionTypeSpec) Write(w *text.Writer) {
	w.W("fn")

	if s.Target != nil {
		w.W("(").W(s.Target).W(") ")
	}

	s.Parameters.Write(w)
	if s.ReturnType != nil {
		w.W(" ").W(s.ReturnType)
	}
}
