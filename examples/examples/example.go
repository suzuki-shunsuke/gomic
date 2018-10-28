// Package examples is an example of gomic.
package examples

import (
	"github.com/suzuki-shunsuke/gomic/examples/examples/examples"
)

type (
	// ChildHello is an example interface.
	ChildHello interface {
		Foo() examples.ChildFoo
	}
)
