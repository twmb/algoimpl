package various

func inversionsCombine(left, right, combined []int) int {
	inversions := 0
	k, li, ri := 0, 0, 0 // index in combined array
	for ; li < len(left) && ri < len(right); k++ {
		if left[li] < right[ri] {
			combined[k] = left[li]
			li++
		} else { // right less than left
			combined[k] = right[ri]
			inversions += len(left) - li // if a right element is larger than a left,
			ri++                         // then it is larger than every element remaining on the left
		}
	}
	for ; li < len(left); li, k = li+1, k+1 {
		combined[k] = left[li]
	}
	for ; ri < len(right); ri, k = ri+1, k+1 {
		combined[k] = right[ri]
	}
	return inversions
}

// performs a mergesort while counting inversions
func inversionsCount(array, sorted []int) int {
	if len(array) <= 1 {
		return 0
	}
	cleft := inversionsCount(array[:len(array)/2], sorted[:len(array)/2])
	cright := inversionsCount(array[len(array)/2:], sorted[len(array)/2:])
	ccross := inversionsCombine(array[:len(array)/2], array[len(array)/2:], sorted)
	return cleft + ccross + cright
}

// Inversions will return the number of inversions in a given input integer array.
// An inversion is when a smaller number appears after a larger number.
// For example, [1,3,5,2,4,6] has three inversions: [3,2], [5,2] and [5,4].
func Inversions(array []int) int {
	sorted := make([]int, len(array))
	count := inversionsCount(array, sorted)
	return count
}
