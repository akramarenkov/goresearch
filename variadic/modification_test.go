package variadic

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func modification(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	slices.Sort(numbers)

	sum := 0

	for id := range numbers {
		sum += numbers[id]
	}

	return sum
}

func TestModification(t *testing.T) {
	first := 5
	second := 4
	third := 3
	fourth := 2
	fifth := 1

	sum := modification(first, second, third, fourth, fifth)

	require.Equal(t, 15, sum)
	require.Equal(t, 5, first)
	require.Equal(t, 4, second)
	require.Equal(t, 3, third)
	require.Equal(t, 2, fourth)
	require.Equal(t, 1, fifth)
}

func TestModificationSlice(t *testing.T) {
	numbers := []int{5, 4, 3, 2, 1}

	sum := modification(numbers...)

	require.Equal(t, 15, sum)
	require.Equal(t, []int{1, 2, 3, 4, 5}, numbers)
}

func BenchmarkModification(b *testing.B) {
	first := 5
	second := 4
	third := 3
	fourth := 2
	fifth := 1

	// sum and require is used to prevent compiler optimizations
	sum := 0

	b.ResetTimer()

	for range b.N {
		sum = modification(first, second, third, fourth, fifth)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}

func BenchmarkModificationSlice(b *testing.B) {
	numbers := []int{5, 4, 3, 2, 1}

	// sum and require is used to prevent compiler optimizations
	sum := 0

	b.ResetTimer()

	for range b.N {
		sum = modification(numbers...)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, sum)
}
