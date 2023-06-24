package fsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitmapFSMIsFreeEmpty(t *testing.T) {
	f := NewBitmapFreeSpaceManager()

	assert.Equal(t, f.FirstFreeIndex(200), 0)
}

func TestBitmapFSMTop(t *testing.T) {
	f := NewBitmapFreeSpaceManager()

	f.MarkBusy(0, defaultWordSize*11)
	assert.Equal(t, defaultWordSize*11, f.FirstFreeIndex(200))
}

func TestBitmapFSMBottom(t *testing.T) {
	f := NewBitmapFreeSpaceManager()

	f.MarkBusy(5000, 1000)
	assert.Equal(t, 0, f.FirstFreeIndex(200))
}

func TestBitmapFSMBetwen(t *testing.T) {
	f := NewBitmapFreeSpaceManager()

	f.MarkBusy(0, 12)
	f.MarkBusy(defaultWordSize*2, 1000)
	assert.Equal(t, defaultWordSize, f.FirstFreeIndex(10))
}

func TestBitmapFSMSimpleWithFree(t *testing.T) {
	f := NewBitmapFreeSpaceManager()

	f.MarkBusy(0, 1000)
	f.MarkFree(0, 1000)
	assert.Equal(t, 0, f.FirstFreeIndex(200))
	f.MarkBusy(0, 234)
	f.MarkBusy(defaultWordSize, 500)
	f.MarkBusy(defaultWordSize*2, 200)
	f.MarkFree(defaultWordSize, 200)
	assert.Equal(t, defaultWordSize, f.FirstFreeIndex(10))

}
