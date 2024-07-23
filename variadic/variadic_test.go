package variadic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func one(first int) int {
	return first
}

func oneVariadic(numbers ...int) int {
	return numbers[0]
}

func two(first, second int) int {
	return first | second
}

func twoVariadic(numbers ...int) int {
	return numbers[0] | numbers[1]
}

func three(first, second, third int) int {
	return first | second | third
}

func threeVariadic(numbers ...int) int {
	return numbers[0] | numbers[1] | numbers[2]
}

func four(first, second, third, fourth int) int {
	return first | second | third | fourth
}

func fourVariadic(numbers ...int) int {
	return numbers[0] | numbers[1] | numbers[2] | numbers[3]
}

func BenchmarkOne(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = one(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkOneVariadic(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = oneVariadic(b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkTwo(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = two(b.N, b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkTwoVariadic(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = twoVariadic(b.N, b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkThree(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = three(b.N, b.N, b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkThreeVariadic(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = threeVariadic(b.N, b.N, b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkFour(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = four(b.N, b.N, b.N, b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkFourVariadic(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = fourVariadic(b.N, b.N, b.N, b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}
