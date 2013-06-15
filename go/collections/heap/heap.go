package heap

import (
	"github.com/twmb/algoimpl/go/sort"
)

func left(i int) int   { return 2*i + 1 }     // 2 * (i + 1) - 1 because 0 indexed instead of 1
func right(i int) int  { return 2*i + 2 }     // 2 * (i + 1) - 1 + 1
func parent(i int) int { return (i+1)/2 - 1 } // (i + 1) / 2 - 1

type Heap struct {
	Collection *sort.Sortable
	Size       int
}

func Heapify(heap Heap, i int) {
	l := left(i)
	r := right(i)
	largestI := i
	if l < heap.Size && (*heap.Collection).Less(i, l) {
		largestI = l
	}
	if r < heap.Size && (*heap.Collection).Less(largestI, r) {
		largestI = r
	}
	if largestI != i {
		(*heap.Collection).Swap(largestI, i)
		Heapify(heap, largestI)
	}
}

// Creates a heap out of an unorganized Sortable collection.
// Runs in O(n) time.
func BuildHeap(stuff sort.Sortable) {
	heap := Heap{&stuff, stuff.Len()}
	for i := stuff.Len()/2 - 1; i >= 0; i-- { // start at first non leaf (equiv. to parent of last leaf)
		Heapify(heap, i)
	}
}

// Runs HeapSort on a Sortable collection.
// Runs in O(n * lg n) time, but amortizes worse than quicksort
func HeapSort(stuff sort.Sortable) {
	BuildHeap(stuff)
	heap := Heap{&stuff, stuff.Len()}
	for i := stuff.Len() - 1; i > 0; i-- {
		stuff.Swap(0, i) // put max at end
		heap.Size--
		Heapify(heap, 0)
	}
}
