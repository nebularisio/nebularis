// Copyright 2019 Nebularis Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package semantics

import (
	"nebularis.io/nebularis/pkg/compiler/common"
)

const (
	DefaultModuleName = common.QualifiedName("$default")
)

type Module struct {
	Name       common.QualifiedName
	Sealed     bool
	Attributes Attributes
	References ModuleRefs                      // TODO:: semantics that will contain Attr
	Types      map[common.Identifier]Type      // TODO: Overloads
	Functions  map[common.Identifier]*Function // TODO: Handle overloads
}

func NewModule(name common.QualifiedName) *Module {
	return &Module{
		Name:       name,
		References: NewModuleRefs(),
		Types:      make(map[common.Identifier]Type),
		Functions:  make(map[common.Identifier]*Function),
	}
}
