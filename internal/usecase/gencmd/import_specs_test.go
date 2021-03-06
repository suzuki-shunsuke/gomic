package gencmd

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

func Test_ImportSpecsAdd(t *testing.T) {
	specs := NewImportSpecs()
	name := "rterror"
	path := "github.com/suzuki-shunsuke/rterror"

	spec, err := specs.Add(importSpec{name: name, path: path})
	require.Nil(t, err)
	require.Equal(t, name, spec.Name())
	require.Equal(t, path, spec.Path())
	require.Exactly(t, map[string]domain.ImportSpec{name: importSpec{name: name, path: path}}, specs.Names())
	require.Exactly(t, map[string]domain.ImportSpec{path: importSpec{name: name, path: path}}, specs.Paths())

	spec, err = specs.Add(importSpec{name: name, path: path})
	require.Nil(t, err)
	require.Equal(t, name, spec.Name())
	require.Equal(t, path, spec.Path())
	require.Exactly(t, map[string]domain.ImportSpec{name: importSpec{name: name, path: path}}, specs.Names())
	require.Exactly(t, map[string]domain.ImportSpec{path: importSpec{name: name, path: path}}, specs.Paths())

	spec, err = specs.Add(importSpec{name: name, path: "io"})
	require.Nil(t, err)
	require.Equal(t, "rterror0", spec.Name())
	require.Equal(t, "io", spec.Path())

	spec, err = specs.Add(importSpec{name: "rterror1", path: "io"})
	require.Nil(t, err)
	require.Equal(t, "rterror0", spec.Name())
	require.Equal(t, "io", spec.Path())
}
