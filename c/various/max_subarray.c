#include "max_subarray.h"

// maximum subarray from left to right
max_info max_subarray(int *arr, int l, int r) {
  max_info max_now = {l, l, 0};
  if (r - l <= 1) {
    max_now.r = r;
    if (r == l) {
      return max_now;
    }
    max_now.sum = arr[l];
    return max_now;
  }
  max_now.sum = arr[l];
  max_info max_so_far = {l, l, arr[l]}; 
  for (l += 1; l < r; l++) {
    if (max_now.sum + arr[l] > arr[l]) { // yet net higher than start
      max_now.r = l;
      max_now.sum += arr[l];
    } else { // new lowest low
      max_now.l = l;
      max_now.r = l;
      max_now.sum = arr[l];
    }
    if (max_now.sum > max_so_far.sum) {
      max_so_far = max_now;
    }
  }
  max_so_far.r++; // increment to one past the actual right index
  return max_so_far;
}

