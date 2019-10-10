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

package common

import (
	"strings"

	"nebularis.io/nebularis/pkg/text"
)

type QualifiedName string

func (q QualifiedName) Write(w *text.Writer) {
	w.W(string(q))
}

func (q QualifiedName) Segments() []Identifier {
	parts := strings.Split(string(q), ".")

	result := make([]Identifier, len(parts))
	for i, p := range parts {
		result[i] = Identifier(p)
	}

	return result
}

func (q QualifiedName) LastSegment() Identifier {
	parts := strings.Split(string(q), ".")
	return Identifier(parts[len(parts)-1])
}
