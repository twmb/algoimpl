#include <stdio.h>
#include <stdbool.h>
#include <string.h>

#include "permute_letters.h"

struct letters {
  char letter;
  bool used;
};

void do_permute_letters(struct letters *all, int len, char *buffer, int position) {
  if (position == len) {
    printf("%s\n", buffer);
    return;
  }
  for (int i = 0; i < len; i++) {
    if (!all[i].used) {
      buffer[position] = all[i].letter;
      all[i].used = true;
      do_permute_letters(all, len, buffer, position+1);
      all[i].used = false;
    }
  }
}

void permute_letters(const char *letters, int len) {
  char buffer[len+1];
  buffer[len] = '\0';
  strcpy(buffer, letters);
  struct letters all[len];
  for (int i = 0; i < len; i++) {
    all[i].letter = letters[i];
    all[i].used = false;
  }
  do_permute_letters(all, len, buffer, 0);
}
