package clone

import (
	"slices"
	"testing"

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
	require.NotSame(b, original, copied)
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
	require.NotSame(b, original, copied)
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
	require.NotSame(b, original, copied)
}

func BenchmarkSlicesClone16B(b *testing.B) {
	original := make([]bool, 1<<4)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = slices.Clone(original)
	}

	b.StopTimer()

	require.NotSame(b, original, copied)
}

func BenchmarkThroughAppend16B(b *testing.B) {
	original := make([]bool, 1<<4)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughAppend(original)
	}

	b.StopTimer()

	require.NotSame(b, original, copied)
}

func BenchmarkThroughCopy16B(b *testing.B) {
	original := make([]bool, 1<<4)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughCopy(original)
	}

	b.StopTimer()

	require.NotSame(b, original, copied)
}

func BenchmarkSlicesClone32B(b *testing.B) {
	original := make([]bool, 1<<5)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = slices.Clone(original)
	}

	b.StopTimer()

	require.NotSame(b, original, copied)
}

func BenchmarkThroughAppend32B(b *testing.B) {
	original := make([]bool, 1<<5)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughAppend(original)
	}

	b.StopTimer()

	require.NotSame(b, original, copied)
}

func BenchmarkThroughCopy32B(b *testing.B) {
	original := make([]bool, 1<<5)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughCopy(original)
	}

	b.StopTimer()

	require.NotSame(b, original, copied)
}

func BenchmarkSlicesClone64B(b *testing.B) {
	original := make([]bool, 1<<6)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = slices.Clone(original)
	}

	b.StopTimer()

	require.NotSame(b, original, copied)
}

func BenchmarkThroughAppend64B(b *testing.B) {
	original := make([]bool, 1<<6)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughAppend(original)
	}

	b.StopTimer()

	require.NotSame(b, original, copied)
}

func BenchmarkThroughCopy64B(b *testing.B) {
	original := make([]bool, 1<<6)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughCopy(original)
	}

	b.StopTimer()

	require.NotSame(b, original, copied)
}

func BenchmarkSlicesClone128B(b *testing.B) {
	original := make([]bool, 1<<7)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = slices.Clone(original)
	}

	b.StopTimer()

	require.NotSame(b, original, copied)
}

func BenchmarkThroughAppend128B(b *testing.B) {
	original := make([]bool, 1<<7)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughAppend(original)
	}

	b.StopTimer()

	require.NotSame(b, original, copied)
}

func BenchmarkThroughCopy128B(b *testing.B) {
	original := make([]bool, 1<<7)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughCopy(original)
	}

	b.StopTimer()

	require.NotSame(b, original, copied)
}

func BenchmarkThroughAppend1KB(b *testing.B) {
	original := make([]bool, 1<<10)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughAppend(original)
	}

	b.StopTimer()

	require.NotSame(b, original, copied)
}

func BenchmarkThroughCopy1KB(b *testing.B) {
	original := make([]bool, 1<<10)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughCopy(original)
	}

	b.StopTimer()

	require.NotSame(b, original, copied)
}

func BenchmarkThroughAppend1MB(b *testing.B) {
	original := make([]bool, 1<<20)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughAppend(original)
	}

	b.StopTimer()

	require.NotSame(b, original, copied)
}

func BenchmarkThroughCopy1MB(b *testing.B) {
	original := make([]bool, 1<<20)

	original[0] = true
	original[len(original)-1] = true

	var copied []bool

	b.ResetTimer()

	for range b.N {
		copied = throughCopy(original)
	}

	b.StopTimer()

	require.NotSame(b, original, copied)
}
