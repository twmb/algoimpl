#include <stdbool.h>

#include "balanced_parens.h"

bool do_is_balanced(const char *string, int *strpos, char curparen) {
  for (; string[*strpos] != '\0'; (*strpos)++) {
    if (string[*strpos] == '(' || string[*strpos] == '{' || string[*strpos] == '[') {
      if (!do_is_balanced(string, strpos, string[(*strpos)++])) {
        return false;
      }
    } else {
      switch (string[*strpos]) {
        case ')': if (curparen == '(') return true; return false; break;
        case ']': if (curparen == '[') return true; return false; break;
        case '}': if (curparen == '{') return true; return false; break;
      }
    }
  }
  if (curparen != '\0') {
    return false;
  }
  return true;
}

bool is_balanced(const char *string) {
  int zPos = 0;
  return do_is_balanced(string, &zPos, '\0');
}
