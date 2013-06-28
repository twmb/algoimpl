#include <stdio.h>
#include <string.h>

#include "reverse_words.h"

int main(int argc, char **argv) {
//  char line[80];
//  fgets(line, sizeof(line), stdin);
  char line[] = "hello this is easy\n";
  printf("%s\n", line);
  reverse_words(line);
  printf("%s\n", line);
  if (strcmp(line, "easy is this hello\n")) {
    printf("error, line %s not equal to 'easy is this hello\n'", line);
  }
  line[0] = '\0';
  reverse_words(line);
  if (strcmp(line, "")) {
    printf("error, line %s not equal to ''", line);
  }
  line[0] = '\n';
  line[1] = '\0';
  reverse_words(line);
  if (strcmp(line, "\n")) {
    printf("error, line %s not equal to '\n'", line);
  }
}

