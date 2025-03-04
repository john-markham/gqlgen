package imports

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/john-markham/gqlgen/internal/code"
)

func TestPrune(t *testing.T) {
	// prime the packages cache so that it's not considered uninitialized

	b, err := Prune("testdata/unused.go", mustReadFile("testdata/unused.go"), code.NewPackages())
	require.NoError(t, err)
	require.Equal(t, strings.ReplaceAll(string(mustReadFile("testdata/unused.expected.go")), "\r\n", "\n"), string(b))
}

func mustReadFile(filename string) []byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return b
}
