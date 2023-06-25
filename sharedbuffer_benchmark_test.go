package sharedbuffer

import (
	"testing"

	"github.com/Giulio2002/sharedbuffer/fsm"
	"github.com/Giulio2002/sharedbuffer/management"
)

func BenchmarkContiguousLongLived(b *testing.B) {
	maxRequest := 64
	f := NewSimpleSharedBuffer(fsm.NewCountingFreeSpaceManager(), management.NewMemoryBuffer())
	for i := 0; i < b.N; i++ {
		f.Make(i % maxRequest)
	}
}

func BenchmarkBitmapLongLived(b *testing.B) {
	maxRequest := 64
	f := NewSimpleSharedBuffer(fsm.NewBitmapFreeSpaceManagerWithWordSize(64), management.NewMemoryBuffer())
	for i := 0; i < b.N; i++ {
		f.Make(i % maxRequest)
	}
}

func BenchmarkITreeLongLived(b *testing.B) {
	maxRequest := 64
	f := NewSimpleSharedBuffer(fsm.NewIntervalTreeFreeSpaceManager(), management.NewMemoryBuffer())
	for i := 0; i < b.N; i++ {
		f.Make(i % maxRequest)
	}
}

func BenchmarkContiguousShortLived(b *testing.B) {
	f := NewSimpleSharedBuffer(fsm.NewCountingFreeSpaceManager(), management.NewMemoryBuffer())
	maxRequest := 100000
	for i := 0; i < b.N; i++ {
		_, cf := f.Make(maxRequest)
		cf()
	}
}

func BenchmarkBitmapShortLived(b *testing.B) {
	f := NewSimpleSharedBuffer(fsm.NewBitmapFreeSpaceManagerWithWordSize(64), management.NewMemoryBuffer())
	maxRequest := 100000
	for i := 0; i < b.N; i++ {
		_, cf := f.Make(maxRequest)
		cf()
	}
}

func BenchmarkITreeShortLived(b *testing.B) {
	f := NewSimpleSharedBuffer(fsm.NewIntervalTreeFreeSpaceManager(), management.NewMemoryBuffer())
	maxRequest := 100000
	for i := 0; i < b.N; i++ {
		_, cf := f.Make(maxRequest)
		cf()
	}
}

func BenchmarkAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]byte, 100000)
	}
}
