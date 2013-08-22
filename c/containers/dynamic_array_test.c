#include <stdio.h>
#include <stdint.h>

#include "dynamic_array.h"

int main(void) {
  dynarr d = create_dynarr();
  for (int64_t i = 0; i < 20; i++) {
    dynarr_append(d, (void *)i);
  }
  printf("printing\n");
  for (int64_t i = dynarr_len(d) - 1; i >= 0; i--) {
    printf("%lu\n", (int64_t)dynarr_at(d, i));
  }
  return 0;
}
