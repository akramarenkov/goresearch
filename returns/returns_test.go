package returns

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type result struct {
	First  int
	Second int
	Third  int
	Fourth int
	Fifth  int
	Sixth  int
}

func doDirect(seed int) (int, int, int, int, int, int) {
	return seed, seed * 2, seed * 3, seed * 4, seed * 5, seed * 6
}

func doStruct(seed int) result {
	result := result{
		First:  seed,
		Second: seed * 2,
		Third:  seed * 3,
		Fourth: seed * 4,
		Fifth:  seed * 5,
		Sixth:  seed * 6,
	}

	return result
}

func doStructPointer(seed int) *result {
	result := &result{
		First:  seed,
		Second: seed * 2,
		Third:  seed * 3,
		Fourth: seed * 4,
		Fifth:  seed * 5,
		Sixth:  seed * 6,
	}

	return result
}

func doStructPointerPassthrough(seed int, result *result) {
	result.First = seed
	result.Second = seed * 2
	result.Third = seed * 3
	result.Fourth = seed * 4
	result.Fifth = seed * 5
	result.Sixth = seed * 6
}

func BenchmarkDirect(b *testing.B) {
	first := 0
	second := 0
	third := 0
	fourth := 0
	fifth := 0
	sixth := 0

	// b.N and require is used to prevent compiler optimizations
	for range b.N {
		first, second, third, fourth, fifth, sixth = doDirect(b.N)
	}

	b.StopTimer()

	require.NotNil(b, first)
	require.NotNil(b, second)
	require.NotNil(b, third)
	require.NotNil(b, fourth)
	require.NotNil(b, fifth)
	require.NotNil(b, sixth)
}

func BenchmarkStruct(b *testing.B) {
	result := result{}

	// b.N and require is used to prevent compiler optimizations
	for range b.N {
		result = doStruct(b.N)
	}

	b.StopTimer()

	require.NotNil(b, result.First)
	require.NotNil(b, result.Second)
	require.NotNil(b, result.Third)
	require.NotNil(b, result.Fourth)
	require.NotNil(b, result.Fifth)
	require.NotNil(b, result.Sixth)
}

func BenchmarkStructPointer(b *testing.B) {
	result := &result{}

	// b.N and require is used to prevent compiler optimizations
	for range b.N {
		result = doStructPointer(b.N)
	}

	b.StopTimer()

	require.NotNil(b, result.First)
	require.NotNil(b, result.Second)
	require.NotNil(b, result.Third)
	require.NotNil(b, result.Fourth)
	require.NotNil(b, result.Fifth)
	require.NotNil(b, result.Sixth)
}

func BenchmarkStructPointerPassthrough(b *testing.B) {
	result := &result{}

	// b.N and require is used to prevent compiler optimizations
	for range b.N {
		doStructPointerPassthrough(b.N, result)
	}

	b.StopTimer()

	require.NotNil(b, result.First)
	require.NotNil(b, result.Second)
	require.NotNil(b, result.Third)
	require.NotNil(b, result.Fourth)
	require.NotNil(b, result.Fifth)
	require.NotNil(b, result.Sixth)
}
