package parser

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestGetModulePathInvalidGoMod(t *testing.T) {
	const gomod = `
modulez github.com/alkbt/gat

go 1.24.2

require golang.org/x/mod v0.24.0
`

	path, err := getModulePath(strings.NewReader(gomod))
	require.Error(t, err)
	assert.Equal(t, "", path)
}

func TestGetModulePathReadError(t *testing.T) {
	expectedErrorText := "read error"
	rdr := MockedReader{}
	rdr.Mock.On("Read", mock.AnythingOfType("[]uint8")).Return(0, errors.New(expectedErrorText))

	path, err := getModulePath(&rdr)
	require.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrorText)
	assert.Empty(t, path)
}

type MockedReader struct {
	mock.Mock
}

func (m *MockedReader) Read(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}
