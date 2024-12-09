package iter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	benchmarkSamples = []string{
		"0123456789",
		"1011121314",
		"1516171819",
		"2021222324",
		"2526272829",
		"3031323334",
		"3536373839",
		"4041424344",
		"4546474849",
		"5051525354",
	}
)

func TestIter(t *testing.T) {
	input := "Hello, 世界"

	expectedIndices := []int{0, 1, 2, 3, 4, 5, 6, 7, 10}
	expectedLengths := []int{1, 1, 1, 1, 1, 1, 1, 3, 3}

	actualIndices := make([]int, 0, len(expectedIndices))
	actualLengths := make([]int, 0, len(expectedLengths))

	for id, symbol := range input {
		actualIndices = append(actualIndices, id)
		actualLengths = append(actualLengths, len(string(symbol)))
	}

	require.Equal(t, expectedIndices, actualIndices)
	require.Equal(t, expectedLengths, actualLengths)
}

func BenchmarkIterByte1(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:1] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range len(input) {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterRune1(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:1] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range input {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterByte2(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:2] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range len(input) {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterRune2(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:2] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range input {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterByte3(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:3] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range len(input) {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterRune3(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:3] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range input {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterByte4(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:4] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range len(input) {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterRune4(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:4] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range input {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterByte5(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:5] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range len(input) {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterRune5(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:5] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range input {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterByte6(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:6] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range len(input) {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterRune6(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:6] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range input {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterByte7(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:7] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range len(input) {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterRune7(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:7] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range input {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterByte8(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:8] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range len(input) {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterRune8(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:8] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range input {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterByte9(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:9] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range len(input) {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterRune9(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:9] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range input {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterByte10(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:10] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range len(input) {
			last = id
		}
	}

	require.NotNil(b, last)
}

func BenchmarkIterRune10(b *testing.B) {
	input := ""

	for _, sample := range benchmarkSamples[:10] {
		input += sample
	}

	b.ResetTimer()

	var last int

	for range b.N {
		for id := range input {
			last = id
		}
	}

	require.NotNil(b, last)
}
