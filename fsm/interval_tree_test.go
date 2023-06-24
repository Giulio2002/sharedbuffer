package fsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestITreeFSMIsFreeEmpty(t *testing.T) {
	f := NewIntervalTreeFreeSpaceManager()

	assert.Equal(t, f.FirstFreeIndex(200), 1)
}

func TestITreeFSMTop(t *testing.T) {
	f := NewIntervalTreeFreeSpaceManager()

	f.MarkBusy(0, 1000)
	assert.Equal(t, 1001, f.FirstFreeIndex(200))
}

func TestITreeFSMBottom(t *testing.T) {
	f := NewIntervalTreeFreeSpaceManager()

	f.MarkBusy(300, 1000)
	assert.Equal(t, 1, f.FirstFreeIndex(200))
}

func TestITreeFSMBetwen(t *testing.T) {
	f := NewIntervalTreeFreeSpaceManager()

	f.MarkBusy(0, 150)
	f.MarkBusy(300, 1000)
	assert.Equal(t, 151, f.FirstFreeIndex(10))
}

func TestITreeFSMSimpleWithFree(t *testing.T) {
	f := NewIntervalTreeFreeSpaceManager()

	f.MarkBusy(0, 1000)
	f.MarkFree(0, 1000)
	assert.Equal(t, 1, f.FirstFreeIndex(200))
	f.MarkBusy(0, 234)
	f.MarkBusy(1000, 500)
	f.MarkBusy(500, 200)
	f.MarkFree(500, 200)
	assert.Equal(t, 235, f.FirstFreeIndex(10))

}
