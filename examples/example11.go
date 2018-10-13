package examples

import (
	"github.com/suzuki-shunsuke/rterror"
)

type (
	// Hello11 is an example interface.
	Hello11 interface {
		Hello() rterror.RuntimeError
	}
)
