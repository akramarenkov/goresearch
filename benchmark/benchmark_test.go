package benchmark

import (
	"testing"

	"github.com/akramarenkov/goresearch/internal/getenv"

	"github.com/stretchr/testify/require"
)

const defaultSize = 1000

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
	size, err := getenv.Size(defaultSize)
	require.NoError(b, err)

	b.ResetTimer()

	for range b.N {
		slice := make([]bool, size)

		_ = slice
	}
}

func BenchmarkMakeWithNLoopWithAssignmentSizeFromEnv(b *testing.B) {
	size, err := getenv.Size(defaultSize)
	require.NoError(b, err)

	b.ResetTimer()

	for range b.N {
		slice := make([]bool, size)

		for id := range slice {
			slice[id] = true
		}
	}
}
