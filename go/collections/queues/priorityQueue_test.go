package queues

import (
	"errors"
	"testing"
)

type Ints []int

func (p Ints) Len() int             { return len(p) }
func (p Ints) Less(i, j int) bool   { return p[i] < p[j] }
func (p Ints) Swap(i, j int)        { p[i], p[j] = p[j], p[i] }
func (p Ints) At(i int) interface{} { return p[i] }
func (p Ints) Set(i int, val interface{}) error {
	v, ok := val.(int)
	if ok {
		p[i] = v
		return nil
	}
	return errors.New("Set() passed in type differs from what the collection can hold")
}
func (p Ints) Append(val interface{}) (ModSortable, error) {
	v, ok := val.(int)
	if ok {
		p = append(p, v)
		return &p, nil
	}
	return &p, errors.New("Append: passed in type is not int")
}
func (p Ints) Delete(index int) (ModSortable, error) {
	if index >= len(p) {
		return &p, errors.New("Can't delete, index too large")
	}
	copy(p[index:], p[index+1:])
	p = p[0 : len(p)-1]
	return &p, nil
}

func TestNewPriorityQueue(t *testing.T) {
	wantInts0 := Ints([]int{})
	wantInts1 := Ints([]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1})
	tests := []struct {
		In   Ints
		Want *ModifiableHeap
	}{
		{ // test empty
			Ints([]int{}),
			&ModifiableHeap{collection: &wantInts0, size: 0},
		},
		{
			Ints([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}),
			&ModifiableHeap{collection: &wantInts1, size: len(wantInts1)},
		},
	}
	for _, test := range tests {
		got := NewPriorityQueue(test.In)
		ints, ok := got.collection.(Ints)
		if !ok {
			t.Errorf("yo %v not ints, wtf yo", ints)
			continue
		}
		failed := false
		for i, v := range ints {
			if v != test.Want.collection.At(i) {
				failed = true
				break
			}
		}
		if failed {
			t.Errorf("Failing Ints: result %v != want %v", test.In, test.Want)
		}
	}
}

// called with a heap
// if not, dumb, but result same
func TestMaximum(t *testing.T) {
	inputInts0 := Ints([]int{})
	inputInts1 := Ints([]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1})
	tests := []struct {
		Caller    *ModifiableHeap
		Want      int
		WantError string
	}{
		{
			&ModifiableHeap{collection: &inputInts0, size: len(inputInts0)},
			0,
			"Cannot call maximum on empty heap",
		},
		{
			&ModifiableHeap{collection: &inputInts1, size: len(inputInts1)},
			16,
			"",
		},
	}
	for _, test := range tests {
		got, gotErr := test.Caller.Maximum()
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

func TestChangeValue(t *testing.T) {
	inputInts0 := Ints([]int{})
	wantInts0 := Ints([]int{})
	inputInts1 := Ints([]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1})
	wantInts1 := Ints([]int{16, 15, 10, 14, 7, 9, 3, 2, 8, 1})
	inputInts2 := Ints([]int{1})
	wantInts2 := Ints([]int{1})
	tests := []struct {
		Caller       *ModifiableHeap
		InIToChange  int
		InIValChange interface{}
		WantInts     Ints
		WantError    string
	}{
		{
			&ModifiableHeap{collection: &inputInts0, size: len(inputInts0)},
			0,
			0,
			wantInts0,
			"Cannot change index 0, out of bounds of collection (length 0)",
		},
		{
			&ModifiableHeap{collection: &inputInts1, size: len(inputInts1)},
			8,
			15,
			wantInts1,
			"",
		},
		{
			&ModifiableHeap{collection: &inputInts2, size: len(inputInts2)},
			0,
			"hello",
			wantInts2,
			"Set() passed in type differs from what the collection can hold",
		},
	}
	for _, test := range tests {
		gotErr := test.Caller.ChangeValue(test.InIToChange, test.InIValChange)
		if gotErr != nil {
			if gotErr.Error() != test.WantError {
				t.Errorf("Incorrect error: received %v, wanted %v", gotErr, test.WantError)
			}
			continue
		}
		changedInts := test.Caller.collection.(*Ints)
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

func TestInsert(t *testing.T) {
	inputInts0 := Ints([]int{})
	wantInts0 := Ints([]int{0})
	inputInts1 := Ints([]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1})
	wantInts1 := Ints([]int{16, 15, 10, 8, 14, 9, 3, 2, 4, 1, 7})
	inputInts2 := Ints([]int{1})
	wantInts2 := Ints([]int{1})
	tests := []struct {
		Caller      *ModifiableHeap
		InValChange interface{}
		WantInts    Ints
		WantError   string
	}{
		{
			&ModifiableHeap{collection: &inputInts0, size: len(inputInts0)},
			0,
			wantInts0,
			"",
		},
		{
			&ModifiableHeap{collection: &inputInts1, size: len(inputInts1)},
			15,
			wantInts1,
			"",
		},
		{
			&ModifiableHeap{collection: &inputInts2, size: len(inputInts2)},
			"hello",
			wantInts2,
			"Append: passed in type is not int",
		},
	}
	for _, test := range tests {
		gotErr := test.Caller.Insert(test.InValChange)
		if gotErr != nil {
			if gotErr.Error() != test.WantError {
				t.Errorf("Incorrect error: received %v, wanted %v", gotErr, test.WantError)
			}
			continue
		}
		changedInts := test.Caller.collection.(*Ints)
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

func TestDelete(t *testing.T) {
	inputInts0 := Ints([]int{})
	wantInts0 := Ints([]int{})
	inputInts1 := Ints([]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1})
	wantInts1 := Ints([]int{16, 8, 10, 4, 7, 9, 3, 2, 1})
	tests := []struct {
		Caller      *ModifiableHeap
		DeleteIndex int
		WantInts    Ints
		WantError   string
	}{
		{
			&ModifiableHeap{collection: &inputInts0, size: len(inputInts0)},
			0,
			wantInts0,
			"Cannot delete index larger than length of collection",
		},
		{
			&ModifiableHeap{collection: &inputInts1, size: len(inputInts1)},
			1,
			wantInts1,
			"",
		},
	}
	for _, test := range tests {
		gotErr := test.Caller.Delete(test.DeleteIndex)
		if gotErr != nil {
			if gotErr.Error() != test.WantError {
				t.Errorf("Incorrect error: received %v, wanted %v", gotErr, test.WantError)
			}
			continue
		}
		changedInts := test.Caller.collection.(*Ints)
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
