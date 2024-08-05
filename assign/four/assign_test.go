// The difference with the third package is the reverse order of writing the
// Benchmark* functions.
package four

import (
	"testing"

	"github.com/akramarenkov/goresearch/internal/getenv"

	"github.com/stretchr/testify/require"
)

const defaultSize = 1000

func BenchmarkAppendSlice(b *testing.B) {
	size, err := getenv.Size(defaultSize)
	require.NoError(b, err)

	slice := make([]bool, 0, size)

	for range b.N {
		slice = slice[:0]

		for range size {
			slice = append(slice, true)
		}
	}

	require.Len(b, slice, size)
}

func BenchmarkAssignmentSlice(b *testing.B) {
	size, err := getenv.Size(defaultSize)
	require.NoError(b, err)

	slice := make([]bool, size)

	for range b.N {
		for id := range size {
			slice[id] = true
		}
	}

	require.Len(b, slice, size)
}
