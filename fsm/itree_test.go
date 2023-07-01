package fsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestITreeFSMIsFreeEmpty(t *testing.T) {
	f := NewIntervalTreeFreeSpaceManager()

	offset, _ := f.Dirty(200)
	assert.Equal(t, offset, 0)
}

func TestITreeFSMTop(t *testing.T) {
	f := NewIntervalTreeFreeSpaceManager()

	f.Dirty(1000)
	f.Dirty(200)
	offset, _ := f.Dirty(300)

	assert.Equal(t, 1200, offset)
}

func TestITreeFSMBottom(t *testing.T) {
	f := NewIntervalTreeFreeSpaceManager()

	_, c := f.Dirty(1000)
	c()
	offset, _ := f.Dirty(200)
	assert.Equal(t, 0, offset)
}

func TestITreeFSMBetwen(t *testing.T) {
	f := NewIntervalTreeFreeSpaceManager()

	f.Dirty(150)
	_, c := f.Dirty(200)
	f.Dirty(1000)
	c()
	offset, _ := f.Dirty(100)
	assert.Equal(t, 150, offset)
}
