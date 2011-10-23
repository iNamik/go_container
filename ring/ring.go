/*
	Package ring implements a bounded queue, conforming to the iNamik Queue interface
*/
package ring

import "github.com/iNamik/container.go/queue"

// NewRing returns a new ring, which implements the inamik/container/queue/Queue interface
func NewRing(capacity int) queue.Queue {
	return &ring{data: make([]interface{}, capacity), head: 0, tail: 0, length: 0}
}

