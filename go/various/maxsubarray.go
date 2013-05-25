// Implements the maximum subarray algorithm on a slice of ints
package various

func maxCrossingSubarray(array []int, from, to int) (li, ri, sum int) {
	mid := (from + to) / 2
	// left index, left side's sum, new running sum
	li, lsum, nsum := mid-1, array[mid-1], array[mid-1]
	for n := li - 1; n >= from; n-- {
		nsum += array[n]
		if nsum > lsum {
			lsum = nsum
			li = n
		}
	}
	ri, rsum, nsum := mid, array[mid], array[mid]
	for n := ri + 1; n < to; n++ {
		nsum += array[n]
		if nsum > rsum {
			rsum = nsum
			ri = n
		}
	}
	return li, ri + 1, lsum + rsum // one after last valid index
}

func MaxSubarray(array []int, from, to int) (li, ri, sum int) {
	if from >= to-1 {
		if to-from == 0 {
			return from, to, 0
		}
		return from, to, array[from]
	} else {
		lli, lri, lv := MaxSubarray(array, from, (from+to)/2)
		rli, rri, rv := MaxSubarray(array, (from+to)/2, to)
		cli, cri, cv := maxCrossingSubarray(array, from, to)
		if lv > rv && lv > cv {
			return lli, lri, lv // left's left index, right index, sum
		} else if rv > lv && rv > cv {
			return rli, rri, rv // right's left index, right index, sum
		} else {
			return cli, cri, cv // crossing left index, right index, sum
		}
	}
}
