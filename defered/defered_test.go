package defered

import "testing"

func calculate(number int) int {
	return number + 2*number + 1
}

func BenchmarkReference(b *testing.B) {
	number := 0

	for range b.N {
		number = calculate(number)
	}
}

func BenchmarkNotDefered(b *testing.B) {
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
}

func BenchmarkDefered(b *testing.B) {
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
}

func BenchmarkReferenceIdle(b *testing.B) {
	for range b.N {
	}
}

func BenchmarkNotDeferedIdle(b *testing.B) {
	wrapper := func() {
	}

	main := func() {
		wrapper()
	}

	for range b.N {
		main()
	}
}

func BenchmarkDeferedIdle(b *testing.B) {
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
	number := 0

	for range b.N {
		number = calculate(number)
		number = calculate(number)
	}
}

func BenchmarkNotDefered2(b *testing.B) {
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
}

func BenchmarkDefered2(b *testing.B) {
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
}

func BenchmarkReference4(b *testing.B) {
	number := 0

	for range b.N {
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
	}
}

func BenchmarkNotDefered4(b *testing.B) {
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
}

func BenchmarkDefered4(b *testing.B) {
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
}

func BenchmarkReference6(b *testing.B) {
	number := 0

	for range b.N {
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
		number = calculate(number)
	}
}

func BenchmarkNotDefered6(b *testing.B) {
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
}

func BenchmarkDefered6(b *testing.B) {
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
}

func BenchmarkReference8(b *testing.B) {
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
}

func BenchmarkNotDefered8(b *testing.B) {
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
}

func BenchmarkDefered8(b *testing.B) {
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
}

func BenchmarkReference10(b *testing.B) {
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
}

func BenchmarkNotDefered10(b *testing.B) {
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
}

func BenchmarkDefered10(b *testing.B) {
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
}

func BenchmarkReference11(b *testing.B) {
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
}

func BenchmarkNotDefered11(b *testing.B) {
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
}

func BenchmarkDefered11(b *testing.B) {
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
}
