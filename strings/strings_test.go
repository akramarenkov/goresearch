package strings_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRange(t *testing.T) {
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
