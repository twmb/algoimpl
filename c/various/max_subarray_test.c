#include <stdio.h>

#include "max_subarray.h"

typedef struct {
  int *In;
  int InStart, InEnd, WantLI, WantRI, WantSum;
} test_vals;

int main(int argc, char **argv) {
  int failed = 0; // false

  int a1[1] = {-1};
  int a2[4] = {3,-1,-1,4}; // whole thing
  int a3[4] = {-1,1,1,-1}; // crossing middle
  int a4[4] = {-1,-2,1,2}; // right side
  int a5[4] = {1,2,-3,-4}; // left side
  int a6[6] = {1,-2,-3,5,6,7}; // 6 length, right side
  int a7[5] = {1,-2,-3,5,6}; //5 length, right side

  int test_count = 10;
  test_vals tests[10] = { // test count
    {a1, 0, 1, 0, 1, -1},
    {a2, 0, 4, 0, 4, 5},
    {a3, 0, 4, 1, 3, 2},
    {a4, 0, 4, 2, 4, 3},
    {a5, 0, 4, 0, 2, 3},
    {a6, 0, 6, 3, 6, 18},
    {a7, 0, 5, 3, 5, 11},
    {a7, 0, 3, 0, 1, 1},
    {a7, 3, 5, 3, 5, 11},
    {a7, 1, 3, 1, 2, -2},
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
    printf("all good\n");
  } else {
    printf("failures\n");
  }
}

