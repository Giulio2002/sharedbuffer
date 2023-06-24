package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopEmptyStack(t *testing.T) {
	c := NewConcurrentStack[int]()
	_, found := c.Pop()
	assert.False(t, found)
	assert.Equal(t, c.Size(), 0)
}

func TestAddAndPop(t *testing.T) {
	c := NewConcurrentStack[int]()
	c.Add(110)
	c.Add(111)
	c.Add(1)
	c.Add(523)
	c.Add(544)
	assert.Equal(t, c.Size(), 5)
	obj, found := c.Pop()
	assert.Equal(t, obj, 544)
	assert.True(t, found)
	c.Pop()
	c.Pop()
	obj, found = c.Pop()
	assert.Equal(t, obj, 111)
	assert.True(t, found)
	assert.Equal(t, c.Size(), 1)
	c.Pop()
	assert.Equal(t, c.Size(), 0)
}
