package sort

func left(i int) int   { return 2*i + 1 }     // 2 * (i + 1) - 1 because 0 indexed instead of 1
func right(i int) int  { return 2*i + 2 }     // 2 * (i + 1) - 1 + 1
func parent(i int) int { return (i+1)/2 - 1 } // (i + 1) / 2 - 1

type heap struct {
	A    *Sortable
	size int
}

func heapify(heap heap, i int) {
	l := left(i)
	r := right(i)
	largestI := i
	if l < heap.size && (*heap.A).Less(i, l) {
		largestI = l
	}
	if r < heap.size && (*heap.A).Less(largestI, r) {
		largestI = r
	}
	if largestI != i {
		(*heap.A).Swap(largestI, i)
		heapify(heap, largestI)
	}
}

// Creates a heap from a partially organized Sortable collection.
// Runs in O(lg n) time, where n is the length of the collection.
// i is the index that may be less than the left and right sorted trees.
func Heapify(stuff Sortable, i int) {
	if i >= stuff.Len() || i < 0 {
		return
	}
	heap := heap{&stuff, stuff.Len()}
	heapify(heap, i)
}

// Creates a heap out of an unorganized Sortable collection.
// Runs in O(n) time.
func BuildHeap(stuff Sortable) {
	heap := heap{&stuff, stuff.Len()}
	for i := stuff.Len()/2 - 1; i >= 0; i-- { // start at first non leaf (equiv. to parent of last leaf)
		heapify(heap, i)
	}
}

// Runs HeapSort on a Sortable collection.
// Runs in O(n * lg n) time, but amortizes worse than quicksort
func HeapSort(stuff Sortable) {
	BuildHeap(stuff)
	heap := heap{&stuff, stuff.Len()}
	for i := stuff.Len() - 1; i > 0; i-- {
		stuff.Swap(0, i) // put max at end
		heap.size--
		heapify(heap, 0)
	}
}
