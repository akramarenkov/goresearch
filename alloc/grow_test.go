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

func researchGrowFactor(depth int) ([]growDiff, []int) {
	diffs := make([]growDiff, 0)
	caps := make([]int, 0)

	slice := make([]bool, 1)
	caps = append(caps, cap(slice))

	previous := slice
	previousDetectedAt := 1

	for range depth {
		slice = append(slice, true) //nolint:makezero

		if unsafe.SliceData(slice) != unsafe.SliceData(previous) {
			// Difference data is filled in as if nothing is known about the
			// current slice
			detectedAt := len(slice)

			from := previousDetectedAt - 1
			to := detectedAt - 1

			previous = slice
			previousDetectedAt = detectedAt

			diff := growDiff{
				DetectedAt: detectedAt,
				From:       from,
				To:         to,
				Factor:     float64(to) / float64(from),
			}

			diffs = append(diffs, diff)

			// However, the capacity of the current slice is known
			caps = append(caps, cap(slice))
		}
	}

	return diffs, caps
}

func TestResearchGrowFactor(t *testing.T) {
	diffs, caps := researchGrowFactor(1 << 26)

	t.Logf("%+v\n%+v", diffs, caps)
}
