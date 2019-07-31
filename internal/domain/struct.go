package domain

import (
	"go/ast"
)

type (
	// Config represents configuration.
	Config struct {
		Default DefaultConfiguration
		Items   []Item
	}

	// DefaultConfiguration represents default configuration.
	DefaultConfiguration struct {
		SrcDefaultConfiguration `yaml:",inline"`
	}

	// SrcDefaultConfiguration represents source's default configuration.
	SrcDefaultConfiguration struct {
		InterfacePrefix string `yaml:"interface_prefix"`
		InterfaceSuffix string `yaml:"interface_suffix"`
		VendorDir       string `yaml:"vendor_dir"`
	}

	// Item represents configuration of each mock.
	Item struct {
		Src                  Src
		Dest                 Dest
		DefaultConfiguration `yaml:",inline"`
	}

	// Src represents source configuration.
	Src struct {
		Package                 string
		Interface               string
		Name                    string
		File                    string
		Dir                     string
		SrcDefaultConfiguration `yaml:",inline"`
	}

	// Dest represents generated file's configuration.
	Dest struct {
		Package string
		File    string
	}

	Interface struct {
		Package   *ast.Package
		File      *ast.File
		Interface *ast.InterfaceType
		Fields    []*ast.Field
		FuncTypes []*ast.FuncType
	}

	Interfaces struct {
		Interfaces []Interface
	}

	Function struct {
		FuncType *ast.FuncType
		Field    *ast.Field
	}
)
