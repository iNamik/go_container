GOROOT ?= $(shell printf 't:;@echo $$(GOROOT)\n' | gomake -f -)
include $(GOROOT)/src/Make.inc

TARG=github.com/iNamik/go_container/queue

GOFILES=\
	impl.go\
	queue.go\

include $(GOROOT)/src/Make.pkg
