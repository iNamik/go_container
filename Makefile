GOROOT ?= $(shell printf 't:;@echo $$(GOROOT)\n' | gomake -f -)
include $(GOROOT)/src/Make.inc

.PHONY: all install clean queue ring

all: queue ring

queue:
	gomake -C queue

ring:
	gomake -C ring

install: queue ring
	gomake -C queue install
	gomake -C ring  install

clean:
	gomake -C queue clean
	gomake -C ring  clean

