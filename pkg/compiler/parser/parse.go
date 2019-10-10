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

package parser

import (
	"nebularis.io/nebularis/gen/parser"
	"nebularis.io/nebularis/pkg/compiler/ast"
	"nebularis.io/nebularis/pkg/compiler/common"
	"nebularis.io/nebularis/pkg/compiler/diag"
	"nebularis.io/nebularis/pkg/compiler/diag/msg"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func Parse(ms *diag.Messages, s antlr.CharStream) *ast.CompilationUnit {
	lexer := parser.NewNebularisLexer(s)
	str := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewNebularisParser(str)
	p.AddErrorListener(&listener{ms})

	b := builder{ms: ms}
	b.compilationUnit(p.CompilationUnit())
	return b.cu
}

func ParseText(ms *diag.Messages, text string) *ast.CompilationUnit {
	s := antlr.NewInputStream(text)
	return Parse(ms, s)
}

type listener struct {
	ms *diag.Messages
}

func (l *listener) SyntaxError(
	recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, m string, e antlr.RecognitionException) {
	l.ms.Append(msg.Syntax(common.Location{Line: int32(line), Column: int32(column)}, m))
}

func (l *listener) ReportAmbiguity(
	recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (l *listener) ReportAttemptingFullContext(
	recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (l *listener) ReportContextSensitivity(
	recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {

}
