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

type StatementKind string

const (
	StVariable          = StatementKind("Variable")
	StReturn            = StatementKind("Return")
	StIf                = StatementKind("If")
	StWhile             = StatementKind("While")
	StLabel             = StatementKind("Label")
	StContinue          = StatementKind("Continue")
	StBreak             = StatementKind("Break")
	StExpression        = StatementKind("Expression")
	StNotYetImplemented = StatementKind("NotYetImplemented")
	StUnary             = StatementKind("Unary")
	StSwitch            = StatementKind("Switch")
	StAssigment         = StatementKind("Assignment")
)

type Statement interface {
	Node
	Kind() StatementKind
}

type Statements []Statement

func (s Statements) Write(w *text.Writer) {
	for _, st := range s {
		w.W(st)
	}
}
