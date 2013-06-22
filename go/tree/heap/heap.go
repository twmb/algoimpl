package queues

import (
	"github.com/twmb/algoimpl/go/tree"
)

// Makes a heap out of the passed in collection that implements
// tree.Interface. Runs in O(n) time, where n = h.Len().
func Init(h tree.Interface) {
	buildHeap(h)
}

// Removes and returns the maximum of the heap and reorganizes.
// The complexity is O(lg n).
func Pop(h tree.Interface) interface{} {
	n := h.Len() - 1
	h.Swap(n, 0)
	shuffleDown(h, 0, n)
	return h.Pop()
}

// This function will push a new value into a priority queue.
func Push(h tree.Interface, val interface{}) {
	h.Push(val)
	shuffleUp(h, h.Len()-1)
}

// Removes and erturns the element at index i
func Remove(h tree.Interface, i int) (v interface{}) {
	n := h.Len() - 1
	if n != i {
		h.Swap(n, i)
		shuffleDown(h, i, n)
		shuffleUp(h, i)
	}
	return h.Pop()
}

// Creates a max heap out of an unorganized tree.Interface collection.
// Runs in O(n) time.
func buildHeap(stuff tree.Interface) {
	for i := stuff.Len()/2 - 1; i >= 0; i-- { // start at first non leaf (equiv. to parent of last leaf)
		shuffleDown(stuff, i, stuff.Len())
	}
}

// Shuffles a smaller value at index i in a heap
// down to the appropriate spot. Complexity is O(lg n).
func shuffleDown(heap tree.Interface, i, end int) {
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
func shuffleUp(heap tree.Interface, i int) {
	for {
		pi := (i - 1) / 2 // (i + 1) / 2 - 1, parent
		if i == pi || !heap.Less(pi, i) {
			break
		}
		heap.Swap(pi, i)
		i = pi
	}
}
