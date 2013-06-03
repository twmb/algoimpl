#include <stdio.h>

#include "max_subarray.h"

typedef struct {
  int *In;
  int InStart, InEnd, WantLI, WantRI, WantSum;
} test_vals;

int main(int argc, char **argv) {
  int failed = 0; // false

  int a0[1] = {-1};
  int a1[4] = {3,-1,-1,4}; // whole thing
  int a2[4] = {-1,1,1,-1}; // crossing middle
  int a3[4] = {-1,-2,1,2}; // right side
  int a4[4] = {1,2,-3,-4}; // left side
  int a5[6] = {1,-2,-3,5,6,7}; // 6 length, right side
  int a6[5] = {1,-2,-3,5,6}; //5 length, right side

  int test_count = 10;
  test_vals tests[10] = { // test count
    {a0, 0, 1, 0, 1, -1}, // array, start, end, want li, want ri, want sum
    {a1, 0, 4, 0, 4, 5},
    {a2, 0, 4, 1, 3, 2},
    {a3, 0, 4, 2, 4, 3},
    {a4, 0, 4, 0, 2, 3},
    {a5, 0, 6, 3, 6, 18},
    {a6, 0, 5, 3, 5, 11},
    {a6, 0, 3, 0, 1, 1},
    {a6, 3, 5, 3, 5, 11},
    {a6, 1, 3, 1, 2, -2},
  };

  for (int i = 0; i < test_count; i++) {
    max_info info = max_subarray(tests[i].In, tests[i].InStart, tests[i].InEnd);
    if (info.l != tests[i].WantLI || 
        info.r != tests[i].WantRI ||
        info.sum != tests[i].WantSum) {
      printf("failure on %d, ret info: (%d, %d, %d), expected: (%d, %d, %d)\n",
          i, info.l, info.r, info.sum, 
          tests[i].WantLI, tests[i].WantRI, tests[i].WantSum);
      failed = 1; // true
    }
  }

  if (!failed) {
    printf("all iterative subarray tests good\n");
  }

  failed = 0;

  for (int i = 0; i < test_count; i++) {
    max_info info = max_subarray_recursive(tests[i].In, tests[i].InStart, tests[i].InEnd);
    if (info.l != tests[i].WantLI || 
        info.r != tests[i].WantRI ||
        info.sum != tests[i].WantSum) {
      printf("failure on %d, ret info: (%d, %d, %d), expected: (%d, %d, %d)\n",
          i, info.l, info.r, info.sum, 
          tests[i].WantLI, tests[i].WantRI, tests[i].WantSum);
      failed = 1; // true
    }
  }

  if (!failed) {
    printf("all recursive subarray tests good\n");
  } 
}

