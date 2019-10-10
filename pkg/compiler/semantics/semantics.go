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

type Semantics struct {
	Modules map[common.QualifiedName]*Module
}

func New() *Semantics {
	return &Semantics{
		Modules: make(map[common.QualifiedName]*Module),
	}
}
