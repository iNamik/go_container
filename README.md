go_container
============

**Container Classes in Go**


PACKAGES:
--------

### import "github.com/iNamik/go_container/queue"

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


### import "github.com/iNamik/go_container/ring"

The 'ring' package implements a bounded FIFO using the 'queue' interface


### import "github.com/iNamik/go_container/stack"

The 'stack' package implements an unbounded FILO using the 'queue' interface


INSTALL:
--------

The packages are built using the Go tool.  Assuming you have correctly set the
$GOPATH variable, you can run the folloing command(s):

	go get github.com/iNamik/go_container/queue
	go get github.com/iNamik/go_container/ring
	go get github.com/iNamik/go_container/stack


AUTHORS:
--------

 * David Farrell

