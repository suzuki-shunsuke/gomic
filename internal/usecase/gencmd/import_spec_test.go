package gencmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

func Test_ImportSpecsAdd(t *testing.T) {
	specs := NewImportSpecs()
	name := "rterror"
	path := "github.com/suzuki-shunsuke/rterror"

	spec, err := specs.Add(importSpec{name: name, path: path})
	assert.Nil(t, err)
	assert.Equal(t, name, spec.Name())
	assert.Equal(t, path, spec.Path())
	assert.Exactly(t, map[string]domain.ImportSpec{name: importSpec{name: name, path: path}}, specs.Names())
	assert.Exactly(t, map[string]domain.ImportSpec{path: importSpec{name: name, path: path}}, specs.Paths())

	spec, err = specs.Add(importSpec{name: name, path: path})
	assert.Nil(t, err)
	assert.Equal(t, name, spec.Name())
	assert.Equal(t, path, spec.Path())
	assert.Exactly(t, map[string]domain.ImportSpec{name: importSpec{name: name, path: path}}, specs.Names())
	assert.Exactly(t, map[string]domain.ImportSpec{path: importSpec{name: name, path: path}}, specs.Paths())

	spec, err = specs.Add(importSpec{name: name, path: "io"})
	assert.Nil(t, err)
	assert.Equal(t, "rterror0", spec.Name())
	assert.Equal(t, "io", spec.Path())

	spec, err = specs.Add(importSpec{name: "rterror1", path: "io"})
	assert.Nil(t, err)
	assert.Equal(t, "rterror0", spec.Name())
	assert.Equal(t, "io", spec.Path())
}

func Test_importSpecHost(t *testing.T) {
	spec := importSpec{}
	assert.Equal(t, "", spec.Host())
	spec.path = "io"
	assert.Equal(t, "", spec.Host())
	spec.path = "github.com/suzuki-shunsuke/gomic"
	assert.Equal(t, "github.com", spec.Host())
}

func Test_importSpecString(t *testing.T) {
	spec := importSpec{name: "io", path: "io"}
	assert.Equal(t, `io "io"`, spec.String())
	spec.str = `"io"`
	assert.Equal(t, `"io"`, spec.String())
}
