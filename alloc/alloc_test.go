package alloc

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestMakeAppend(t *testing.T) {
	testMakeAppend(t, 30, func(number int) int { return number })
	testMakeAppend(t, 30, func(number int) int { return number - 1 })
	testMakeAppend(t, 30, func(number int) int { return number - number%10 })
	testMakeAppend(t, 30, func(number int) int { return number - number%100 })
	testMakeAppend(t, 30, func(number int) int { return number - number%1000 })
}

func testMakeAppend(t *testing.T, power int, modifier func(number int) int) {
	for number := 1; number <= 1<<power; number *= 2 {
		length := modifier(number)

		slice := make([]bool, length)

		require.Len(t, slice, length)
		require.Equal(t, length, cap(slice))

		previous := slice

		slice = append(slice, true) //nolint:makezero
		require.NotSame(t, unsafe.SliceData(previous), unsafe.SliceData(slice))
	}
}
