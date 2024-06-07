package benchmark

import (
	"os"
	"strconv"
	"testing"

	"github.com/akramarenkov/goresearch/internal/consts"

	"github.com/stretchr/testify/require"
)

func BenchmarkMakeWithoutNLoop(b *testing.B) {
	slice := make([]bool, 1000)

	for id := range slice {
		slice[id] = true
	}
}

func BenchmarkMakeWithNLoop(b *testing.B) {
	for run := 0; run < b.N; run++ {
		slice := make([]bool, 1000)

		for id := range slice {
			slice[id] = true
		}
	}
}

func BenchmarkMakeWithNLoopSizeFromEnv(b *testing.B) {
	env := os.Getenv(consts.EnvPrefix + "SIZE")

	size, err := strconv.Atoi(env)
	require.NoError(b, err)

	for run := 0; run < b.N; run++ {
		slice := make([]bool, size)

		for id := range slice {
			slice[id] = true
		}
	}
}

func BenchmarkAssignmentSlice(b *testing.B) {
	env := os.Getenv(consts.EnvPrefix + "SIZE")

	size, err := strconv.Atoi(env)
	require.NoError(b, err)

	slice := make([]bool, size)

	for run := 0; run < b.N; run++ {
		for id := range cap(slice) {
			slice[id] = true
		}
	}

	require.Len(b, slice, size)
	require.Equal(b, size, cap(slice))
}

func BenchmarkAppendSlice(b *testing.B) {
	env := os.Getenv(consts.EnvPrefix + "SIZE")

	size, err := strconv.Atoi(env)
	require.NoError(b, err)

	slice := make([]bool, 0, size)

	for run := 0; run < b.N; run++ {
		slice = slice[:0]

		for range cap(slice) {
			slice = append(slice, true)
		}
	}

	require.Len(b, slice, size)
	require.Equal(b, size, cap(slice))
}
