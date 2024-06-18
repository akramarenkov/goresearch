package benchmark

import (
	"os"
	"strconv"
	"testing"

	"github.com/akramarenkov/goresearch/internal/consts"

	"github.com/stretchr/testify/require"
)

func BenchmarkMakeWithoutNLoop(*testing.B) {
	slice := make([]bool, 1000)

	_ = slice
}

func BenchmarkMakeWithoutNLoopWithAssignment(*testing.B) {
	slice := make([]bool, 1000)

	for id := range slice {
		slice[id] = true
	}
}

func BenchmarkMakeWithNLoop(b *testing.B) {
	for range b.N {
		slice := make([]bool, 1000)

		_ = slice
	}
}

func BenchmarkMakeWithNLoopWithAssignment(b *testing.B) {
	for range b.N {
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

	b.ResetTimer()

	for range b.N {
		slice := make([]bool, size)

		_ = slice
	}
}

func BenchmarkMakeWithNLoopSizeFromEnvWithAssignment(b *testing.B) {
	env := os.Getenv(consts.EnvPrefix + "SIZE")

	size, err := strconv.Atoi(env)
	require.NoError(b, err)

	b.ResetTimer()

	for range b.N {
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

	b.ResetTimer()

	for range b.N {
		for id := range size {
			slice[id] = true
		}
	}

	b.StopTimer()

	require.Len(b, slice, size)
	require.Equal(b, size, cap(slice))
}

func BenchmarkAssignmentSliceUseCap(b *testing.B) {
	env := os.Getenv(consts.EnvPrefix + "SIZE")

	size, err := strconv.Atoi(env)
	require.NoError(b, err)

	slice := make([]bool, size)

	b.ResetTimer()

	for range b.N {
		for id := range cap(slice) {
			slice[id] = true
		}
	}

	b.StopTimer()

	require.Len(b, slice, size)
	require.Equal(b, size, cap(slice))
}

func BenchmarkAssignmentSliceConsiderMake(b *testing.B) {
	env := os.Getenv(consts.EnvPrefix + "SIZE")

	size, err := strconv.Atoi(env)
	require.NoError(b, err)

	b.ResetTimer()

	slice := make([]bool, size)

	for range b.N {
		for id := range size {
			slice[id] = true
		}
	}

	b.StopTimer()

	require.Len(b, slice, size)
	require.Equal(b, size, cap(slice))
}

func BenchmarkAppendSlice(b *testing.B) {
	env := os.Getenv(consts.EnvPrefix + "SIZE")

	size, err := strconv.Atoi(env)
	require.NoError(b, err)

	slice := make([]bool, 0, size)

	b.ResetTimer()

	for range b.N {
		slice = slice[:0]

		for range size {
			slice = append(slice, true)
		}
	}

	b.StopTimer()

	require.Len(b, slice, size)
	require.Equal(b, size, cap(slice))
}

func BenchmarkAppendSliceUseCap(b *testing.B) {
	env := os.Getenv(consts.EnvPrefix + "SIZE")

	size, err := strconv.Atoi(env)
	require.NoError(b, err)

	slice := make([]bool, 0, size)

	b.ResetTimer()

	for range b.N {
		slice = slice[:0]

		// If you uncomment these lines, then all three benchmarks from
		// the BenchmarkAppendSlice* group can begin to run at the same speed equal to
		// the speed of BenchmarkAppendSlice
		/*if cap(slice) != size {
			b.FailNow()
		}*/

		for range cap(slice) {
			slice = append(slice, true)
		}
	}

	b.StopTimer()

	require.Len(b, slice, size)
	require.Equal(b, size, cap(slice))
}

func BenchmarkAppendSliceConsiderMake(b *testing.B) {
	env := os.Getenv(consts.EnvPrefix + "SIZE")

	size, err := strconv.Atoi(env)
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
	require.Equal(b, size, cap(slice))
}
