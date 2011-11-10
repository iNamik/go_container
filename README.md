container.go
============

**Container Classes in Go**


PACKAGES:
--------

### import "github.com/iNamik/container.go/queue"

The 'queue' package implements an unbounded FIFO with the following interface:

	// queue.Interface
	type Interface interface {
		Add       (interface{})
		Remove    ()    interface{}
		Peek      (int) interface{}
		Clear     ()
		Len       () int
		Empty     () bool
		Cap       () int
		AtCapactiy() bool
	}


### import "github.com/iNamik/container.go/ring"

The 'ring' package implements a bounded FIFO using the 'queue' interface


INSTALL:
--------

  To install the packages set manually

	git clone https://github.com/iNamik/container.go
	cd container.go
	gomake
	gomake install

  Or you can install each package via goinstall

	goinstall github.com/iNamik/container.go/queue
	goinstall github.com/iNamik/container.go/ring


AUTHORS:
--------

 * David Farrell

