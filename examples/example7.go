package examples

import (
	"io"

	"github.com/suzuki-shunsuke/gomic/examples/examples"
)

type (
	// Hello7 is an example interface.
	Hello7 interface {
		io.Reader
		examples.ChildHello
		Child() examples.ChildHello
	}
)
