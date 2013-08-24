CFLAGS		:= -Wall -pedantic -std=c11

special: heap_test_special

heap_test_special: heap_test_special.c heap.c dynamic_array.c heap.h dynamic_array.h
	clang ${CFLAGS} -o $@ heap_test_special.c dynamic_array.c

clean:
	rm -f heap_test_special
