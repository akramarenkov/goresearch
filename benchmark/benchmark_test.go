package benchmark

import (
	"os"
	"strconv"
	"testing"

	"github.com/akramarenkov/goresearch/internal/consts"
	"github.com/stretchr/testify/require"
)

const defaultSize = 1000

func getSizeFromEnv(b *testing.B) int {
	env := os.Getenv(consts.EnvPrefix + "SIZE")

	if env == "" {
		return defaultSize
	}

	size, err := strconv.Atoi(env)
	require.NoError(b, err)

	return size
}

func BenchmarkMakeWithoutNLoop(*testing.B) {
	slice := make([]bool, defaultSize)

	_ = slice
}

func BenchmarkMakeWithoutNLoopWithAssignment(*testing.B) {
	slice := make([]bool, defaultSize)

	for id := range slice {
		slice[id] = true
	}
}

func BenchmarkMakeWithNLoop(b *testing.B) {
	for range b.N {
		slice := make([]bool, defaultSize)

		_ = slice
	}
}

func BenchmarkMakeWithNLoopWithAssignment(b *testing.B) {
	for range b.N {
		slice := make([]bool, defaultSize)

		for id := range slice {
			slice[id] = true
		}
	}
}

func BenchmarkMakeWithNLoopScopeEscape(b *testing.B) {
	var slice []bool

	for range b.N {
		slice = make([]bool, defaultSize)

		_ = slice
	}
}

func BenchmarkMakeWithNLoopWithAssignmentScopeEscape(b *testing.B) {
	var slice []bool

	for range b.N {
		slice = make([]bool, defaultSize)

		for id := range slice {
			slice[id] = true
		}
	}
}

func BenchmarkMakeWithNLoopSizeFromEnv(b *testing.B) {
	size := getSizeFromEnv(b)

	b.ResetTimer()

	for range b.N {
		slice := make([]bool, size)

		_ = slice
	}
}

func BenchmarkMakeWithNLoopWithAssignmentSizeFromEnv(b *testing.B) {
	size := getSizeFromEnv(b)

	b.ResetTimer()

	for range b.N {
		slice := make([]bool, size)

		for id := range slice {
			slice[id] = true
		}
	}
}

func BenchmarkAssignmentSlice(b *testing.B) {
	size := getSizeFromEnv(b)

	b.ResetTimer()

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
	size := getSizeFromEnv(b)

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
