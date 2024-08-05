package alloc

import (
	"testing"
	"unsafe"
)

// Difference between two independent slices
type diff struct {
	// Iteration of append at which the slice address change was determined
	DetectedAt int
	// Slice lengths ratio
	Factor float64
	// Length of the independent slice that came before the last known
	From int
	// Length of the last known independent slice
	To int
}

func researchGrowFactor(depth int) []diff {
	diffs := make([]diff, 0)

	slice := make([]bool, 1)

	previousAt := 1
	previousPtr := unsafe.SliceData(slice)

	for range depth {
		slice = append(slice, true)

		currentPtr := unsafe.SliceData(slice)

		if currentPtr != previousPtr {
			detectedAt := len(slice)

			from := previousAt - 1
			to := detectedAt - 1

			previousAt = detectedAt
			previousPtr = currentPtr

			diff := diff{
				DetectedAt: detectedAt,
				From:       from,
				To:         to,
				Factor:     float64(to) / float64(from),
			}

			diffs = append(diffs, diff)
		}
	}

	return diffs
}

func TestResearchGrowFactor(t *testing.T) {
	t.Logf("%+v", researchGrowFactor(1<<26))
}
