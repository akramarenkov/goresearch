// The main package which was originally intended to be the only one.
package assign

import (
	"testing"

	"github.com/akramarenkov/goresearch/internal/getenv"

	"github.com/stretchr/testify/require"
)

const defaultSize = 1000

func BenchmarkAssignmentSlice(b *testing.B) {
	size, err := getenv.Size(defaultSize)
	require.NoError(b, err)

	b.ResetTimer()

	slice := make([]byte, size)

	for range b.N {
		for id := range size {
			slice[id] = 1
		}
	}

	b.StopTimer()

	require.Len(b, slice, size)
}

func BenchmarkAppendSlice(b *testing.B) {
	size, err := getenv.Size(defaultSize)
	require.NoError(b, err)

	b.ResetTimer()

	slice := make([]byte, 0, size)

	for range b.N {
		slice = slice[:0]

		for range size {
			slice = append(slice, 1)
		}
	}

	b.StopTimer()

	require.Len(b, slice, size)
}
