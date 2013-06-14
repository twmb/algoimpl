package sort

// uses Ints from sort_test.go

import "testing"

func TestHeapify(t *testing.T) {
	tests := []struct {
		In, Want   Ints
		misplacedI int
	}{
		{ // test empty
			Ints([]int{}),
			Ints([]int{}),
			0,
		},
		{ // from CLRS, 4 is out of place
			Ints([]int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}),
			Ints([]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}),
			1,
		},
	}

	for _, test := range tests {
		Heapify(test.In, test.misplacedI)
		failed := false
		for i, v := range test.In {
			if v != test.Want[i] {
				failed = true
				break
			}
		}
		if failed {
			t.Errorf("Failing Ints: result %v != want %v", test.In, test.Want)
		}
	}
}

func TestBuildHeap(t *testing.T) {
	tests := []struct {
		In, Want Ints
	}{
		{Ints([]int{}), Ints([]int{})},
		{
			Ints([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}),
			Ints([]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}),
		},
	}
	for _, test := range tests {
		BuildHeap(test.In)
		failed := false
		for i, v := range test.In {
			if v != test.Want[i] {
				failed = true
				break
			}
		}
		if failed {
			t.Errorf("Failing Ints: result %v != want %v", test.In, test.Want)
		}
	}
}
