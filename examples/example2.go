package examples

import (
	"io"
)

type (
	// Bar is an example struct.
	Bar struct{}

	foo interface { // nolint: megacheck
		Foo(bar Bar) string
	}

	// Hello2 is an example interface.
	Hello2 interface {
		foo
		io.Reader
	}
)
