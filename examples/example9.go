package examples

import (
	"go/ast"
	"io"
	"os"
)

type (
	// Hello9 is an example interface.
	Hello9 interface {
		Hello(*ast.File) map[io.Reader]os.File
		Hello2() func() string
		Hello3() struct{}
	}
)
