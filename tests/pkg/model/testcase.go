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
	indexSuffixPattern = regexp.MustCompile(`\.\d+`) // nolint:gochecknoglobals
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

func (c *TestCase) hasSemantics() bool {
	if c.Semantics != nil {
		return true
	}

	if c.CheckErrors != nil {
		return true
	}

	return false
}

func (c *TestCase) isEmpty() bool {
	if c.hasSemantics() {
		return false
	}

	for _, a := range c.Asts {
		if a != nil {
			return false
		}
	}

	for _, e := range c.ParseErrors {
		if e != nil {
			return false
		}
	}

	return true
}

func (c *TestCase) Run(t *testing.T, outfolder string) {
	t.Run(c.Name, func(t *testing.T) {

		for i := range c.Inputs {
			index := i
			if c.Asts[index] != nil {

				if c.ParseErrors[index] == nil {
					t.Run(testName("ast", index), func(t *testing.T) {
						c.runAstTest(t, index, outfolder)
					})
					t.Run(testName("roundtrip", index), func(t *testing.T) {
						c.runRoundtrippedTest(t, index, outfolder)
					})
				} else {
					t.Run(testName("parse error", index), func(t *testing.T) {
						c.runParseErrorTest(t, index, outfolder)
					})
				}
			}
		}

		if c.hasSemantics() {
			var cus []*ast.CompilationUnit
			for _, in := range c.Inputs {
				var m diag.Messages
				cu := parser.ParseText(&m, in)
				cus = append(cus, cu)
			}

			o := analyzer.Options{
				Sources: cus,
			}
			r := analyzer.Analyze(o)

			if c.Semantics != nil {
				t.Run(testName("semantics", 0), func(t *testing.T) {
					c.runSemanticsTest(t, r, outfolder)
				})
			}

			if c.CheckErrors != nil {
				t.Run(testName("semantic errors", 0), func(t *testing.T) {
					c.runSemanticErrorsTest(t, r, outfolder)
				})
			} else if r.Messages.HasErrors() {
				t.Fatalf("Unexpected semantic errors:\n%v\n", r.Messages)
			}
		}
	})
}

func (c *TestCase) runSemanticsTest(t *testing.T, r analyzer.Result, outfolder string) {
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
}

func (c *TestCase) runSemanticErrorsTest(t *testing.T, r analyzer.Result, outfolder string) {
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
}

func (c *TestCase) runAstTest(t *testing.T, index int, outfolder string) {
	var ms diag.Messages
	cu := parser.ParseText(&ms, c.Inputs[index])

	if ms.HasErrors() {
		t.Fatalf("Unexpected parse errors:\n%v\n", ms)
	}

	actualAst := text.ToReflectString(cu)

	writeFileOrFail(t, c.AstOutFilePath(index, outfolder), actualAst)
	if !compare(actualAst, *c.Asts[index]) {
		log(t, "Ast", actualAst, *c.Asts[index])
		if UpdateBaselines() {
			t.Logf("Updating Ast baselines...")
			writeFileOrFail(t, c.AstFilePath(index), actualAst)
		} else {
			t.Fail()
		}
	}
}

func (c *TestCase) runParseErrorTest(t *testing.T, index int, outfolder string) {
	var ms diag.Messages
	_ = parser.ParseText(&ms, *c.Asts[index])

	actualErrs := ms.String()
	if !compare(actualErrs, *c.ParseErrors[index]) {
		log(t, "Parse Errors", actualErrs, *c.ParseErrors[index])
		writeFileOrFail(t, c.ParseErrsOutFilePath(index, outfolder), actualErrs)
		if UpdateBaselines() {
			t.Logf("Updating Ast baselines...")
			writeFileOrFail(t, c.ParseErrsFilePath(index), actualErrs)
		} else {
			t.Fail()
		}
	}
}

func (c *TestCase) runRoundtrippedTest(t *testing.T, index int, outfolder string) {
	var ms diag.Messages
	cu := parser.ParseText(&ms, *c.Asts[index])

	var w text.Writer
	cu.Write(&w)
	txt := w.String()
	ms = diag.Messages{}
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

	sortedCases := sortCases(cases)
	return verifyCases(sortedCases)
}

func sortCases(cases map[string]*TestCase) []*TestCase {
	sortedCases := make([]*TestCase, 0, len(cases))
	for _, v := range cases {
		sortedCases = append(sortedCases, v)
	}
	sort.SliceStable(sortedCases, func(i, j int) bool {
		return strings.Compare(sortedCases[i].Name, sortedCases[j].Name) < 0
	})
	return sortedCases
}

func verifyCases(sortedCases []*TestCase) ([]*TestCase, error) {
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
