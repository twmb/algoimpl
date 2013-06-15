package queues

import (
	"errors"
	"fmt"
	"github.com/twmb/algoimpl/go/sort"
)

func left(i int) int  { return 2*i + 1 } // 2 * (i + 1) - 1 because 0 indexed instead of 1
func right(i int) int { return 2*i + 2 } // 2 * (i + 1) - 1 + 1

// This interface embeds the Sortable interface, meaning elements
// can be compared and swapped. It extends the Sortable interface by
// allowing access to individual elements or by appending to the collection.
type Interface interface {
	sort.Sortable
	// Returns the value at the index.
	At(int) interface{}
	// Sets the value at the index to a new value.
	Set(i int, val interface{}) error
	// Adds a value to the end of the collection.
	Push(val interface{}) error
	// Removes the value at the end of the collection.
	Pop() (interface{}, error)
}

// shuffle down if value at i is smaller than children
// O(lg n), O(1) if it is not
func shuffleDown(heap Interface, i int) {
	l := left(i)
	r := right(i)
	largestI := i
	if l < heap.Len() && heap.Less(i, l) {
		largestI = l
	}
	if r < heap.Len() && heap.Less(largestI, r) {
		largestI = r
	}
	if largestI != i {
		heap.Swap(largestI, i)
		shuffleDown(heap, largestI)
	}
}

// shuffle up if the value at index i larger than the parent
// O(lg n), O(1) if it is not
func shuffleUp(h Interface, i int) {
	for {
		pi := (i - 1) / 2 // parent: (i + 1) / 2 - 1
		if i == pi || !h.Less(pi, i) {
			break
		}
		h.Swap(pi, i)
		i = pi
	}
}

// Returns a new priority queue. It can be modified using the heap
// functions below, all of which run in O(lg n) time.
func NewPriorityQueue(stuff Interface) {
	sort.BuildHeap(stuff)
}

// Returns the maximum of the heap.
func Maximum(h Interface) (interface{}, error) {
	if h.Len() > 0 {
		return h.At(0), nil
	}
	return nil, errors.New("Cannot call maximum on empty heap")
}

// This function will change the value at the index and will reorganize
// the priority queue as appropriate.
func Change(h Interface, i int, to interface{}) error {
	if i >= h.Len() {
		return fmt.Errorf("Cannot change index %v, out of bounds of collection (length %v)", i, h.Len())
	}
	err := h.Set(i, to)
	if err != nil {
		return err
	}
	shuffleUp(h, i)
	shuffleDown(h, i)
	return nil
}

// This function will push a new value into a priority queue should
// error if the value is not of the same type.
func Push(h Interface, val interface{}) error {
	err := h.Push(val)
	if err != nil {
		return err
	}
	shuffleUp(h, h.Len()-1)
	return nil
}

// Returns the maximum of the heap and reorganizes.
// The complexity is O(lg n)
func Pop(h Interface) (max interface{}, err error) {
	max, err = h.Pop()
	if err != nil {
		return
	}
	h.Swap(h.Len()-1, 0)
	max, err = h.Pop()
	shuffleDown(h, 0)
	return
}

// Removes and erturns the element at index i
func Remove(h Interface, i int) (v interface{}, err error) {
	if i > h.Len()-1 {
		return nil, errors.New("Index too large")
	}
	n := h.Len() - 1
	h.Swap(n, i)
	v, err = h.Pop()
	if n != i {
		shuffleDown(h, i)
		shuffleUp(h, i)
	}
	return
}
