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

func TestSlowMultiplicationF64(t *testing.T) {
	require.True(t, cpuid.CPU.Supports(cpuid.AVX))

	for first := range uint64(1 << 10) {
		for second := range uint64(1 << 10) {
			require.InDelta(
				t,
				float64(first*second),
				SlowMultiplicationF64(float64(first), float64(second)),
				0,
			)
			require.InDelta(
				t,
				float64(first*second),
				SlowMultiplicationF64SIMD(float64(first), float64(second)),
				0,
			)
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

func BenchmarkSlowMultiplicationF64_64(b *testing.B) {
	benchmarkSlowMultiplicationF64(b, 1<<6)
}

func BenchmarkSlowMultiplicationF64SIMD_64(b *testing.B) {
	benchmarkSlowMultiplicationF64SIMD(b, 1<<6)
}

func BenchmarkSlowMultiplicationF64_128(b *testing.B) {
	benchmarkSlowMultiplicationF64(b, 1<<7)
}

func BenchmarkSlowMultiplicationF64SIMD_128(b *testing.B) {
	benchmarkSlowMultiplicationF64SIMD(b, 1<<7)
}

func BenchmarkSlowMultiplicationF64_256(b *testing.B) {
	benchmarkSlowMultiplicationF64(b, 1<<8)
}

func BenchmarkSlowMultiplicationF64SIMD_256(b *testing.B) {
	benchmarkSlowMultiplicationF64SIMD(b, 1<<8)
}

func BenchmarkSlowMultiplicationF64_512(b *testing.B) {
	benchmarkSlowMultiplicationF64(b, 1<<9)
}

func BenchmarkSlowMultiplicationF64SIMD_512(b *testing.B) {
	benchmarkSlowMultiplicationF64SIMD(b, 1<<9)
}

func BenchmarkSlowMultiplicationF64_1024(b *testing.B) {
	benchmarkSlowMultiplicationF64(b, 1<<10)
}

func BenchmarkSlowMultiplicationF64SIMD_1024(b *testing.B) {
	benchmarkSlowMultiplicationF64SIMD(b, 1<<10)
}

func BenchmarkSlowMultiplicationF64_65536(b *testing.B) {
	benchmarkSlowMultiplicationF64(b, 1<<16)
}

func BenchmarkSlowMultiplicationF64SIMD_65536(b *testing.B) {
	benchmarkSlowMultiplicationF64SIMD(b, 1<<16)
}

func benchmarkSlowMultiplicationF64(b *testing.B, second uint64) {
	product := float64(0)

	first := float64(3)

	for range b.N {
		product = SlowMultiplicationF64(first, float64(second))
	}

	b.StopTimer()

	require.InDelta(b, first*float64(second), product, 0)
}

func benchmarkSlowMultiplicationF64SIMD(b *testing.B, second uint64) {
	product := float64(0)

	first := float64(3)

	for range b.N {
		product = SlowMultiplicationF64SIMD(first, float64(second))
	}

	b.StopTimer()

	require.InDelta(b, first*float64(second), product, 0)
}
