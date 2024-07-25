package object

import (
	"bytes"
	"fmt"
	"interpreter/ast"
	"strings"
)

const (
	INTEGER_OBJ 			= "INTEGER"
	BOOLEAN_OBJ 			= "BOOLEAN"
	VOID_OBJ					= "VOID"
	RETURN_VALUE_OBJ	= "RETURN_VALUE"
	ERROR_OBJ					= "ERROR"
	FUNCTION_OBJ			= "FUNCTION"
)

type Function struct {
	Parameters []*ast.Identifier
	ReturnType *ast.Identifier
	Body *ast.BlockStatement
	Env *Environment
}
func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) ReturnTypeLiteral() string {
	if f.ReturnType != nil {
		return f.ReturnType.Literal()
	}
	return ""
}
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	if f.ReturnType != nil {
		out.WriteString(f.ReturnTypeLiteral())
	}
	out.WriteString("{\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

type Error struct {
	Message string
}
func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string { return "SKILL ISSUE: " + e.Message }

type ReturnValue struct {
	Value Object
}
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

type ObjectType string
type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

type Boolean struct {
	Value bool
}
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

type Void struct{}
func (v *Void) Type() ObjectType { return VOID_OBJ }
func (v *Void) Inspect() string { return "void" }
