package management

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryBuffer(t *testing.T) {
	b := NewMemoryBuffer()
	got := b.Get(200, 100) // from buffer position 200, get 100
	assert.Equal(t, 100, len(got))
}
