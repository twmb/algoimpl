package various

import "testing"

func TestMaxSubarray(t *testing.T) {
	tests := []struct {
		In                      []int
		WantLI, WantRI, WantSum int
	}{
		{[]int{}, 0, 0, 0},
		{[]int{-1}, 0, 1, -1},
		{[]int{3, -1, -1, 4}, 0, 4, 5},        //whole thing
		{[]int{-1, 1, 1, -1}, 1, 3, 2},        //crossing
		{[]int{-1, -2, 1, 2}, 2, 4, 3},        //right side
		{[]int{1, 2, -3, -4}, 0, 2, 3},        //left side
		{[]int{1, -2, -3, 5, 6, 7}, 3, 6, 18}, // 6 length, right side
		{[]int{1, -2, -3, 5, 6}, 3, 5, 11},    // 5 length, right side
	}

	for _, test := range tests {
		li, ri, sum := MaxSubarray(test.In, 0, len(test.In))
		if li != test.WantLI || ri != test.WantRI || sum != test.WantSum {
			t.Errorf("test results (%v, %v, %v) != test expected (%v, %v, %v) for array %v", li, ri, sum, test.WantLI, test.WantRI, test.WantSum, test.In)
		}
	}
}
