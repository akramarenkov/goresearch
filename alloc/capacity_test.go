package alloc

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestCapacity(t *testing.T) {
	testCapacity(t, 30, func(number int) int { return number })
	testCapacity(t, 30, func(number int) int { return number - 1 })
	testCapacity(t, 30, func(number int) int { return number - number%10 })
	testCapacity(t, 30, func(number int) int { return number - number%100 })
	testCapacity(t, 30, func(number int) int { return number - number%1000 })
}

func testCapacity(t *testing.T, power int, modifier func(number int) int) {
	for number := 1; number <= 1<<power; number *= 2 {
		length := modifier(number)

		slice := make([]byte, length)

		require.Len(t, slice, length)
		require.Equal(t, length, cap(slice))

		previous := slice

		slice = append(slice, 1) //nolint:makezero
		require.NotSame(t, unsafe.SliceData(previous), unsafe.SliceData(slice))
	}
}
