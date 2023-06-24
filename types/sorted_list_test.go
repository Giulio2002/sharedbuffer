package types

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedList(t *testing.T) {
	lessFunc := func(a, b float64) bool {
		return a < b
	}
	// Add and check
	s := NewSortedList(lessFunc)
	s.Add(2.0)
	s.Add(23.0)
	s.Add(1.0)
	s.Add(0.0)
	s.Add(-2.0)
	cpyList := s.UnderlyingCopy()
	assert.Equal(t, len(cpyList), 5)
	assert.True(t, sort.SliceIsSorted(cpyList, func(i, j int) bool {
		return cpyList[i] < cpyList[j]
	}))
	// Try right shift.
	s.Set(0, 92)
	assert.Equal(t, 92.0, s.Get(4))
	// try leftshit
	s.Set(4, 5.0)
	assert.Equal(t, 5.0, s.Get(3))
	cpyList = s.UnderlyingCopy()
	fmt.Println(cpyList)
	// now do the search with success
	res, idx, found := s.Search(func(a float64) bool {
		return a >= 1 // we search for 1.
	})
	assert.True(t, found)
	assert.Equal(t, idx, 1)
	assert.Equal(t, res, 1.0)
}
