package insertsort

import "testing"

func TestSort(t *testing.T) {
	// 100% line coverage
	tests := []struct {
		In, Want []int
	}{
		{[]int{21, -10, 54, 0, 1098309}, []int{-10, 0, 21, 54, 1098309}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		Sort(test.In)
		failed := false
		for i, v := range test.In {
			if v != test.Want[i] {
				t.Errorf("Sort, position %v, sortval %v, supposed to be %v", i, v, test.Want[i])
				failed = true
			}
		}
		if failed {
			t.Errorf("Failing slices: %v != %v", test.In, test.Want)
		}
	}
}
