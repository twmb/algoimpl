// Implements merge sort on []ints.
// Knowing that the underlying type is a slice allows for using channels.
// This also allows for using goroutines.
package dupsort

func mergeCombineInts(lch, rch <-chan int, tch chan<- int) {
	lv, lopen := <-lch
	rv, ropen := <-rch
	for lopen && ropen {
		if lv < rv {
			tch <- lv
			lv, lopen = <-lch
		} else {
			tch <- rv
			rv, ropen = <-rch
		}
	}
	for lopen {
		tch <- lv
		lv, lopen = <-lch
	}
	for ropen {
		tch <- rv
		rv, ropen = <-rch
	}
	close(tch)
}

func MergeSortInts(me []int, from, to int, tch chan<- int) {
	if from < to-1 {
		lch, rch := make(chan int), make(chan int)
		go MergeSortInts(me, from, (from+to)/2, lch)
		go MergeSortInts(me, (from+to)/2, to, rch)
		mergeCombineInts(lch, rch, tch)
	} else {
		for i := from; i < to; i++ {
			tch <- me[i]
		}
		close(tch)
	}
}
