// Package stack implements a First-In-First-Out stack, conforming to the iNamik Queue interface.
package stack

import "github.com/iNamik/go_container/queue"

// NewStack returns a new ring, which implements the inamik/container/queue/Queue interface
func New(capacity int) queue.Interface {
	return &stack{data: make([]interface{}, capacity), length: 0}
}
