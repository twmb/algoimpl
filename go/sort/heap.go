package sort

func left(i int) int   { return 2*i + 1 }     // 2 * (i + 1) - 1 because 0 indexed instead of 1
func right(i int) int  { return 2*i + 2 }     // 2 * (i + 1) - 1 + 1
func parent(i int) int { return (i+1)/2 - 1 } // (i + 1) / 2 - 1

type heap struct {
	collection *Sortable
	size       int
}

func heapify(heap heap, i int) {
	l := left(i)
	r := right(i)
	largestI := i
	if l < heap.size && (*heap.collection).Less(i, l) {
		largestI = l
	}
	if r < heap.size && (*heap.collection).Less(largestI, r) {
		largestI = r
	}
	if largestI != i {
		(*heap.collection).Swap(largestI, i)
		heapify(heap, largestI)
	}
}

// Creates a heap out of an unorganized Sortable collection.
// Runs in O(n) time.
func buildHeap(stuff Sortable) {
	heap := heap{&stuff, stuff.Len()}
	for i := stuff.Len()/2 - 1; i >= 0; i-- { // start at first non leaf (equiv. to parent of last leaf)
		heapify(heap, i)
	}
}

// Runs HeapSort on a Sortable collection.
// Runs in O(n * lg n) time, but amortizes worse than quicksort
func HeapSort(stuff Sortable) {
	buildHeap(stuff)
	heap := heap{&stuff, stuff.Len()}
	for i := stuff.Len() - 1; i > 0; i-- {
		stuff.Swap(0, i) // put max at end
		heap.size--
		heapify(heap, 0)
	}
}
