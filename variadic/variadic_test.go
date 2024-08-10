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

func six(first, second, third, fourth, fifth, sixth int) int {
	return first | second | third | fourth | fifth | sixth
}

func sixVariadic(numbers ...int) int {
	return numbers[0] | numbers[1] | numbers[2] | numbers[3] | numbers[4] | numbers[5]
}

func eight(first, second, third, fourth, fifth, sixth, seventh, eighth int) int {
	return first | second | third | fourth | fifth | sixth | seventh | eighth
}

func eightVariadic(numbers ...int) int {
	return numbers[0] | numbers[1] | numbers[2] | numbers[3] | numbers[4] | numbers[5] |
		numbers[6] | numbers[7]
}

func ten(
	first,
	second,
	third,
	fourth,
	fifth,
	sixth,
	seventh,
	eighth,
	ninth,
	tenth int,
) int {
	return first | second | third | fourth | fifth | sixth | seventh | eighth |
		ninth | tenth
}

func tenVariadic(numbers ...int) int {
	return numbers[0] | numbers[1] | numbers[2] | numbers[3] | numbers[4] | numbers[5] |
		numbers[6] | numbers[7] | numbers[8] | numbers[9]
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

func BenchmarkOneVariadicSlice(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	args := make([]int, 1)

	for range b.N {
		args[0] = b.N

		number = oneVariadic(args...)
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

func BenchmarkTwoVariadicSlice(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	args := make([]int, 2)

	for range b.N {
		args[0] = b.N
		args[1] = b.N

		number = twoVariadic(args...)
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

func BenchmarkThreeVariadicSlice(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	args := make([]int, 3)

	for range b.N {
		args[0] = b.N
		args[1] = b.N
		args[2] = b.N

		number = threeVariadic(args...)
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

func BenchmarkFourVariadicSlice(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	args := make([]int, 4)

	for range b.N {
		args[0] = b.N
		args[1] = b.N
		args[2] = b.N
		args[3] = b.N

		number = fourVariadic(args...)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkSix(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = six(b.N, b.N, b.N, b.N, b.N, b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkSixVariadic(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = sixVariadic(b.N, b.N, b.N, b.N, b.N, b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkSixVariadicSlice(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	args := make([]int, 6)

	for range b.N {
		args[0] = b.N
		args[1] = b.N
		args[2] = b.N
		args[3] = b.N
		args[4] = b.N
		args[5] = b.N

		number = sixVariadic(args...)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkEight(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = eight(b.N, b.N, b.N, b.N, b.N, b.N, b.N, b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkEightVariadic(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = eightVariadic(b.N, b.N, b.N, b.N, b.N, b.N, b.N, b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkEightVariadicSlice(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	args := make([]int, 8)

	for range b.N {
		args[0] = b.N
		args[1] = b.N
		args[2] = b.N
		args[3] = b.N
		args[4] = b.N
		args[5] = b.N
		args[6] = b.N
		args[7] = b.N

		number = eightVariadic(args...)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkTen(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = ten(b.N, b.N, b.N, b.N, b.N, b.N, b.N, b.N, b.N, b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkTenVariadic(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = tenVariadic(b.N, b.N, b.N, b.N, b.N, b.N, b.N, b.N, b.N, b.N)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkTenVariadicSlice(b *testing.B) {
	// number, b.N and require is used to prevent compiler optimizations
	number := 0

	args := make([]int, 10)

	for range b.N {
		args[0] = b.N
		args[1] = b.N
		args[2] = b.N
		args[3] = b.N
		args[4] = b.N
		args[5] = b.N
		args[6] = b.N
		args[7] = b.N
		args[8] = b.N
		args[9] = b.N

		number = tenVariadic(args...)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}
