#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

#include "heap.h"

typedef struct dynint {
  int64_t *ints;
  int len;
  int cap;
} dynint;

dynint create_dynint(void) {
  dynint r;
  r.len = 0;
  r.cap = 0;
  return r;
}

void *append(dynint *array, int64_t newint) {
  if (array->cap == 0) {
    array->ints = malloc(sizeof(int) * 1);
    array->cap = 1;
  }
  if (array->len == array->cap) {
    if (array->cap > 1000) {
      array->ints = realloc(array->ints, (int)(sizeof(int64_t) * 1.2 * array->cap));
      array->cap = (int)(array->cap * 1.2);
    } else {
      array->ints = realloc(array->ints, sizeof(int64_t) * 2 * array->cap);
      array->cap *= 2;
    }
  }
  if (array->ints != NULL) {
    array->ints[array->len] = newint;
  }
  array->len++;
  return array->ints;
}

dynint slice(dynint *array, int from, int to) {
  if (to > array->len) {
    exit(1);
  }
  dynint new;
  new.ints = &array->ints[from];
  new.len = to-from;
  new.cap = array->cap-from;
  return new;
}

void destroy_dynint(dynint array) {
  free(array.ints);
}

bool less(void *container, int left, int right) {
  dynint *d = container;
  return d->ints[left] > d->ints[right];
}
int len(void *container) {
  return ((dynint *)container)->len;
}
void swap(void *container, int left, int right) {
  if (left != right) {
    dynint *d = container;
    d->ints[left] ^= d->ints[right];
    d->ints[right] ^= d->ints[left];
    d->ints[left] ^= d->ints[right];
  }
}
void ints_push(void *container, void *elem) {
  append((dynint *)container, (int64_t)elem);
}
void *ints_pop(void *container) {
  dynint *d = container;
  int64_t end = d->ints[d->len-1];
  *d = slice(d, 0, d->len - 1);
  return (void *)end;
}



int main(void) {
  Heap myheap = create_heap();

  dynint array = create_dynint();

  set_heap_container(myheap, &array);
  set_heap_lessfunc(myheap, &less);
  set_heap_lenfunc(myheap, &len);
  set_heap_swapfunc(myheap, &swap);
  set_heap_pushfunc(myheap, &ints_push);
  set_heap_popfunc(myheap, &ints_pop);
  for (int64_t i = 0; i < 20; i++) {
    //heap_push(myheap, (void *)i);
    append(&array, i);
  }
  heapify(myheap);
  for (int i = 0; i < array.len; i++) {
    printf("%lu\n", array.ints[i]);
  }
  while (array.len > 0) {
    int64_t popped = (int64_t)heap_pop(myheap);
    printf("popped: %lu\n", popped);
  }

  destroy_heap(myheap);
  destroy_dynint(array);
  return 0;
}

