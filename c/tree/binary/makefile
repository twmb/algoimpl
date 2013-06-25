TESTS=binary_tree
CC=gcc
CFLAGS=-Wall -g -pedantic -std=c11 -o3

all: $(TESTS)

$(TESTS): %: %_test.c %.o
	$(CC) $(CFLAGS) -o $@ $^

%.o: %.c %.h
	$(CC) -c $(CFLAGS) -o $@ $<

clean:
	rm -f *.o $(TESTS)
