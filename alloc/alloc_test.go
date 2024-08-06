package alloc

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestMake(t *testing.T) {
	testMake(t, 30, func(number int) int { return number })
	testMake(t, 30, func(number int) int { return number/2 + 1 })
}

func testMake(t *testing.T, power int, modifier func(number int) int) {
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

func TestMakeAppend(t *testing.T) {
	testMakeAppend(t, 30, func(number int) int { return number })
	testMakeAppend(t, 30, func(number int) int { return number/2 + 1 })
}

func testMakeAppend(t *testing.T, power int, modifier func(number int) int) {
	for number := 1; number <= 1<<power; number *= 2 {
		length := modifier(number)

		slice := make([]bool, length)

		for id := range slice {
			slice[id] = true
		}

		require.Len(t, slice, length)
		require.Equal(t, length, cap(slice))

		previous := slice

		slice = append(slice, true)
		require.NotSame(t, unsafe.SliceData(previous), unsafe.SliceData(slice))
	}
}
