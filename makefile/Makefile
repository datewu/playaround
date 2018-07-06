##### makefile for compile C programs
CC = gcc # Compiler to use
OPTIONS = -O2 -g -Wall # -g for debug, -O2 for optimise and -Wall additonal messages
INCLUDES = -I . # Directory for header file
OBJS = main.o module.o # List of objects to be build
.PHONY: all list clean

all: ${OBJS}
	@echo "Building..." # To print "Building..." message
	${CC} ${OPTIONS} ${INCLUDES} ${OBJS} -o main_bin

%.o: %.c  # % pattern wildcard matching
	${CC} ${OPTIONS} ${INCLUDES} -c %.c

list:
	@echo $(shell ls) # To print output of command `ls`

clean:
	@echo Cleaning up...
	-rm -rf *.0 # - prefix for ignoring errors and continue execution
	-rm main_bin


#### makefile for img manage
FILES = $(shell find imgs -type f -iname "*.jpg" | sed 's/imgs/thumb/g')
CONVERT_CMD = convert -resize "100x100" $< $@
MSG = "\nUpdating thumbnail" $@

all_thumb: ${FILES}

thumb/%.jpg: imgs/%.jpg
	${MSG}
	${CONVERT_CMD}

thumb/%.JPG: imgs/%.JPG
	${MSG}
	${CONVERT_CMD}

clean_all:
	@echo Cleaning up files...
	-rm -rf thumb/*.{jpg,JPG}
