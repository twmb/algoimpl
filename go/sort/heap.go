package sort

// Shuffles a smaller value at index i in a heap
// down to the appropriate spot. Complexity is O(lg n).
func ShuffleDown(heap Sortable, i, end int) {
	for {
		l := 2*i + 1
		if l >= end { // int overflow? (in go source)
			break
		}
		li := l
		if r := l + 1; r < end && heap.Less(l, r) {
			li = r // 2*i + 2
		}
		if heap.Less(li, i) {
			break
		}
		heap.Swap(li, i)
		i = li
	}
}

// Shuffles a larger value in a heap at index i
// up to the appropriate spot. Complexity is O(lg n).
func ShuffleUp(heap Sortable, i int) {
	for {
		pi := (i - 1) / 2 // (i + 1) / 2 - 1, parent
		if i == pi || !heap.Less(pi, i) {
			break
		}
		heap.Swap(pi, i)
		i = pi
	}
}

// Creates a max heap out of an unorganized Sortable collection.
// Runs in O(n) time.
func BuildHeap(stuff Sortable) {
	for i := stuff.Len()/2 - 1; i >= 0; i-- { // start at first non leaf (equiv. to parent of last leaf)
		ShuffleDown(stuff, i, stuff.Len())
	}
}

// Runs HeapSort on a Sortable collection.
// Runs in O(n * lg n) time, but amortizes worse than quicksort
func HeapSort(stuff Sortable) {
	BuildHeap(stuff)
	for i := stuff.Len() - 1; i > 0; i-- {
		stuff.Swap(0, i) // put max at end
		ShuffleDown(stuff, 0, i)
	}
}
