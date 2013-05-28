TESTS=insertion_sort_test merge_sort_test
CC=gcc
CFLAGS=-Wall -pedantic -std=c11

all: $(TESTS)

# Note: I don't know why I can't take out the _test, but hey, it works.
%_test: %_test.c %.o
	$(CC) $(CFLAGS) -o $@ $^

%.o: %.c %.h
	$(CC) -c $(CFLAGS) -o $@ $<

clean:
	rm -f *.o $(TESTS)
