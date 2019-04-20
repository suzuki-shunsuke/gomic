package gencmd

import (
	"go/parser"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_toString(t *testing.T) {
	a := "[]string"
	arr, err := parser.ParseExpr(a)
	if err != nil {
		t.Fatal(err)
	}
	s, err := toString(arr)
	require.Nil(t, err)
	require.Equal(t, a, s)
}

func Test_isPublicIdent(t *testing.T) {
	require.False(t, isPublicIdent(""))
	require.False(t, isPublicIdent("foo"))
	require.True(t, isPublicIdent("Foo"))
}
