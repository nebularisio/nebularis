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

type NullableTypeSpec struct {
	NodeBase
	TypeParameters TypeParameters
	Element        TypeSpec
}

var _ TypeSpec = &NullableTypeSpec{}

func (p *NullableTypeSpec) Kind() TypeKind { return Nullable }

func (p *NullableTypeSpec) Write(w *text.Writer) {
	w.W(p.TypeParameters).W("?").W(p.Element)
}
