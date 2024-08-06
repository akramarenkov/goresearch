package alloc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAllocCap(t *testing.T) {
	testAllocCap(t, 32, func(number int) int { return number })
	testAllocCap(t, 32, func(number int) int { return number/2 + 1 })
}

func testAllocCap(t *testing.T, power int, modifier func(number int) int) {
	for number := 1; number <= 1<<power; number *= 2 {
		length := modifier(number)

		slice := make([]bool, length)

		for id := range slice {
			slice[id] = true
		}

		require.Len(t, slice, length)
		require.Equal(t, length, cap(slice))
	}
}
