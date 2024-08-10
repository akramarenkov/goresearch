package deferred

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func calculate(number int) int {
	return number + 2*number + 1
}

func BenchmarkIdleReference(b *testing.B) {
	for range b.N {
		_ = b.N
	}
}

func BenchmarkIdleNotDeferred(b *testing.B) {
	wrapper := func() {
	}

	main := func() {
		wrapper()
	}

	for range b.N {
		main()
	}
}

func BenchmarkIdleDeferred(b *testing.B) {
	wrapper := func() {
	}

	main := func() {
		defer wrapper()
	}

	for range b.N {
		main()
	}
}

func BenchmarkOneReference(b *testing.B) {
	// number and require is used to prevent compiler optimizations
	number := 0

	for range b.N {
		number = calculate(number)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, number)
}

func BenchmarkOneNotDeferred(b *testing.B) {
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

func BenchmarkOneDeferred(b *testing.B) {
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

func BenchmarkTwoReference(b *testing.B) {
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

func BenchmarkTwoNotDeferred(b *testing.B) {
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

func BenchmarkTwoDeferred(b *testing.B) {
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

func BenchmarkFourReference(b *testing.B) {
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

func BenchmarkFourNotDeferred(b *testing.B) {
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

func BenchmarkFourDeferred(b *testing.B) {
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

func BenchmarkSixReference(b *testing.B) {
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

func BenchmarkSixNotDeferred(b *testing.B) {
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

func BenchmarkSixDeferred(b *testing.B) {
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

func BenchmarkEightReference(b *testing.B) {
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

func BenchmarkEightNotDeferred(b *testing.B) {
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

func BenchmarkEightDeferred(b *testing.B) {
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

func BenchmarkTenReference(b *testing.B) {
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

func BenchmarkTenNotDeferred(b *testing.B) {
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

func BenchmarkTenDeferred(b *testing.B) {
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

func BenchmarkElevenReference(b *testing.B) {
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

func BenchmarkElevenNotDeferred(b *testing.B) {
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

func BenchmarkElevenDeferred(b *testing.B) {
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
