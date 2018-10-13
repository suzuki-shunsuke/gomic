package gencmd

import (
	"fmt"
	"strings"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

type (
	importSpec struct {
		name string
		path string
		str  string
	}

	// ImportSpecs implements domain.ImportSpecs .
	ImportSpecs struct {
		names map[string]domain.ImportSpec
		paths map[string]domain.ImportSpec
	}
)

// NewImportSpecs is a constructor of ImportSpecs .
func NewImportSpecs() domain.ImportSpecs {
	return &ImportSpecs{
		names: map[string]domain.ImportSpec{},
		paths: map[string]domain.ImportSpec{},
	}
}

// Add implements domain.ImportSpecs#Add .
func (specs *ImportSpecs) Add(spec domain.ImportSpec) (domain.ImportSpec, error) {
	if s, ok := specs.paths[spec.Path()]; ok {
		return s, nil
	}
	if s, ok := specs.names[spec.Name()]; ok {
		if spec.Path() == s.Path() {
			return spec, nil
		}
		for i := 0; i < 100; i++ {
			name := fmt.Sprintf("%s%d", spec.Name(), i)
			if _, ok := specs.names[name]; ok {
				continue
			}
			s := importSpec{name: name, path: spec.Path()}
			specs.names[name] = s
			specs.paths[s.Path()] = s
			return s, nil
		}
		return nil, fmt.Errorf("failed to add import %s %s", spec.Name(), spec.Path())
	}
	specs.names[spec.Name()] = spec
	specs.paths[spec.Path()] = spec
	return spec, nil
}

// Names implements domain.ImportSpecs#Names .
func (specs ImportSpecs) Names() map[string]domain.ImportSpec {
	return specs.names
}

// Paths implements domain.ImportSpecs#Paths .
func (specs ImportSpecs) Paths() map[string]domain.ImportSpec {
	return specs.paths
}

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
