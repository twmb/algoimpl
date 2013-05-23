package insertsort

import "fmt"
/* 
 * Performs insertion sort on a slice of ints
 */
func sort(nums []int) {
  for j := 1; j < len(nums); j++ { // from the second spot to the last
    current := nums[j] // save the current number you are inserting
    var i int;
    for i = j-1; i >= 0 && nums[i] > current; i-- { // while the next left number is larger
      nums[i+1] = nums[i]                           // slide that number right one position
    } // i will end at -1
    nums[i+1] = current; // sent the leftmost position == current
  } // sorted
}
