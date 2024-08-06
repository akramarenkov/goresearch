package clone

import (
	"slices"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func throughAppend[Type any](slice []Type) []Type {
	return append([]Type(nil), slice...)
}

func throughCopy[Type any](slice []Type) []Type {
	copied := make([]Type, len(slice))

	copy(copied, slice)

	return copied
}

func BenchmarkSlicesClone(b *testing.B) {
	expected := []int{0, 1, 2, 3, 4, 5}
	original := []int{0, 1, 2, 3, 4, 5}

	var copied []int

	b.ResetTimer()

	for range b.N {
		copied = slices.Clone(original)
	}

	b.StopTimer()

	require.Equal(b, expected, copied)
	require.Equal(b, expected, original)
	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func BenchmarkThroughAppend(b *testing.B) {
	expected := []int{0, 1, 2, 3, 4, 5}
	original := []int{0, 1, 2, 3, 4, 5}

	var copied []int

	b.ResetTimer()

	for range b.N {
		copied = throughAppend(original)
	}

	b.StopTimer()

	require.Equal(b, expected, copied)
	require.Equal(b, expected, original)
	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func BenchmarkThroughCopy(b *testing.B) {
	expected := []int{0, 1, 2, 3, 4, 5}
	original := []int{0, 1, 2, 3, 4, 5}

	var copied []int

	b.ResetTimer()

	for range b.N {
		copied = throughCopy(original)
	}

	b.StopTimer()

	require.Equal(b, expected, copied)
	require.Equal(b, expected, original)
	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_16B_SlicesClone(b *testing.B) {
	original := make([]bool, 1<<4)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = slices.Clone(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_16B_ThroughAppend(b *testing.B) {
	original := make([]bool, 1<<4)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughAppend(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_16B_ThroughCopy(b *testing.B) {
	original := make([]bool, 1<<4)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughCopy(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_32B_SlicesClone(b *testing.B) {
	original := make([]bool, 1<<5)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = slices.Clone(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_32B_ThroughAppend(b *testing.B) {
	original := make([]bool, 1<<5)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughAppend(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_32B_ThroughCopy(b *testing.B) {
	original := make([]bool, 1<<5)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughCopy(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_64B_SlicesClone(b *testing.B) {
	original := make([]bool, 1<<6)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = slices.Clone(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_64B_ThroughAppend(b *testing.B) {
	original := make([]bool, 1<<6)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughAppend(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_64B_ThroughCopy(b *testing.B) {
	original := make([]bool, 1<<6)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughCopy(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_128B_SlicesClone(b *testing.B) {
	original := make([]bool, 1<<7)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = slices.Clone(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_128B_ThroughAppend(b *testing.B) {
	original := make([]bool, 1<<7)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughAppend(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_128B_ThroughCopy(b *testing.B) {
	original := make([]bool, 1<<7)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughCopy(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_1KB_SlicesClone(b *testing.B) {
	original := make([]bool, 1<<10)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = slices.Clone(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_1KB_ThroughAppend(b *testing.B) {
	original := make([]bool, 1<<10)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughAppend(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_1KB_ThroughCopy(b *testing.B) {
	original := make([]bool, 1<<10)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughCopy(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_1MB_SlicesClone(b *testing.B) {
	original := make([]bool, 1<<20)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = slices.Clone(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_1MB_ThroughAppend(b *testing.B) {
	original := make([]bool, 1<<20)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughAppend(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_1MB_ThroughCopy(b *testing.B) {
	original := make([]bool, 1<<20)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughCopy(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}
