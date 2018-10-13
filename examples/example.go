// Package examples is an example of gomic.
package examples

type (
	// Hello is an example interface.
	Hello interface {
		ExistFile(string) bool
		MkdirAll(string)
	}
)
