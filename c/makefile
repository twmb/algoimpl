ALL_TEST_SRC 	:= $(shell find . -name '*_test.c')
ALL_TEST_EXE	:= ${ALL_TEST_SRC:.c=}

CFLAGS		:= -Wall -pedantic -std=c99

all:	${ALL_TEST_EXE}

clean:
	rm -f ${ALL_TEST_EXE}

test:	all
	@for test in ${ALL_TEST_EXE};  \
		do                           \
			echo testing $${test} ;    \
			./$${test} ;               \
		  if [ $$? -ne 0 ] ;         \
		  then                       \
		    echo failed ;            \
		  else                       \
		    echo passed ;            \
		  fi ;                       \
		  echo ;                     \
		done 

%_test:	%_test.c %.c %.h
	clang ${CFLAGS} -o $@ $(filter %_test.c %.c, $^)

