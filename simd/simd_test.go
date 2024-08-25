package simd

import (
	"testing"

	"github.com/klauspost/cpuid/v2"
	"github.com/stretchr/testify/require"
)

func TestSlowMultiplication(t *testing.T) {
	require.True(t, cpuid.CPU.Supports(cpuid.AVX2))

	for first := range uint64(1 << 10) {
		for second := range uint64(1 << 10) {
			require.Equal(t, first*second, SlowMultiplication(first, second))
			require.Equal(t, first*second, SlowMultiplicationSIMD(first, second))
		}
	}
}

func BenchmarkSlowMultiplication_64(b *testing.B) {
	benchmarkSlowMultiplication(b, 1<<6)
}

func BenchmarkSlowMultiplicationSIMD_64(b *testing.B) {
	benchmarkSlowMultiplicationSIMD(b, 1<<6)
}

func BenchmarkSlowMultiplication_128(b *testing.B) {
	benchmarkSlowMultiplication(b, 1<<7)
}

func BenchmarkSlowMultiplicationSIMD_128(b *testing.B) {
	benchmarkSlowMultiplicationSIMD(b, 1<<7)
}

func BenchmarkSlowMultiplication_256(b *testing.B) {
	benchmarkSlowMultiplication(b, 1<<8)
}

func BenchmarkSlowMultiplicationSIMD_256(b *testing.B) {
	benchmarkSlowMultiplicationSIMD(b, 1<<8)
}

func BenchmarkSlowMultiplication_512(b *testing.B) {
	benchmarkSlowMultiplication(b, 1<<9)
}

func BenchmarkSlowMultiplicationSIMD_512(b *testing.B) {
	benchmarkSlowMultiplicationSIMD(b, 1<<9)
}

func BenchmarkSlowMultiplication_1024(b *testing.B) {
	benchmarkSlowMultiplication(b, 1<<10)
}

func BenchmarkSlowMultiplicationSIMD_1024(b *testing.B) {
	benchmarkSlowMultiplicationSIMD(b, 1<<10)
}

func BenchmarkSlowMultiplication_65536(b *testing.B) {
	benchmarkSlowMultiplication(b, 1<<16)
}

func BenchmarkSlowMultiplicationSIMD_65536(b *testing.B) {
	benchmarkSlowMultiplicationSIMD(b, 1<<16)
}

func benchmarkSlowMultiplication(b *testing.B, second uint64) {
	product := uint64(0)

	first := uint64(3)

	for range b.N {
		product = SlowMultiplication(first, second)
	}

	b.StopTimer()

	require.Equal(b, first*second, product)
}

func benchmarkSlowMultiplicationSIMD(b *testing.B, second uint64) {
	product := uint64(0)

	first := uint64(3)

	for range b.N {
		product = SlowMultiplicationSIMD(first, second)
	}

	b.StopTimer()

	require.Equal(b, first*second, product)
}
