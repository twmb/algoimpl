package insertsort

import "testing"

func TestSort(t *testing.T) {
	// 100% line coverage
	scrambled := []int{21, -10, 54, 0, 1098309}
	sorted := []int{-10, 0, 21, 54, 1098309}
	Sort(scrambled)
	testSort(t, scrambled, sorted)

	scrambled = []int{}
	sorted = []int{}
	Sort(scrambled)
	testSort(t, scrambled, sorted)
}

func testSort(t *testing.T, scrambled, sorted []int) {
	failed := false
	for i, v := range scrambled {
		if v != sorted[i] {
			t.Errorf("Sort, position %v, sortval %v, supposed to be %v", i, v, sorted[i])
			failed = true
		}
	}
	if failed {
		t.Errorf("Failing slices: %v != %v", scrambled, sorted)
	}
}
