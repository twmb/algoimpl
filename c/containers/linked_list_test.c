#include <stdio.h>
#include <stdint.h>

#include "linked_list.h"

int main() {
  linked_list mine = new_linked_list();
  push_ll(&mine, (void *)3);
  push_ll(&mine, (void *)2);
  push_ll(&mine, (void *)2);
  item *head = mine.head;
  while (head != NULL) {
    printf("hello, %lu\n", (int64_t)head->data);
    head = head->next;
  }
  printf("\n");
  remove_element_ll(&mine, (void *)2);
  head = mine.head;
  while (head != NULL) {
    printf("hello, %lu\n", (int64_t)head->data);
    head = head->next;
  }
  delete_ll(&mine);
}
