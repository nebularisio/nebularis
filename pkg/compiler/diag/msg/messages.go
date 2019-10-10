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

package msg

import (
	"fmt"

	"nebularis.io/nebularis/pkg/compiler/common"
	"nebularis.io/nebularis/pkg/compiler/diag"
)

func Internal(loc common.Location, msg string, args ...interface{}) diag.Message {
	return err(loc, 0, "Internal Error: %s", fmt.Sprintf(msg, args...))
}

func NYI(loc common.Location, msg string, args ...interface{}) diag.Message {
	return err(loc, 1, "NYI: %s", fmt.Sprintf(msg, args...))
}

func Syntax(loc common.Location, msg string) diag.Message {
	return err(loc, 1000, "Syntax Error: '%s'", msg)
}

func MultipleDefaults(loc common.Location) diag.Message {
	return err(loc, 1001, "Multiple default blocks found in switch")
}

func UnresolvedImport(loc common.Location, name common.QualifiedName) diag.Message {
	return err(loc, 2001, "Unresolved import: '%s'", name)
}

func MemberAlreadyDeclared(loc common.Location, name common.Identifier) diag.Message {
	return err(loc, 2002, "Element is already declared: '%s'", name)
}

func MethodAlreadyDeclared(loc common.Location, name common.Identifier) diag.Message {
	return err(loc, 2003, "Method is already declared: '%s'", name)
}

func ParameterAlreadyDeclared(loc common.Location, name common.Identifier) diag.Message {
	return err(loc, 2004, "Parameter is already declared: '%s'", name)
}

func FieldAlreadyDeclared(loc common.Location, name common.Identifier) diag.Message {
	return err(loc, 2005, "ResolvedTo is already declared: '%s'", name)
}

func AlreadyNullable(loc common.Location) diag.Message {
	return err(loc, 2006, "Type is already nullable")
}

func TypeNotFound(loc common.Location, name common.QualifiedName) diag.Message {
	return err(loc, 2007, "Type not found: %v", name)
}

func TypeAlreadyDeclared(loc common.Location, name common.Identifier) diag.Message {
	return err(loc, 2008, "Type is already declared: %v", name)
}

//func InvalidBaseType(loc common.Location, name common.QualifiedName) diagnostic.Message {
//	return err(loc, 2006, "Invalid base type: %s", name)
//}

//func MemberNeedsEitherTypeOrInitializer(loc ast.Location, name ast.Identifier) diagnostic.Message {
//	return err(loc,  4, "Element declaration requires either an explicit type, or an initializer: '%s'", name)
//}

func err(loc common.Location, code int, fmt string, params ...interface{}) diag.Message {
	return diag.NewMessage(
		diag.Error,
		loc,
		code,
		fmt,
		params)
}
