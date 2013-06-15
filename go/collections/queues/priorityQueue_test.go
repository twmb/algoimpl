package queues

import (
	"errors"
	"testing"
)

type Ints []int

func (p *Ints) Len() int             { return len(*p) }
func (p *Ints) Less(i, j int) bool   { return (*p)[i] < (*p)[j] }
func (p *Ints) Swap(i, j int)        { (*p)[i], (*p)[j] = (*p)[j], (*p)[i] }
func (p *Ints) At(i int) interface{} { return (*p)[i] }
func (p *Ints) Set(i int, val interface{}) error {
	v, ok := val.(int)
	if ok {
		(*p)[i] = v
		return nil
	}
	return errors.New("Set() passed in type differs from what the collection can hold")
}
func (p *Ints) Push(val interface{}) error {
	v, ok := val.(int)
	if ok {
		*p = append(*p, v)
		return nil
	}
	return errors.New("Push: passed in type is not int")
}
func (p *Ints) Pop() (v interface{}, err error) {
	if p.Len() < 1 {
		err = errors.New("Cannot delete index larger than length of collection")
		return
	}
	*p, v = (*p)[:p.Len()-1], (*p)[p.Len()-1]
	return
}

func TestNewPriorityQueue(t *testing.T) {
	tests := []struct {
		In   Ints
		Want Interface
	}{
		{ // test empty
			Ints([]int{}),
			Interface((*Ints)(&[]int{})),
		},
		{
			Ints([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}),
			Interface((*Ints)(&[]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1})),
		},
	}
	for _, test := range tests {
		NewPriorityQueue(&test.In)
		ints := Ints(test.In)
		failed := false
		for i, v := range ints {
			if v != test.Want.At(i) {
				failed = true
				break
			}
		}
		if failed {
			t.Errorf("Failing Ints: result %v != want %v", test.In, test.Want)
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		In        Interface
		Want      int
		WantError string
	}{
		{
			Interface((*Ints)(&[]int{})),
			0,
			"Cannot call maximum on empty heap",
		},
		{
			Interface((*Ints)(&[]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1})),
			16,
			"",
		},
	}
	for _, test := range tests {
		got, gotErr := Maximum(test.In)
		if gotErr != nil {
			if gotErr.Error() != test.WantError {
				t.Errorf("Incorrect error received")
			}
			continue
		}
		gotInt, ok := got.(int)
		if !ok {
			t.Errorf("Could not convert returned value %v to int", gotInt)
			continue
		}
		if gotInt != test.Want {
			t.Errorf("Incorrect maximum, received %v, wanted %v", gotInt, test.Want)
		}
	}
}

func TestChange(t *testing.T) {
	tests := []struct {
		In           Interface
		InIToChange  int
		InIValChange interface{}
		WantInts     Ints
		WantError    string
	}{
		{
			Interface((*Ints)(&[]int{})),
			0,
			0,
			Ints([]int{}),
			"Cannot change index 0, out of bounds of collection (length 0)",
		},
		{
			Interface((*Ints)(&[]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1})),
			8,
			15,
			Ints([]int{16, 15, 10, 14, 7, 9, 3, 2, 8, 1}),
			"",
		},
		{
			Interface((*Ints)(&[]int{1})),
			0,
			"hello",
			Ints([]int{1}),
			"Set() passed in type differs from what the collection can hold",
		},
	}
	for _, test := range tests {
		gotErr := Change(test.In, test.InIToChange, test.InIValChange)
		if gotErr != nil {
			if gotErr.Error() != test.WantError {
				t.Errorf("Incorrect error: received %v, wanted %v", gotErr, test.WantError)
			}
			continue
		}
		changedInts := test.In.(*Ints)
		failed := false
		for i, v := range *changedInts {
			if v != test.WantInts[i] {
				failed = true
				break
			}
		}
		if failed {
			t.Errorf("Failing Ints: result %v != want %v", changedInts, test.WantInts)
		}
	}
}

func TestPush(t *testing.T) {
	tests := []struct {
		In        Interface
		PushVal   interface{}
		WantInts  Ints
		WantError string
	}{
		{
			Interface((*Ints)(&[]int{})),
			0,
			Ints([]int{0}),
			"",
		},
		{
			Interface((*Ints)(&[]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1})),
			15,
			Ints([]int{16, 15, 10, 8, 14, 9, 3, 2, 4, 1, 7}),
			"",
		},
		{
			Interface((*Ints)(&[]int{1})),
			"hello",
			Ints([]int{1}),
			"Push: passed in type is not int",
		},
	}
	for _, test := range tests {
		gotErr := Push(test.In, test.PushVal)
		if gotErr != nil {
			if gotErr.Error() != test.WantError {
				t.Errorf("Incorrect error: received %v, wanted %v", gotErr, test.WantError)
			}
			continue
		}
		changedInts := test.In.(*Ints)
		failed := false
		for i, v := range *changedInts {
			if v != test.WantInts[i] {
				failed = true
				break
			}
		}
		if failed {
			t.Errorf("Failing Ints: result %v != want %v", changedInts, test.WantInts)
		}
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		In        Interface
		PopIndex  int
		WantInts  Ints
		WantV     int
		WantError string
	}{
		{
			Interface((*Ints)(&[]int{})),
			0,
			Ints([]int{}),
			0,
			"Cannot delete index larger than length of collection",
		},
		{
			Interface((*Ints)(&[]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1})),
			1,
			Ints([]int{14, 8, 10, 4, 7, 9, 3, 2, 1}),
			//Ints([]int{16, 8, 10, 4, 7, 9, 3, 2, 1}),
			16,
			"",
		},
	}
	for _, test := range tests {
		got, gotErr := Pop(test.In)
		if gotErr != nil {
			if gotErr.Error() != test.WantError {
				t.Errorf("Incorrect error: received %v, wanted %v", gotErr, test.WantError)
			}
			continue
		}
		if test.WantV != got.(int) {
			t.Errorf("Return value %v != wanted %v", got, test.WantV)
		}
		changedInts := test.In.(*Ints)
		failed := false
		for i, v := range *changedInts {
			if v != test.WantInts[i] {
				failed = true
				break
			}
		}
		if failed {
			t.Errorf("Failing Ints: result %v != want %v", changedInts, test.WantInts)
		}
	}
}

// Taken from Go source code "heap_test.go" and modified to fit my structures
// I must learn to make tests this easy...
func (h Ints) verify(t *testing.T, i int) {
	n := h.Len()
	left := 2*i + 1
	right := 2*i + 2
	if left < n {
		if h.Less(i, left) {
			t.Errorf("heap invariant invalidated [%d] = %d > [%d] = %d", i, h[i], left, h[left])
			return
		}
		h.verify(t, left)
	}
	if right < n {
		if h.Less(i, right) {
			t.Errorf("heap invariant invalidated [%d] = %d > [%d] = %d", i, h[i], left, h[right])
			return
		}
		h.verify(t, right)
	}
}

func TestRemove1(t *testing.T) {
	h := new(Ints)
	for i := 0; i < 10; i++ {
		Push(h, i)
	}
	h.verify(t, 0)
	// removes the max every time
	for i := 0; h.Len() > 0; i++ {
		x, _ := Remove(h, 0)
		if x.(int) != 9-i {
			t.Errorf("Remove(0) got %d; want %d", x, i)
		}
		h.verify(t, 0)
	}
}

func TestRemove2(t *testing.T) {
	N := 10
	h := new(Ints)
	for i := 0; i < N; i++ {
		Push(h, i)
	}
	h.verify(t, 0)
	// tests that it removed all
	m := make(map[int]bool)
	for h.Len() > 0 {
		x, _ := Remove(h, (h.Len()-1)/2) // remove from middle
		m[x.(int)] = true
		h.verify(t, 0)
	}
	if len(m) != N {
		t.Errorf("len(m) = %d; want %d", len(m), N)
	}
	// and removed all correctly
	for i := 0; i < len(m); i++ {
		if !m[i] {
			t.Errorf("m[%d] doesn't exist", i)
		}
	}
}
