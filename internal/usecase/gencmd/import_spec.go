package gencmd

import (
	"fmt"
	"strings"
)

type (
	importSpec struct {
		name string
		path string
		str  string
	}
)

// Host implements domain.ImportSpec#Host .
func (spec importSpec) Host() string {
	h := strings.Split(spec.path, "/")[0]
	if strings.Contains(h, ".") {
		return h
	}
	return ""
}

// Name implements domain.ImportSpec#Name .
func (spec importSpec) Name() string {
	return spec.name
}

// Path implements domain.ImportSpec#Path .
func (spec importSpec) Path() string {
	return spec.path
}

// String implements domain.ImportSpec#String .
func (spec importSpec) String() string {
	if spec.str != "" {
		return spec.str
	}
	return fmt.Sprintf(`%s "%s"`, spec.Name(), spec.Path())
}
