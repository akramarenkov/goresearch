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

func benchmarkBytesSlicesClone(b *testing.B, length int) {
	original := make([]byte, length)

	var copied []byte

	b.ResetTimer()

	for range b.N {
		copied = slices.Clone(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func benchmarkBytesThroughAppend(b *testing.B, length int) {
	original := make([]byte, length)

	var copied []byte

	b.ResetTimer()

	for range b.N {
		copied = throughAppend(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func benchmarkBytesThroughCopy(b *testing.B, length int) {
	original := make([]byte, length)

	var copied []byte

	b.ResetTimer()

	for range b.N {
		copied = throughCopy(original)
	}

	b.StopTimer()

	require.NotSame(b, unsafe.SliceData(original), unsafe.SliceData(copied))
}

func Benchmark_16B_SlicesClone(b *testing.B) {
	benchmarkBytesSlicesClone(b, 1<<4)
}

func Benchmark_16B_ThroughAppend(b *testing.B) {
	benchmarkBytesThroughAppend(b, 1<<4)
}

func Benchmark_16B_ThroughCopy(b *testing.B) {
	benchmarkBytesThroughCopy(b, 1<<4)
}

func Benchmark_32B_SlicesClone(b *testing.B) {
	benchmarkBytesSlicesClone(b, 1<<5)
}

func Benchmark_32B_ThroughAppend(b *testing.B) {
	benchmarkBytesThroughAppend(b, 1<<5)
}

func Benchmark_32B_ThroughCopy(b *testing.B) {
	benchmarkBytesThroughCopy(b, 1<<5)
}

func Benchmark_64B_SlicesClone(b *testing.B) {
	benchmarkBytesSlicesClone(b, 1<<6)
}

func Benchmark_64B_ThroughAppend(b *testing.B) {
	benchmarkBytesThroughAppend(b, 1<<6)
}

func Benchmark_64B_ThroughCopy(b *testing.B) {
	benchmarkBytesThroughCopy(b, 1<<6)
}

func Benchmark_128B_SlicesClone(b *testing.B) {
	benchmarkBytesSlicesClone(b, 1<<7)
}

func Benchmark_128B_ThroughAppend(b *testing.B) {
	benchmarkBytesThroughAppend(b, 1<<7)
}

func Benchmark_128B_ThroughCopy(b *testing.B) {
	benchmarkBytesThroughCopy(b, 1<<7)
}

func Benchmark_1KB_SlicesClone(b *testing.B) {
	benchmarkBytesSlicesClone(b, 1<<10)
}

func Benchmark_1KB_ThroughAppend(b *testing.B) {
	benchmarkBytesThroughAppend(b, 1<<10)
}

func Benchmark_1KB_ThroughCopy(b *testing.B) {
	benchmarkBytesThroughCopy(b, 1<<10)
}

func Benchmark_1MB_SlicesClone(b *testing.B) {
	benchmarkBytesSlicesClone(b, 1<<20)
}

func Benchmark_1MB_ThroughAppend(b *testing.B) {
	benchmarkBytesThroughAppend(b, 1<<20)
}

func Benchmark_1MB_ThroughCopy(b *testing.B) {
	benchmarkBytesThroughCopy(b, 1<<20)
}
