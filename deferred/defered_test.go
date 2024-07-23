package deferred

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func calculate(number int) int {
	return number + 1
}

func BenchmarkReference(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = calculate(number)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkNotDeferred(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	wrapper := func() {
		number = calculate(number)
	}

	main := func() {
		wrapper()
	}

	for range b.N {
		main()
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkDeferred(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	wrapper := func() {
		number = calculate(number)
	}

	main := func() {
		defer wrapper()
	}

	for range b.N {
		main()
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkReferenceIdle(b *testing.B) {
	for range b.N {
		_ = b.N
	}
}

func BenchmarkNotDeferredIdle(b *testing.B) {
	wrapper := func() {
	}

	main := func() {
		wrapper()
	}

	for range b.N {
		main()
	}
}

func BenchmarkDeferredIdle(b *testing.B) {
	wrapper := func() {
	}

	main := func() {
		defer wrapper()
	}

	for range b.N {
		main()
	}
}

func BenchmarkReference2(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = calculate(number)
		number = calculate(number)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkNotDeferred2(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	wrapper := func() {
		number = calculate(number)
	}

	main := func() {
		wrapper()
		wrapper()
	}

	for range b.N {
		main()
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkDeferred2(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	wrapper := func() {
		number = calculate(number)
	}

	main := func() {
		defer wrapper()
		defer wrapper()
	}

	for range b.N {
		main()
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkReference4(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkNotDeferred4(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	wrapper := func() {
		number = calculate(number)
	}

	main := func() {
		wrapper()
		wrapper()
		wrapper()
		wrapper()
	}

	for range b.N {
		main()
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkDeferred4(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	wrapper := func() {
		number = calculate(number)
	}

	main := func() {
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
	}

	for range b.N {
		main()
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkReference6(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkNotDeferred6(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	wrapper := func() {
		number = calculate(number)
	}

	main := func() {
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
	}

	for range b.N {
		main()
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkDeferred6(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	wrapper := func() {
		number = calculate(number)
	}

	main := func() {
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
	}

	for range b.N {
		main()
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkReference8(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkNotDeferred8(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	wrapper := func() {
		number = calculate(number)
	}

	main := func() {
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
	}

	for range b.N {
		main()
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkDeferred8(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	wrapper := func() {
		number = calculate(number)
	}

	main := func() {
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
	}

	for range b.N {
		main()
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkReference10(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkNotDeferred10(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	wrapper := func() {
		number = calculate(number)
	}

	main := func() {
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
	}

	for range b.N {
		main()
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkDeferred10(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	wrapper := func() {
		number = calculate(number)
	}

	main := func() {
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
	}

	for range b.N {
		main()
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkReference11(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkNotDeferred11(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	wrapper := func() {
		number = calculate(number)
	}

	main := func() {
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
		wrapper()
	}

	for range b.N {
		main()
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkDeferred11(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	wrapper := func() {
		number = calculate(number)
	}

	main := func() {
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
		defer wrapper()
	}

	for range b.N {
		main()
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}
