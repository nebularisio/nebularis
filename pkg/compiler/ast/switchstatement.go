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

type SwitchStatement struct {
	NodeBase
	Expression Expression
	Cases      CaseBlocks
	Default    *DefaultBlock
}

var _ Statement = &SwitchStatement{}

func (SwitchStatement) Kind() StatementKind { return StSwitch }

func (s *SwitchStatement) Write(w *text.Writer) {
	w.W("switch ").W(s.Expression).W(" {")
	w.Indent()

	for _, c := range s.Cases {
		w.Ln().W(c)
	}

	if s.Default != nil {
		w.Ln().W(s.Default)
	}

	w.Dedent()

	w.W("}").Ln()
}

type CaseBlock struct {
	NodeBase
	Expressions Expressions
	Body        *CodeBlock
}

func (c *CaseBlock) Write(w *text.Writer) {
	w.W("case ")

	for i, exp := range c.Expressions {
		if i > 0 {
			w.W(", ")
		}

		w.W(exp)
	}

	w.W(":")

	if !c.Body.Inline {
		w.W(" ")
	}

	w.W(c.Body)

	if !c.Body.Inline || len(c.Body.Statements) == 0 {
		w.Ln()
	}
}

type CaseBlocks []*CaseBlock

type DefaultBlock struct {
	NodeBase
	Body *CodeBlock
}

func (d *DefaultBlock) Write(w *text.Writer) {
	w.W("default:")

	if d.Body != nil {
		if !d.Body.Inline {
			w.W(" ")
		}

		w.W(d.Body)

		if !d.Body.Inline {
			w.Ln()
		}
	}
}
