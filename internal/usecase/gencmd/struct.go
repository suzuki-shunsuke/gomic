package gencmd

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

type (
	// MockTemplateArg implements domain.MockTemplateArg .
	MockTemplateArg struct {
		methods []domain.Method
		imports [][]string
		Intf    *ast.InterfaceType
		Item    domain.Item
	}

	// Method implements domain.Method .
	Method struct {
		params                      []domain.Var
		results                     []domain.Var
		Field                       *ast.Field
		FuncType                    *ast.FuncType
		Decl                        string
		definition                  string
		setReturnDefinition         string
		setReturnInternalDefinition string
		name                        string
		decl                        string
		isEllipsis                  bool
		imports                     map[string]domain.ImportSpec
	}

	// Var implements domain.Var .
	Var struct {
		name    string
		host    string
		t       string
		pkgName string
	}
)

// Name implements domain.Var#Name .
func (v Var) Name() string {
	return v.name
}

// Host implements domain.Var#Host .
func (v Var) Host() string {
	return v.host
}

// Type implements domain.Var#Type .
func (v Var) Type() string {
	return v.t
}

// PkgName implements domain.Var#PkgName .
func (v Var) PkgName() string {
	return v.pkgName
}

// PkgPath implements domain.Var#PkgPath .
func (v Var) PkgPath() string {
	return ""
}

// Name implements domain.Method#Name .
func (method Method) Name() string {
	return method.name
}

// Imports implements domain.Method#Imports .
func (method Method) Imports() map[string]domain.ImportSpec {
	return method.imports
}

// Declaration implements domain.Method#Declaration .
func (method Method) Declaration() string {
	// "Foo func() error"
	// {{.Name}} func() error"
	return method.decl
}

// Definition implements domain.Method#Definition .
func (method Method) Definition() string {
	// "Foo() error"
	//  {{.Name}}() error"
	return method.definition
}

// SetReturnDefinition implements domain.Method#SetReturnDefinition .
func (method Method) SetReturnDefinition() string {
	return method.setReturnDefinition
}

// SetReturnInternalDefinition implements domain.Method#SetReturnInternalDefinition .
func (method Method) SetReturnInternalDefinition() string {
	return method.setReturnInternalDefinition
}

// Params implements domain.Method#Params .
func (method Method) Params() []domain.Var {
	return method.params
}

// ParamsStr implements domain.Method#ParamsStr .
func (method Method) ParamsStr() string {
	// "a0, a1, ..."
	if len(method.Params()) == 0 {
		return ""
	}
	size := len(method.Params())
	arr := make([]string, size)
	for i, v := range method.Params() {
		arr[i] = v.Name()
	}
	if method.IsEllipsis() {
		arr[size-1] = fmt.Sprintf("%s...", method.Params()[size-1].Name())
	}
	return strings.Join(arr, ", ")
}

// Results implements domain.Method#Results .
func (method Method) Results() []domain.Var {
	return method.results
}

// ResultValuesStr implements domain.Method#ResultValuesStr .
func (method Method) ResultValuesStr() string {
	// "a0, a1, ..."
	results := method.Results()
	if len(results) == 0 {
		return ""
	}
	arr := make([]string, len(results))
	for i, v := range results {
		name := v.Name()
		if name == "" {
			name = fmt.Sprintf("r%d", i)
		}
		arr[i] = name
	}
	return strings.Join(arr, ", ")
}

// IsEllipsis implements domain.Method#IsEllipsis .
func (method Method) IsEllipsis() bool {
	return method.isEllipsis
}

// Version implements domain.MockTemplateArg#Version .
func (mock MockTemplateArg) Version() string {
	return domain.Version
}

// URL implements domain.MockTemplateArg#URL .
func (mock MockTemplateArg) URL() string {
	return domain.URL
}

// PackageName implements domain.MockTemplateArg#PackageName .
func (mock MockTemplateArg) PackageName() string {
	return mock.Item.Dest.Package
}

// MockName implements domain.MockTemplateArg#MockName .
func (mock MockTemplateArg) MockName() string {
	return mock.Item.Src.Name
}

// Imports implements domain.MockTemplateArg#Imports .
func (mock MockTemplateArg) Imports() [][]string {
	return mock.imports
}

// Methods implements domain.MockTemplateArg#Methods .
func (mock MockTemplateArg) Methods() []domain.Method {
	return mock.methods
}
