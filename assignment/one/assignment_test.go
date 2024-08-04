package one

import (
	"testing"

	"github.com/akramarenkov/goresearch/internal/getenv"

	"github.com/stretchr/testify/require"
)

const defaultSize = 1000

func BenchmarkAssignmentSlice(b *testing.B) {
	size, err := getenv.Size(defaultSize)
	require.NoError(b, err)

	slice := make([]bool, size)

	for range b.N {
		for id := range size {
			slice[id] = true
		}
	}

	b.StopTimer()

	require.Len(b, slice, size)
}

func BenchmarkAppendSlice(b *testing.B) {
	size, err := getenv.Size(defaultSize)
	require.NoError(b, err)

	b.ResetTimer()

	slice := make([]bool, 0, size)

	for range b.N {
		slice = slice[:0]

		for range size {
			slice = append(slice, true)
		}
	}

	b.StopTimer()

	require.Len(b, slice, size)
}
