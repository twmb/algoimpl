#ifndef TWMB_BIN_TREE
#define TWMB_BIN_TREE

#include <stdbool.h>
#include <stdio.h>

struct node {
  struct node *lchild;
  struct node *rchild;
  struct node *parent;
  void *data;
};

struct binary_tree {
  struct node *root;
};

struct binary_tree new_binary_tree();
int push_binary_tree(struct binary_tree *tree, void *element,
    bool (*less)(void *left, void *right));
void walk_node(struct node *node, void (*print_node)(void *data));

#endif
