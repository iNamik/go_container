GOROOT ?= $(shell printf 't:;@echo $$(GOROOT)\n' | gomake -f -)
include $(GOROOT)/src/Make.inc

.PHONY: default all install clean nuke

default: all

all: install

install: 
	$(MAKE) -C queue install
	$(MAKE) -C ring  install

clean:
	$(MAKE) -C queue clean
	$(MAKE) -C ring  clean

nuke:
	$(MAKE) -C queue nuke
	$(MAKE) -C ring  nuke
