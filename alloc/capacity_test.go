package alloc

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestMakeCapacity(t *testing.T) {
	testMakeCapacity(t, 30, func(number int) int { return number })
	testMakeCapacity(t, 30, func(number int) int { return number - 1 })
	testMakeCapacity(t, 30, func(number int) int { return number - number%10 })
	testMakeCapacity(t, 30, func(number int) int { return number - number%100 })
	testMakeCapacity(t, 30, func(number int) int { return number - number%1000 })
}

func testMakeCapacity(t *testing.T, power int, modifier func(number int) int) {
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

func TestAppendCapacity(t *testing.T) {
	testAppendCapacity(t, 30, func(number int) int { return number })
	testAppendCapacity(t, 30, func(number int) int { return number - 1 })
	testAppendCapacity(t, 30, func(number int) int { return number - number%10 })
	testAppendCapacity(t, 30, func(number int) int { return number - number%100 })
	testAppendCapacity(t, 30, func(number int) int { return number - number%1000 })
}

func testAppendCapacity(t *testing.T, power int, modifier func(number int) int) {
	equalities := 0
	discrepancy := 0

	same := 0
	notSame := 0

	for number := 1; number <= 1<<power; number *= 2 {
		length := modifier(number)

		slice := append([]byte(nil), make([]byte, length)...)

		if cap(slice) == length {
			equalities++
		} else {
			discrepancy++
		}

		require.Len(t, slice, length)

		previous := slice

		slice = append(slice, 1)

		if unsafe.SliceData(previous) == unsafe.SliceData(slice) {
			same++
		} else {
			notSame++
		}
	}

	require.NotZero(t, equalities)
	require.NotZero(t, discrepancy)
	require.NotZero(t, same)
	require.NotZero(t, notSame)
}
