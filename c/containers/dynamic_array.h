#ifndef TWMB_DYNARRAY
#define TWMB_DYNARRAY

typedef struct dynamic_array *dynarr;

dynarr create_dynarr(void);
dynarr make_dynarr(int len, int cap);
void destroy_dynarr(dynarr array);
void *dynarr_append(dynarr array, void *element);
dynarr dynarr_slice(dynarr array, int from, int to);
void *dynarr_at(dynarr array, int position);
void dynarr_set(dynarr array, int position, void *element);
int dynarr_len(dynarr array);

#endif
