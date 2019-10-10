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

package text

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func ToString(i interface{}) string {
	var w Writer
	w.W(i)
	return w.String()
}

func ToReflectString(i interface{}) string {
	m := make(map[interface{}]struct{})
	var w Writer
	writeInterface(&w, m, i)
	return w.String()
}

func writeInterface(w *Writer, m map[interface{}]struct{}, iface interface{}) {
	if iface == nil {
		w.W("<nil>")
		return
	}

	t := reflect.TypeOf(iface)
	switch t.Kind() {
	case reflect.Ptr:
		if _, found := m[iface]; found {
			w.W("...")
			return
		}
		m[iface] = struct{}{}
		if reflect.ValueOf(iface).IsNil() {
			return
		}
		writeInterface(w, m, reflect.ValueOf(iface).Elem().Interface())

	case reflect.String:
		w.W(reflect.ValueOf(iface).String())

	case reflect.Bool:
		w.W(reflect.ValueOf(iface).Bool())

	case reflect.Slice:
		w.W("[]{").Indent()
		needComma := false
		v := reflect.ValueOf(iface)
		for i := 0; i < v.Len(); i++ {
			if needComma {
				w.W(",")
			}
			needComma = true
			w.Ln()
			writeInterface(w, m, v.Index(i).Interface())
		}
		w.Dedent().Ln().W("}")

	case reflect.Struct:
		var names []string
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			names = append(names, f.Name)
		}
		sort.Strings(names)

		w.W(t.Name()).W(" {").Indent()
		needComma := false
		for _, name := range names {
			if strings.ToLower(string(name[0])) == string(name[0]) {
				continue
			}
			f, _ := t.FieldByName(name)
			val := reflect.ValueOf(iface).FieldByName(name).Interface()
			if val == nil || reflect.ValueOf(val).IsZero() {
				continue
			}
			if needComma {
				w.W(",")
			}
			needComma = true
			w.Ln()

			w.W(f.Name).W(": ")
			writeInterface(w, m, val)
		}
		w.Dedent().Ln().W("}")

	case reflect.Map:
		keys := reflect.ValueOf(iface).MapKeys()
		sort.SliceStable(keys, func(i, j int) bool {
			return strings.Compare(keys[i].String(), keys[j].String()) < 0
		})
		w.W("{").Indent()
		for j, k := range keys {
			if j > 0 {
				w.W(",")
			}
			w.Ln()
			w.W(k.String()).W(": ")
			writeInterface(w, m, reflect.ValueOf(iface).MapIndex(k).Interface())
		}

		w.Dedent().Ln().W("}")

	case reflect.Int32, reflect.Int:
		w.W(reflect.ValueOf(iface).Int())

	default:
		panic(fmt.Sprintf("NYI: %v", t.Kind()))
	}
}
