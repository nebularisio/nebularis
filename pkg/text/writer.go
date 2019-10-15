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
	"bytes"
	"fmt"
	"strings"
)

type Writer struct {
	b      bytes.Buffer
	indent int

	inline bool
}

func (w *Writer) Indent() *Writer {
	w.indent += 2
	return w
}

func (w *Writer) Dedent() *Writer {
	w.indent -= 2
	return w
}

func (w *Writer) W(v interface{}) *Writer {
	w.ensureIndent()

	if wr, ok := v.(Writeable); ok {
		wr.Write(w)
	} else {
		w.Fmt("%v", v)
	}

	return w
}

func (w *Writer) Fmt(format string, args ...interface{}) *Writer {
	w.ensureIndent()
	w.write(format, args...)

	return w
}

func (w *Writer) FmtLn(format string, args ...interface{}) *Writer {
	w.ensureIndent()
	w.write(format, args...)
	w.Ln()

	return w
}

func (w *Writer) String() string {
	return w.b.String()
}

func (w *Writer) ensureIndent() {
	if !w.inline {
		_, _ = fmt.Fprint(&w.b, strings.Repeat(" ", w.indent*2))
		w.inline = true
	}
}

func (w *Writer) write(format string, args ...interface{}) {
	for i, a := range args {
		if w, ok := a.(Writeable); ok {
			wr := Writer{}
			w.Write(&wr)
			args[i] = wr.String()
		}
	}

	_, _ = fmt.Fprintf(&w.b, format, args...)
	w.inline = true
}

func (w *Writer) Ln() *Writer {
	_, _ = fmt.Fprintln(&w.b)
	w.inline = false

	return w
}
