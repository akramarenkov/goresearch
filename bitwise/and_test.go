package bitwise

import (
	"testing"

	"github.com/akramarenkov/goresearch/internal/getenv"

	"github.com/stretchr/testify/require"
)

func BenchmarkAnd2(b *testing.B) {
	greaterZero := false
	lessZero := false

	first, err := getenv.Int("", 1)
	require.NoError(b, err)

	second, err := getenv.Int("", 2)
	require.NoError(b, err)

	firstN := -first
	secondN := -second

	for range b.N {
		greaterZero = first > 0 && second > 0
		lessZero = firstN < 0 && secondN < 0
	}

	require.True(b, greaterZero)
	require.True(b, lessZero)
}

func BenchmarkAnd2Bitwise(b *testing.B) {
	greaterZero := false
	lessZero := false

	first, err := getenv.Int("", 1)
	require.NoError(b, err)

	second, err := getenv.Int("", 2)
	require.NoError(b, err)

	firstN := -first
	secondN := -second

	for range b.N {
		greaterZero = first|second > 0
		lessZero = firstN&secondN < 0
	}

	require.True(b, greaterZero)
	require.True(b, lessZero)
}

func BenchmarkAnd3(b *testing.B) {
	greaterZero := false
	lessZero := false

	first, err := getenv.Int("", 1)
	require.NoError(b, err)

	second, err := getenv.Int("", 2)
	require.NoError(b, err)

	third, err := getenv.Int("", 3)
	require.NoError(b, err)

	firstN := -first
	secondN := -second
	thirdN := -third

	for range b.N {
		greaterZero = first > 0 && second > 0 && third > 0
		lessZero = firstN < 0 && secondN < 0 && thirdN < 0
	}

	require.True(b, greaterZero)
	require.True(b, lessZero)
}

func BenchmarkAnd3Bitwise(b *testing.B) {
	greaterZero := false
	lessZero := false

	first, err := getenv.Int("", 1)
	require.NoError(b, err)

	second, err := getenv.Int("", 2)
	require.NoError(b, err)

	third, err := getenv.Int("", 3)
	require.NoError(b, err)

	firstN := -first
	secondN := -second
	thirdN := -third

	for range b.N {
		greaterZero = first|second|third > 0
		lessZero = firstN&secondN&thirdN < 0
	}

	require.True(b, greaterZero)
	require.True(b, lessZero)
}

func BenchmarkAnd4(b *testing.B) {
	greaterZero := false
	lessZero := false

	first, err := getenv.Int("", 1)
	require.NoError(b, err)

	second, err := getenv.Int("", 2)
	require.NoError(b, err)

	third, err := getenv.Int("", 3)
	require.NoError(b, err)

	fourth, err := getenv.Int("", 4)
	require.NoError(b, err)

	firstN := -first
	secondN := -second
	thirdN := -third
	fourthN := -fourth

	for range b.N {
		greaterZero = first > 0 && second > 0 && third > 0 && fourth > 0
		lessZero = firstN < 0 && secondN < 0 && thirdN < 0 && fourthN < 0
	}

	require.True(b, greaterZero)
	require.True(b, lessZero)
}

func BenchmarkAnd4Bitwise(b *testing.B) {
	greaterZero := false
	lessZero := false

	first, err := getenv.Int("", 1)
	require.NoError(b, err)

	second, err := getenv.Int("", 2)
	require.NoError(b, err)

	third, err := getenv.Int("", 3)
	require.NoError(b, err)

	fourth, err := getenv.Int("", 4)
	require.NoError(b, err)

	firstN := -first
	secondN := -second
	thirdN := -third
	fourthN := -fourth

	for range b.N {
		greaterZero = first|second|third|fourth > 0
		lessZero = firstN&secondN&thirdN&fourthN < 0
	}

	require.True(b, greaterZero)
	require.True(b, lessZero)
}

func BenchmarkAnd5(b *testing.B) {
	greaterZero := false
	lessZero := false

	first, err := getenv.Int("", 1)
	require.NoError(b, err)

	second, err := getenv.Int("", 2)
	require.NoError(b, err)

	third, err := getenv.Int("", 3)
	require.NoError(b, err)

	fourth, err := getenv.Int("", 4)
	require.NoError(b, err)

	fifth, err := getenv.Int("", 5)
	require.NoError(b, err)

	firstN := -first
	secondN := -second
	thirdN := -third
	fourthN := -fourth
	fifthN := -fifth

	for range b.N {
		greaterZero = first > 0 && second > 0 && third > 0 && fourth > 0 && fifth > 0
		lessZero = firstN < 0 && secondN < 0 && thirdN < 0 && fourthN < 0 && fifthN < 0
	}

	require.True(b, greaterZero)
	require.True(b, lessZero)
}

func BenchmarkAnd5Bitwise(b *testing.B) {
	greaterZero := false
	lessZero := false

	first, err := getenv.Int("", 1)
	require.NoError(b, err)

	second, err := getenv.Int("", 2)
	require.NoError(b, err)

	third, err := getenv.Int("", 3)
	require.NoError(b, err)

	fourth, err := getenv.Int("", 4)
	require.NoError(b, err)

	fifth, err := getenv.Int("", 5)
	require.NoError(b, err)

	firstN := -first
	secondN := -second
	thirdN := -third
	fourthN := -fourth
	fifthN := -fifth

	for range b.N {
		greaterZero = first|second|third|fourth|fifth > 0
		lessZero = firstN&secondN&thirdN&fourthN&fifthN < 0
	}

	require.True(b, greaterZero)
	require.True(b, lessZero)
}
