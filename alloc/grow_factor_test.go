package alloc

import (
	"testing"
	"unsafe"
)

// Difference between two independent slices.
type growDiff struct {
	// Iteration of append at which the slice address change was determined
	DetectedAt int
	// Slice lengths ratio
	Factor float64
	// Length of the independent slice that came before the some known
	From int
	// Length of some known independent slice
	To int
}

func researchGrowFactor(depth int) []growDiff {
	diffs := make([]growDiff, 0)

	slice := make([]byte, 0)

	previous := slice

	for range depth {
		slice = append(slice, 1)

		if unsafe.SliceData(slice) != unsafe.SliceData(previous) {
			diff := growDiff{
				DetectedAt: len(slice),
				From:       cap(previous),
				To:         cap(slice),
				Factor:     float64(cap(slice)) / float64(cap(previous)),
			}

			previous = slice

			diffs = append(diffs, diff)
		}
	}

	return diffs
}

func TestResearchGrowFactor(t *testing.T) {
	t.Logf("%+v", researchGrowFactor(1<<26))
}
