CC=gcc
CFLAGS=-Wall -pedantic -std=c11

all: insertion_sort_test

insertion_sort_test: insertion_sort_test.c sort.o
	$(CC) $(CFLAGS) -o $@ $< sort.o

sort.o: insertion_sort.c sort.h
	$(CC) -c $(CFLAGS) -o sort.o $<

clean:
	rm -f sort.o insertion_sort_test
