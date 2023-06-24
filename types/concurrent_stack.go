package types

import "sync"

type ConcurrentStack[T any] struct {
	l []T

	mu sync.Mutex
}

func NewConcurrentStack[T any]() *ConcurrentStack[T] {
	return &ConcurrentStack[T]{}
}

func (c *ConcurrentStack[T]) Add(elem T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.l = append(c.l, elem)
}

func (c *ConcurrentStack[T]) Size() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.l)
}

func (c *ConcurrentStack[T]) Pop() (obj T, found bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.l) == 0 {
		return
	}
	obj = c.l[len(c.l)-1]
	c.l = c.l[:len(c.l)-1]
	found = true
	return
}
