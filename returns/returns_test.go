package returns

import (
	"math/rand"
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

func process(seed int) (int, int, int, int, int, int) {
	return seed, seed * 2, seed * 3, seed * 4, seed * 5, seed * 6
}

func processStruct(seed int) result {
	result := result{
		First: seed,
	}

	result.Second = seed * 2
	result.Third = seed * 3
	result.Fourth = seed * 4
	result.Fifth = seed * 5
	result.Sixth = seed * 6

	return result
}

func processStructPointer(seed int) *result {
	result := &result{
		First: seed,
	}

	result.Second = seed * 2
	result.Third = seed * 3
	result.Fourth = seed * 4
	result.Fifth = seed * 5
	result.Sixth = seed * 6

	return result
}

func processStructPointerPassthrough(seed int, result *result) {
	result.Second = seed * 2
	result.Third = seed * 3
	result.Fourth = seed * 4
	result.Fifth = seed * 5
	result.Sixth = seed * 6
}

func BenchmarkDirect(b *testing.B) {
	// seed and require is used to prevent compiler optimizations
	seed := rand.Int() //nolint:gosec

	b.ResetTimer()

	first := 0
	second := 0
	third := 0
	fourth := 0
	fifth := 0
	sixth := 0

	for range b.N {
		first, second, third, fourth, fifth, sixth = process(seed)
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
	// seed and require is used to prevent compiler optimizations
	seed := rand.Int() //nolint:gosec

	b.ResetTimer()

	result := result{}

	for range b.N {
		result = processStruct(seed)
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
	// seed and require is used to prevent compiler optimizations
	seed := rand.Int() //nolint:gosec

	b.ResetTimer()

	result := &result{}

	for range b.N {
		result = processStructPointer(seed)
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
	// seed and require is used to prevent compiler optimizations
	seed := rand.Int() //nolint:gosec

	b.ResetTimer()

	result := &result{}

	for range b.N {
		processStructPointerPassthrough(seed, result)
	}

	b.StopTimer()

	require.NotNil(b, result.First)
	require.NotNil(b, result.Second)
	require.NotNil(b, result.Third)
	require.NotNil(b, result.Fourth)
	require.NotNil(b, result.Fifth)
	require.NotNil(b, result.Sixth)
}
