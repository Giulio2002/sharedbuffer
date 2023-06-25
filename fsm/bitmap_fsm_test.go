package fsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitmapFSMIsFreeEmpty(t *testing.T) {
	f := NewBitmapFreeSpaceManager()
	offset, _ := f.Dirty(200)
	assert.Equal(t, offset, 0)
}

func TestBitmapFSMTop(t *testing.T) {
	f := NewBitmapFreeSpaceManager()

	f.Dirty(defaultWordSize * 11)
	offset, _ := f.Dirty(1)
	assert.Equal(t, defaultWordSize*11, offset)
}

func TestBitmapFSMBottom(t *testing.T) {
	f := NewBitmapFreeSpaceManager()

	_, c := f.Dirty(defaultWordSize * 11)
	f.Dirty(defaultWordSize * 11)
	c()
	offset, _ := f.Dirty(1)

	assert.Equal(t, 0, offset)
}

func TestBitmapFSMBetwen(t *testing.T) {
	f := NewBitmapFreeSpaceManager()

	f.Dirty(defaultWordSize)
	_, c := f.Dirty(defaultWordSize)
	f.Dirty(defaultWordSize)
	c()
	offset, _ := f.Dirty(defaultWordSize)
	assert.Equal(t, defaultWordSize, offset)
}
