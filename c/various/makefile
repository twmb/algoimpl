TESTS=max_subarray reverse_words
CC=gcc
CFLAGS=-Wall -pedantic -std=c11

all: $(TESTS)

$(TESTS): %: %_test.c %.o
	$(CC) $(CFLAGS) -o $@ $^

%.o: %.c %.h
	$(CC) -c $(CFLAGS) -o $@ $<

clean:
	rm -f *.o $(TESTS)
