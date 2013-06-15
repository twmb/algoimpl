package queues

import (
	"errors"
	"fmt"
	"github.com/twmb/algoimpl/go/sort"
)

func left(i int) int   { return 2*i + 1 }     // 2 * (i + 1) - 1 because 0 indexed instead of 1
func right(i int) int  { return 2*i + 2 }     // 2 * (i + 1) - 1 + 1
func parent(i int) int { return (i+1)/2 - 1 } // (i + 1) / 2 - 1

// This interface embeds the Sortable interface, meaning elements
// can be compared and swapped. It extends the Sortable interface by
// allowing access to individual elements or by appending to the collection.
type ModSortable interface {
	sort.Sortable
	At(int) interface{}
	Set(i int, val interface{}) error
	Append(val interface{}) (ModSortable, error)
	Delete(index int) (ModSortable, error)
}

// A heap that can be appended to or can have individual values accessed
// or changed.
type ModifiableHeap struct {
	collection ModSortable
	size       int
}

func heapify(heap ModifiableHeap, i int) {
	l := left(i)
	r := right(i)
	largestI := i
	if l < heap.size && heap.collection.Less(i, l) {
		largestI = l
	}
	if r < heap.size && heap.collection.Less(largestI, r) {
		largestI = r
	}
	if largestI != i {
		heap.collection.Swap(largestI, i)
		heapify(heap, largestI)
	}
}

// Returns a new priority queue. It can be modified using the heap
// functions below, all of which run in O(lg n) time.
func NewPriorityQueue(stuff ModSortable) *ModifiableHeap {
	sort.BuildHeap(stuff)
	return &ModifiableHeap{stuff, stuff.Len()}
}

// Returns the maximum of the heap.
func (h *ModifiableHeap) Maximum() (interface{}, error) {
	if h.collection.Len() > 0 {
		return h.collection.At(0), nil
	}
	return nil, errors.New("Cannot call maximum on empty heap")
}

// Returns the maximum of the heap, reorganizes and decrements the size
// by one.
func (h *ModifiableHeap) ExtractMax() (interface{}, error) {
	if h.size < 0 {
		return 0, errors.New("heap underflow")
	}
	max := h.collection.At(0)
	h.collection.Swap(0, h.size-1)
	h.size--
	heapify(*h, 0)
	return max, nil
}

// This function will change the value at the index and will reorganize
// the priority queue as appropriate.
func (h *ModifiableHeap) ChangeValue(i int, to interface{}) error {
	if i >= h.collection.Len() {
		return fmt.Errorf("Cannot change index %v, out of bounds of collection (length %v)", i, h.collection.Len())
	}
	err := h.collection.Set(i, to)
	if err != nil {
		return err
	}
	pi := parent(i)
	// shuffle up if the changed value is now larger than the parents
	// O(lg n), O(1) if it is not
	for i > 0 && h.collection.Less(pi, i) {
		h.collection.Swap(pi, i)
		i = pi
		pi = parent(i)
	}
	// shuffle down if changed value smaller than children
	// O(lg n), O(1) if it is not
	heapify(*h, i)
	return nil
}

// The caller must make sure that the Append function returns the same
// type that is originally used in NewPriorityQueue.
// This function will insert a new value into a priority queue and will
// error if the value is not of the same type. If the return value of
// Append is not the same as what was originally used in NewPriorityQueue,
// this function will panic.
func (h *ModifiableHeap) Insert(val interface{}) error {
	returned, err := h.collection.Append(val)
	if err != nil {
		return err
	}
	h.collection = returned
	h.size++
	err = h.ChangeValue(h.size-1, val) // yes, changes twice, cum at me bra
	if err != nil {
		return err
	}
	return nil
}

func (h *ModifiableHeap) Delete(i int) error {
	if i >= h.collection.Len() {
		return errors.New("Cannot delete index larger than length of collection")
	}
	h.collection.Swap(h.collection.Len()-1, i)
	returned, err := h.collection.Delete(h.collection.Len() - 1)
	if err != nil { // should never happen
		return err
	}
	h.collection = returned
	heapify(*h, i)
	return nil
}
