package fsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingFSMIsFreeEmpty(t *testing.T) {
	f := NewCountingFreeSpaceManager()

	assert.Equal(t, f.FirstFreeIndex(200), 0)
}

func TestCountingFSMTop(t *testing.T) {
	f := NewCountingFreeSpaceManager()

	f.MarkBusy(0, 1000)
	assert.Equal(t, 1000, f.FirstFreeIndex(200))
}

func TestCountingFSMBottom(t *testing.T) {
	f := NewCountingFreeSpaceManager()

	f.MarkBusy(300, 1000)
	assert.Equal(t, 0, f.FirstFreeIndex(200))
}

func TestCountingFSMBetwen(t *testing.T) {
	f := NewCountingFreeSpaceManager()

	f.MarkBusy(0, 150)
	f.MarkBusy(300, 1000)
	assert.Equal(t, 150, f.FirstFreeIndex(10))
}

func TestCountingFSMSimpleWithFree(t *testing.T) {
	f := NewCountingFreeSpaceManager()

	f.MarkBusy(0, 1000)
	f.MarkFree(0, 1000)
	assert.Equal(t, 0, f.FirstFreeIndex(200))
	f.MarkBusy(0, 234)
	f.MarkBusy(1000, 500)
	f.MarkBusy(500, 200)
	f.MarkFree(500, 200)
	assert.Equal(t, 234, f.FirstFreeIndex(10))

}
