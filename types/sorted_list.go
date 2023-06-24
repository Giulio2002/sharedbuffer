package types

import (
	"sort"
)

type lessFunc[T any] func(a T, b T) bool
type cond[T any] func(a T) bool

type SortedList[T any] struct {
	u    []T
	less lessFunc[T]
}

func NewSortedList[T any](less lessFunc[T]) *SortedList[T] {
	return &SortedList[T]{
		less: less,
	}
}

func (s *SortedList[T]) Add(elem T) {
	s.u = append(s.u, elem)
	for i := len(s.u) - 2; i >= 0; i-- {
		if !s.less(s.u[i+1], s.u[i]) {
			break
		}
		s.u[i+1], s.u[i] = s.u[i], s.u[i+1]
	}
}

func (s *SortedList[T]) Get(idx int) T {
	return s.u[idx]
}

func (s *SortedList[T]) Set(idx int, elem T) {
	s.u[idx] = elem
	leftshift := idx > 0 && s.less(s.u[idx], s.u[idx-1])
	// shift to the left :(
	if leftshift {
		for i := idx - 1; i >= 0; i-- {
			if !s.less(s.u[i+1], s.u[i]) {
				break
			}
			s.u[i+1], s.u[i] = s.u[i], s.u[i+1]
		}
		return
	}
	// shift to the right :)
	for i := idx; i < len(s.u)-1; i++ {
		if s.less(s.u[i], s.u[i+1]) {
			break
		}
		s.u[i+1], s.u[i] = s.u[i], s.u[i+1]
	}
	// we did it.
}

func (s *SortedList[T]) Search(condition cond[T]) (obj T, idx int, found bool) {
	idx = sort.Search(len(s.u), func(i int) bool {
		return condition(s.u[i])
	})
	if idx >= len(s.u) {
		return
	}
	found = true
	obj = s.u[idx]
	return
}

func (s *SortedList[T]) Remove(idx int) {
	s.u = append(s.u[:idx], s.u[idx+1:]...)
}

func (s *SortedList[T]) UnderlyingCopy() (cpy []T) {
	cpy = make([]T, len(s.u))
	copy(cpy, s.u)
	return
}

func (s *SortedList[T]) Range(f func(val T, idx, l int) bool) {
	for i, e := range s.u {
		if !f(e, i, len(s.u)) {
			break
		}
	}
}
