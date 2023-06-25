package fsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingFSMIsFreeEmpty(t *testing.T) {
	f := NewCountingFreeSpaceManager()

	offset, _ := f.Dirty(200)
	assert.Equal(t, offset, 0)
}

func TestCountingFSMTop(t *testing.T) {
	f := NewCountingFreeSpaceManager()

	f.Dirty(1000)
	offset, _ := f.Dirty(200)

	assert.Equal(t, 1000, offset)
}

func TestCountingFSMBottom(t *testing.T) {
	f := NewCountingFreeSpaceManager()

	_, c := f.Dirty(1000)
	c()
	offset, _ := f.Dirty(200)
	assert.Equal(t, 0, offset)
}

func TestCountingFSMBetwen(t *testing.T) {
	f := NewCountingFreeSpaceManager()

	f.Dirty(150)
	_, c := f.Dirty(200)
	f.Dirty(1000)
	c()
	offset, _ := f.Dirty(100)
	assert.Equal(t, 150, offset)
}
