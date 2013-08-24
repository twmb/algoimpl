#include <stdlib.h> 
//#include <pthread.h>

#include "dynamic_array.h"

struct dynamic_array {
  void **elements;
  int len;
  int cap;

//  int *slicecount;
//  
//  pthread_mutexattr_t attr;
//  pthread_mutex_t mutex;
};

dynarr create_dynarr(void) {
  dynarr r = malloc(sizeof(dynarr));
  r->len = 0;
  r->cap = 0;
  
//  pthread_mutexattr_init(&r->attr);
//  pthread_mutex_init(&r->mutex, &r->attr);
  return r;
}

dynarr make_dynarr(int len, int cap) {
  dynarr r = malloc(sizeof(dynarr));
  if (len < 0) {
    exit(-1);
  }
  if (cap < len) {
    cap = len;
  }
  r->len = len;
  r->cap = cap;
  if (cap > 0) {
    r->elements = malloc(cap * sizeof(void*));
  }
  return r;
}

void destroy_dynarr(dynarr array) {
  // pthread_mutex_lock(&array->mutex);
  // if 
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

void *dynarr_at(dynarr array, int position) {
  return array->elements[position];
}

void dynarr_set(dynarr array, int position, void *element) {
  array->elements[position] = element;
}

int dynarr_len(dynarr array) {
  return array->len;
}
