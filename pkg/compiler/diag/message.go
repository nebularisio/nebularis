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

package diag

import (
	"fmt"

	"nebularis.io/nebularis/pkg/compiler/common"
)

type Message struct {
	Level    Level
	Location common.Location
	Code     int
	Format   string
	Params   []interface{}
}

func NewMessage(level Level, loc common.Location, code int, format string, params []interface{}) Message {
	return Message{
		Level:    level,
		Location: loc,
		Code:     code,
		Format:   format,
		Params:   params,
	}
}

func (m Message) String() string {
	return fmt.Sprintf("%v %04d %v %s",
		m.Level,
		m.Code,
		m.Location,
		fmt.Sprintf(m.Format, m.Params...))
}
