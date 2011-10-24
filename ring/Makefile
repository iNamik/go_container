GOROOT ?= $(shell printf 't:;@echo $$(GOROOT)\n' | gomake -f -)
include $(GOROOT)/src/Make.inc

TARG=github.com/iNamik/container.go/ring

GCIMPORTS+="-I ../"

GOFILES=\
	impl.go\
	ring.go\

GC=${O}g -I../queue/_obj
LD=${O}l -L../queue/_obj

include $(GOROOT)/src/Make.pkg
