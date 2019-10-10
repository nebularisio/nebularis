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

package model

import (
	"fmt"
	"io/ioutil"
	"path"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"

	"nebularis.io/nebularis/pkg/compiler/analyzer"
	"nebularis.io/nebularis/pkg/compiler/ast"
	"nebularis.io/nebularis/pkg/compiler/diag"
	"nebularis.io/nebularis/pkg/compiler/parser"
	"nebularis.io/nebularis/pkg/text"
)

var (
	indexSuffixPattern = regexp.MustCompile(`\.\d+`)
)

type TestCase struct {
	Folder      string
	Name        string
	Inputs      []string
	Asts        []*string
	ParseErrors []*string
	Semantics   *string
	CheckErrors *string
}

func (c *TestCase) InputFilePath(index *int) string {
	var indexExt string
	if index != nil {
		indexExt = "." + strconv.FormatInt(int64(*index), 10)
	}
	return path.Join(c.Folder, c.Name+indexExt+".ns")
}

func (c *TestCase) AstFilePathFor(index *int) string {
	var indexExt string
	if index != nil {
		indexExt = "." + strconv.FormatInt(int64(*index), 10)
	}
	return path.Join(c.Folder, c.Name+indexExt+".ns.ast")
}

func (c *TestCase) AstFilePath(index int) string {
	return c.AstOutFilePath(index, c.Folder)
}

func (c *TestCase) AstOutFilePath(index int, folder string) string {
	var indexExt string
	if index != 0 || len(c.Asts) > 1 {
		indexExt = "." + strconv.FormatInt(int64(index), 10)
	}
	return path.Join(folder, c.Name+indexExt+".ns.ast")
}

func (c *TestCase) ParseErrsFilePathFor(index *int) string {
	var indexExt string
	if index != nil {
		indexExt = "." + strconv.FormatInt(int64(*index), 10)
	}
	return path.Join(c.Folder, c.Name+indexExt+".ns.ast.errs")
}

func (c *TestCase) ParseErrsFilePath(index int) string {
	return c.ParseErrsOutFilePath(index, c.Folder)
}

func (c *TestCase) ParseErrsOutFilePath(index int, folder string) string {
	var indexExt string
	if index != 0 || len(c.Asts) > 1 {
		indexExt = "." + strconv.FormatInt(int64(index), 10)
	}
	return path.Join(folder, c.Name+indexExt+".ns.ast.errs")
}

func (c *TestCase) SemFilePath() string {
	return c.SemOutFilePath(c.Folder)
}

func (c *TestCase) SemOutFilePath(folder string) string {
	return path.Join(folder, c.Name+".ns.sem")
}

func (c *TestCase) SemErrsFilePath() string {
	return c.SemErrsOutFilePath(c.Folder)
}

func (c *TestCase) SemErrsOutFilePath(folder string) string {
	return path.Join(folder, c.Name+".ns.sem.errs")
}

func (c *TestCase) EnsureSize(size int) {
	i := make([]string, size)
	copy(i, c.Inputs)
	c.Inputs = i

	arr := make([]*string, size)
	copy(arr, c.Asts)
	c.Asts = arr

	arr = make([]*string, size)
	copy(arr, c.ParseErrors)
	c.ParseErrors = arr
}

func testName(baseName string, index int) string {
	if index == 0 {
		return baseName
	}
	return fmt.Sprintf("%s[%d]", baseName, index)
}

func (tc *TestCase) hasSemantics() bool {
	if tc.Semantics != nil {
		return true
	}

	if tc.CheckErrors != nil {
		return true
	}

	return false
}

func (tc *TestCase) isEmpty() bool {
	if tc.hasSemantics() {
		return false
	}

	for _, a := range tc.Asts {
		if a != nil {
			return false
		}
	}

	for _, e := range tc.ParseErrors {
		if e != nil {
			return false
		}
	}

	return true
}

func (c *TestCase) Run(t *testing.T, outfolder string) {
	t.Run(c.Name, func(t *testing.T) {

		var cus []*ast.CompilationUnit

		for i, s := range c.Inputs {
			var ms diag.Messages
			cu := parser.ParseText(&ms, s)
			if c.Asts[i] != nil {
				t.Run(testName("ast", i), func(t *testing.T) {
					actualAst := text.ToReflectString(cu)

					writeFileOrFail(t, c.AstOutFilePath(i, outfolder), actualAst)
					if !compare(actualAst, *c.Asts[i]) {
						log(t, "Ast", actualAst, *c.Asts[i])
						if UpdateBaselines() {
							t.Logf("Updating Ast baselines...")
							writeFileOrFail(t, c.AstFilePath(i), actualAst)
						} else {
							t.Fail()
						}
					}
				})
			}

			if c.ParseErrors[i] != nil {
				t.Run(testName("parse error", i), func(t *testing.T) {
					actualErrs := ms.String()
					if !compare(actualErrs, *c.ParseErrors[i]) {
						log(t, "Parse Errors", actualErrs, *c.ParseErrors[i])
						writeFileOrFail(t, c.ParseErrsOutFilePath(i, outfolder), actualErrs)
						if UpdateBaselines() {
							t.Logf("Updating Ast baselines...")
							writeFileOrFail(t, c.ParseErrsFilePath(i), actualErrs)
						} else {
							t.Fail()
						}
					}
				})
			} else {
				if ms.HasErrors() {
					t.Fatalf("Unexpected parse errors:\n%v\n", ms)
				}

				t.Run(testName("roundtrip", i), func(t *testing.T) {
					var w text.Writer
					cu.Write(&w)
					txt := w.String()
					var ms diag.Messages
					roundtrippedCu := parser.ParseText(&ms, txt)
					if ms.HasErrors() {
						t.Fatalf("Unexpected parse errors:\n%s\n%v\n", txt, ms)
					}

					w = text.Writer{}
					roundtrippedCu.Write(&w)
					writeFileOrFail(t, path.Join(outfolder, c.Name+".ns"), w.String())
					if !reflect.DeepEqual(roundtrippedCu, cu) {
						t.Logf("----")
						t.Logf("Output:\n%s\n", txt)
						t.Logf("----")
						t.Logf("Original: \n%s\n", text.ToReflectString(cu))
						t.Logf("----")
						t.Logf("Rounttripped: \n%s\n", text.ToReflectString(roundtrippedCu))
						t.Logf("----")
						t.Fail()
					}
				})
			}

			cus = append(cus, cu)
		}

		if c.hasSemantics() {
			o := analyzer.Options{
				Sources: cus,
			}
			r := analyzer.Analyze(o)

			if c.Semantics != nil {
				t.Run(testName("semantics", 0), func(t *testing.T) {
					actualSem := text.ToReflectString(r.Semantics)
					if !compare(actualSem, *c.Semantics) {
						log(t, "semantics", actualSem, *c.Semantics)
						writeFileOrFail(t, c.SemOutFilePath(outfolder), actualSem)
						if UpdateBaselines() {
							t.Logf("Updating semantics baselines...")
							writeFileOrFail(t, c.SemFilePath(), actualSem)
						} else {
							t.Fail()
						}
					}
				})
			}

			if c.CheckErrors != nil {
				t.Run(testName("semantic errors", 0), func(t *testing.T) {
					actualErrs := r.Messages.String()
					if !compare(actualErrs, *c.CheckErrors) {
						log(t, "semantic errors", actualErrs, *c.CheckErrors)
						writeFileOrFail(t, c.SemErrsOutFilePath(outfolder), actualErrs)
						if UpdateBaselines() {
							t.Logf("Updating semantic error baselines...")
							writeFileOrFail(t, c.SemErrsFilePath(), actualErrs)
						} else {
							t.Fail()
						}
					}
				})
			} else if r.Messages.HasErrors() {
				t.Fatalf("Unexpected semantic errors:\n%v\n", r.Messages)
			}
		}
	})
}

func Load(folder string) ([]*TestCase, error) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	cases := make(map[string]*TestCase)
	for _, f := range files {
		if !isInputFile(f.Name()) {
			continue
		}

		baseName, index, err := parseInputName(f.Name())
		if err != nil {
			return nil, err
		}

		tc := cases[baseName]
		if tc == nil {
			tc = &TestCase{
				Folder: folder,
				Name:   baseName,
			}

			if tc.CheckErrors, err = tryReadFile(tc.SemErrsFilePath()); err != nil {
				return nil, err
			}

			if tc.Semantics, err = tryReadFile(tc.SemFilePath()); err != nil {
				return nil, err
			}

			cases[baseName] = tc
		}

		if index == nil {
			tc.EnsureSize(1)
		} else if len(tc.Inputs) <= *index {
			tc.EnsureSize(*index + 1)
		}

		idx := 0
		if index != nil {
			idx = *index
		}
		if tc.Inputs[idx], err = readFile(tc.InputFilePath(index)); err != nil {
			return nil, err
		}

		if tc.ParseErrors[idx], err = tryReadFile(tc.ParseErrsFilePathFor(index)); err != nil {
			return nil, err
		}

		if tc.Asts[idx], err = tryReadFile(tc.AstFilePathFor(index)); err != nil {
			return nil, err
		}
	}

	var sortedCases []*TestCase
	for _, v := range cases {
		sortedCases = append(sortedCases, v)
	}
	sort.SliceStable(sortedCases, func(i, j int) bool {
		return strings.Compare(sortedCases[i].Name, sortedCases[j].Name) < 0
	})

	for _, s := range sortedCases {
		if len(s.Inputs) == 0 {
			return nil, fmt.Errorf("test case with no inputs: %s", s.Name)
		}

		if s.isEmpty() {
			return nil, fmt.Errorf("test case with no expected files: %s", s.Name)
		}
	}

	return sortedCases, nil
}
