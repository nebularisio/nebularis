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

type CodeBlock struct {
	NodeBase
	Statements Statements
	Inline     bool
}

var _ Node = &CodeBlock{}

func (e *CodeBlock) Write(w *text.Writer) {
	if !e.Inline {
		w.W("{")
	}

	if len(e.Statements) > 0 {
		w.Ln().Indent()

		for _, s := range e.Statements {
			w.W(s)
		}

		w.Dedent()
	}

	if !e.Inline {
		w.W("}")
	}
}
