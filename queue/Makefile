GOROOT ?= $(shell printf 't:;@echo $$(GOROOT)\n' | gomake -f -)
include $(GOROOT)/src/Make.inc

TARG=github.com/iNamik/container.go/queue

GOFILES=\
	impl.go\
	queue.go\

include $(GOROOT)/src/Make.pkg
