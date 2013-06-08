package matrix

import (
  "errors"
)

// I do not check if any row length differs from others: don't do that.
// You know who does that? Crazy people.

func BasicMultiply(A, B [][]int) ([][]int, error) {
  if len(A) == 0 || len(B) == 0 {
    return nil, errors.New("Cannot multiply empty matrices")
  }
  if len(A[0]) == 0 || len(B[0]) == 0 {
    return nil, errors.New("Cannot multiply empty matrices")
  }
  if len(A[0]) != len(B) {
    return nil, errors.New("Dimension mismatch")
  }
  C := make([][]int, len(A))
  for r := range A {
    C[r] = make([]int, len(B[0]))
    for c := range B[0] {
      for k := range A[r] {
        C[r][c] += A[r][k] * B[k][c]
      }
    }
  }
  return C, nil
}

