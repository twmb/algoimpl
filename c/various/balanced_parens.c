#include <stdbool.h>

#include "balanced_parens.h"

bool do_is_balanced(const char *string, int *strpos, char curparen) {
  for (; string[*strpos] != '\0'; (*strpos)++) {
    switch (string[*strpos]) {
      case '(':
      case '{':
      case '[':
        if (!do_is_balanced(string, strpos, string[(*strpos)++])) {
          return false;
        }
        break;
      case ')': if (curparen == '(') return true; return false; break;
      case ']': if (curparen == '[') return true; return false; break;
      case '}': if (curparen == '{') return true; return false; break;
    }
  }
  return curparen == '\0';
}

bool is_balanced(const char *string) {
  int zPos = 0;
  return do_is_balanced(string, &zPos, '\0');
}
