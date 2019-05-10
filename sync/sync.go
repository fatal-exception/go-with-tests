package sync

import "sync"

// Counter holds a monotonically incrementing value
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc increments count in struct
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns counter value
func (c *Counter) Value() (i int) {
	return c.value
}

// NewCounter constructs Counter objects
func NewCounter() *Counter {
	return &Counter{}
}
