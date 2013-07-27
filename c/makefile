ALL_TEST_SRC 	:= $(shell find . -name '*_test.c')
ALL_TEST_EXE	:= ${ALL_TEST_SRC:.c=}

CFLAGS		:= -Wall -pedantic -std=c11

all:	${ALL_TEST_EXE}

clean:
	rm -f ${ALL_TEST_EXE}

test:	all
	@for test in ${ALL_TEST_EXE};              \
		do                                       \
			echo testing $${test} ;                \
			./$${test} ;                           \
		  if [ $$? -ne 0 ] ;                     \
		  then                                   \
		    /bin/echo -e "\e[01;31mfailed\e[m" ; \
		  else                                   \
		    echo passed ;                        \
		  fi ;                                   \
		  echo ;                                 \
		done 

%_test:	%_test.c %.c %.h
	clang ${CFLAGS} -o $@ $(filter %_test.c %.c, $^)

