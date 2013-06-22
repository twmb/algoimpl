package queues

import (
	"github.com/twmb/algoimpl/go/sort"
)

// Any type that implements heap.Interface may be used as a
// max-heap. First call Init before using any of the heap functions.
// The heap functions will panic if called inappropriately (as in, you
// cannot call Peek on an empty heap).
//
// This interface embeds the Sortable interface, meaning elements
// can be compared and swapped. It extends the Sortable interface by
// allowing access to individual elements or by appending to the collection.
//
// Note that Push and Pop are for this package to call. To add or remove
// elements from a heap, use heap.Push or heap.Pop.
type Interface interface {
	sort.Sortable
	// Returns the value at the index.
	At(int) interface{}
	// Sets the value at the index to a new value.
	Set(i int, val interface{})
	// Adds a value to the end of the collection.
	Push(val interface{})
	// Removes the value at the end of the collection.
	Pop() interface{}
}

// Makes a heap out of the passed in collection that implements
// heap.Interface. Runs in O(n) time, where n = h.Len().
func Init(h Interface) {
	sort.BuildHeap(h)
}

// Removes and returns the maximum of the heap and reorganizes.
// The complexity is O(lg n).
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(n, 0)
	sort.ShuffleDown(h, 0, n)
	return h.Pop()
}

// This function will push a new value into a priority queue.
func Push(h Interface, val interface{}) {
	h.Push(val)
	sort.ShuffleUp(h, h.Len()-1)
}

// Returns the maximum of the heap.
func Peek(h Interface) interface{} {
	return h.At(0)
}

// This function will change the value at the index and will reorganize
// the priority queue as appropriate.
func Change(h Interface, i int, to interface{}) {
	h.Set(i, to)
	sort.ShuffleUp(h, i)
	sort.ShuffleDown(h, i, h.Len())
}

// Removes and erturns the element at index i
func Remove(h Interface, i int) (v interface{}) {
	n := h.Len() - 1
	if n != i {
		h.Swap(n, i)
		sort.ShuffleDown(h, i, n)
		sort.ShuffleUp(h, i)
	}
	return h.Pop()
}
