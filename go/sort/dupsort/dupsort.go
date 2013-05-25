// Package dupsort, again, is my own implementation of different sort
// functions. The difference between dupsort and sort is that this package
// contains all algorithms that must actually access an element. This new
// requirement means that dupsort needs a larger interface (by one function).
//
// Generally, an algorithm needs to access an element if it needs to duplicate
// elements. Otherwise they could just be compared and swapped. This means
// that this package includes all algorithms that need to copy what is being
// sorted.
package dupsort

type DupSortable interface {
	// Len is the number of elements in the collection
	Len() int
	// Less returns whether the element i should
	// sort before the element j
	Less(i, j interface{}) bool
	// At accesses the element at index i
	At(i int) interface{}
	// Set sets the value at index i
	Set(i int, val interface{})
	// New returns a new DupSortable of length i
	New(i int) DupSortable
}

func mergeCombine(l DupSortable, r DupSortable) DupSortable {
	combined := l.New(r.Len() + l.Len())
	li, ri, ci := 0, 0, 0
	for ; li < l.Len() && ri < r.Len(); ci++ {
		if l.Less(l.At(li), r.At(ri)) {
			combined.Set(ci, l.At(li))
			li++
		} else {
			combined.Set(ci, r.At(ri))
			ri++
		}
	}
	for ; li < l.Len(); ci, li = ci+1, li+1 {
		combined.Set(ci, l.At(li))
	}
	for ; ri < r.Len(); ci, ri = ci+1, ri+1 {
		combined.Set(ci, r.At(ri))
	}
	return combined
}

func MergeSort(me DupSortable, from, to int) DupSortable {
	if from < to-1 {
		left := MergeSort(me, from, (from+to)/2)
		right := MergeSort(me, (from+to)/2, to)
		combined := mergeCombine(left, right)
		return combined
	}
	ele := me.New(to - from)
	if to-from > 0 {
		ele.Set(0, me.At(from))
	}
	return ele
}
