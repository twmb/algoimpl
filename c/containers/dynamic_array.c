#include <stdlib.h> 

#include "dynamic_array.h"

struct dynamic_array {
  void **elements;
  int len;
  int cap;
};

dynarr create_dynarr(void) {
  dynarr r = malloc(sizeof(dynarr));
  r->len = 0;
  r->cap = 0;
  return r;
}

dynarr make_dynarr(int len, int cap) {
  dynarr r = malloc(sizeof(dynarr));
  r->len = len;
  r->cap = cap;
  r->elements = malloc(cap * sizeof(void*));
  return r;
}

void destroy_dynarr(dynarr array) {
  free(array->elements);
  free(array);
}

void *dynarr_append(dynarr array, void *element) {
  if (array->cap == 0) {
    array->elements = malloc(sizeof(void*));
    array->cap = 1;
  }
  if (array->len == array->cap) {
    if (array->cap > 1000) {
      array->elements = realloc(array->elements, (int)(sizeof(void*) * 1.2 * array->cap));
      array->cap = (int)(array->cap * 1.2);
    } else {
      array->elements = realloc(array->elements, sizeof(void*) * 2 * array->cap);
      array->cap *= 2;
    }
  }
  if (array->elements != NULL) {
    array->elements[array->len] = element;
  }
  array->len++;
  return array->elements;
}

dynarr dynarr_slice(dynarr array, int from, int to) {
  if (to > array->len) {
    exit(1);
  }
  dynarr new = create_dynarr();
  new->elements = &array->elements[from];
  new->len = to-from;
  new->cap = array->cap-from;
  return new;
}

void *dynarr_at(dynarr array, int i) {
  return array->elements[i];
}

int dynarr_len(dynarr array) {
  return array->len;
}
