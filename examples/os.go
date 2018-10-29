// Package examples is an example of gomic.
package examples

import (
	"os"
)

type (
	// OS is an example interface.
	OS interface {
		Getwd() (string, error)
		Mkdir(name string, perm os.FileMode) error
	}
)
