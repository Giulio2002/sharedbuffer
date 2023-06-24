package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitmap(t *testing.T) {
	b := NewBitmap()
	b.Set(2, true)
	assert.True(t, true, b.Get(2))
	b.Set(2, false)
	assert.False(t, false, b.Get(2))
}

func TestBitmapUnsetted(t *testing.T) {
	b := NewBitmap()
	assert.False(t, false, b.Get(2))
}
