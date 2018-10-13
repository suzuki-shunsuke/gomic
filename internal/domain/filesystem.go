package domain

import (
	"io"
)

type (
	// FileSystem abstracts operations of the file system.
	FileSystem interface {
		Exist(string) bool
		MkdirAll(string) error
		Write(string, []byte) error
		GetWriteCloser(string) (io.WriteCloser, error)
		Getwd() (string, error)
	}
)
