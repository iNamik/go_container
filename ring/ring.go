// Package ring implements a bounded First-In-First-Out queue, conforming to the iNamik Queue interface.
package ring

import "github.com/iNamik/go_container/queue"

// NewRing returns a new ring, which implements the inamik/container/queue/Queue interface
func New(capacity int) queue.Interface {
	return &ring{data: make([]interface{}, capacity), head: 0, tail: 0, length: 0}
}
