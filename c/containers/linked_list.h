#ifndef GITHUB_TWMB_C_LL
#define GITHUB_TWMB_C_LL

// Push, remove
typedef struct item {
  struct item *next;
  void *data;
} item;

typedef struct {
  item *head;
} linked_list;

linked_list new_linked_list();
void push_front_ll(linked_list *list, void *data);
void *remove_front_ll(linked_list *list);
void *remove_element_ll(linked_list *list, void *data);

#endif
