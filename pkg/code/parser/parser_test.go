package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetModulePath(t *testing.T) {
	const gomod = `
module github.com/alkbt/gat

go 1.24.2

require golang.org/x/mod v0.24.0
`

	path, err := getModulePath(strings.NewReader(gomod))
	require.NoError(t, err)
	assert.Equal(t, "github.com/alkbt/gat", path)
}
