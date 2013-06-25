#include <stdio.h>
#include <stdint.h>
#include <stdbool.h>

#include "binary_tree.h"

bool less(void *left_data, void *right_data) {
  if ((int64_t)left_data < (int64_t)right_data) {
    return true;
  }
  return false;
}

void print_node(void *data) {
  printf("%lu ", (int64_t)data);
}

int main() {
  struct binary_tree tree = new_binary_tree(); // tree of ints
  push_binary_tree(&tree, (void *)1, &less);
  push_binary_tree(&tree, (void *)0, &less);
  printf("root: %lu\n", (int64_t)tree.root->data);
  printf("lchild: %lu\n", (int64_t)tree.root->lchild->data);
  push_binary_tree(&tree, (void *)1, &less);
  printf("rchild: %lu\n", (int64_t)tree.root->rchild->data);
  push_binary_tree(&tree, (void *)3, &less);
  push_binary_tree(&tree, (void *)4, &less);
  push_binary_tree(&tree, (void *)2, &less);
  push_binary_tree(&tree, (void *)1, &less);
  push_binary_tree(&tree, (void *)2, &less);
  push_binary_tree(&tree, (void *)3, &less);
  push_binary_tree(&tree, (void *)9, &less);
  push_binary_tree(&tree, (void *)11, &less);
  push_binary_tree(&tree, (void *)8, &less);
  walk_node(tree.root, &print_node);
}

