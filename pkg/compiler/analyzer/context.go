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

package analyzer

import (
	"nebularis.io/nebularis/pkg/compiler/ast"
	"nebularis.io/nebularis/pkg/compiler/common"
	"nebularis.io/nebularis/pkg/compiler/diag"
	"nebularis.io/nebularis/pkg/compiler/semantics"
)

type context struct {
	sources []*ast.CompilationUnit
	s       *semantics.Semantics
	m       diag.Messages

	cuToModules map[*ast.CompilationUnit]*semantics.Module
	decls       map[semantics.DeclRef]semantics.Node

	modules map[common.QualifiedName]*semantics.Module
}

func newContext(sources []*ast.CompilationUnit) *context {
	return &context{
		s:           semantics.New(),
		sources:     sources,
		cuToModules: make(map[*ast.CompilationUnit]*semantics.Module),
		decls:       make(map[semantics.DeclRef]semantics.Node),
		modules:     make(map[common.QualifiedName]*semantics.Module),
	}
}

func (c *context) moduleFor(decl *ast.ModuleDeclaration) *semantics.Module {
	id := semantics.DefaultModuleName
	if decl != nil {
		id = decl.Name
	}

	m := c.s.Modules[id]
	if m == nil {
		m = semantics.NewModule(id)
		c.s.Modules[id] = m
	}

	return m
}

func identOrEmpty(i *common.Identifier) common.Identifier {
	if i == nil {
		return ""
	}
	return *i
}
