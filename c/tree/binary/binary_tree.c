#include <stdlib.h>

#include "binary_tree.h"

struct binary_tree new_binary_tree() {
  struct binary_tree new;
  new.root = NULL;
  return new;
}

int push_binary_tree(struct binary_tree *tree, void *data,
    bool (*less)(void *left_data, void *right_data)) {
  struct node *currentParent = tree->root;
  struct node *current = tree->root;
  while (current != NULL) {
    if (less(data, current->data)) {
      currentParent = current;
      current = current->lchild;
    } else {
      currentParent = current;
      current = current->rchild;
    }
  }
  struct node *new_node = (struct node *)malloc(sizeof(struct node));
  if (!new_node) {
    return -1;
  }
  new_node->parent = currentParent;
  new_node->data = data;
  new_node->lchild = NULL;
  new_node->rchild = NULL;
  if (currentParent == NULL) { // no root element
    tree->root = new_node;
    return 0;
  }
  if (less(new_node->data, currentParent->data)) {
    currentParent->lchild = new_node;
  } else {
    currentParent->rchild = new_node;
  }
  return 0;
}

void walk_node(struct node *node, void (*print_node)(void *data)) {
  if (node != NULL) {
    walk_node(node->lchild, print_node);
    print_node(node->data);
    walk_node(node->rchild, print_node);
  }
}

