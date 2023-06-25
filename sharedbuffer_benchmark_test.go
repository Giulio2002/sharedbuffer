package sharedbuffer

import (
	"testing"

	"github.com/Giulio2002/sharedbuffer/fsm"
	"github.com/Giulio2002/sharedbuffer/management"
)

func BenchmarkContiguousLongLived(b *testing.B) {
	maxRequest := 64
	f := NewSharedBuffer(fsm.NewCountingFreeSpaceManager(), management.NewMemoryBuffer())
	for i := 0; i < b.N; i++ {
		f.Make(b.N % maxRequest)
	}
}

func BenchmarkBitmapLongLived(b *testing.B) {
	maxRequest := 64
	f := NewSharedBuffer(fsm.NewBitmapFreeSpaceManagerWithWordSize(64), management.NewMemoryBuffer())
	for i := 0; i < b.N; i++ {
		f.Make(b.N % maxRequest)
	}
}

func BenchmarkITreeLongLived(b *testing.B) {
	maxRequest := 64
	f := NewSharedBuffer(fsm.NewIntervalTreeFreeSpaceManager(), management.NewMemoryBuffer())
	for i := 0; i < b.N; i++ {
		f.Make(b.N % maxRequest)
	}
}

func BenchmarkContiguousShortLived(b *testing.B) {
	f := NewSharedBuffer(fsm.NewCountingFreeSpaceManager(), management.NewMemoryBuffer())
	maxRequest := 64
	batchSize := 64
	for i := 0; i < b.N/batchSize; i++ {
		cs := []func(){}
		for j := 0; j < batchSize; j++ {
			_, cf := f.Make(b.N % maxRequest)
			cs = append(cs, cf)
		}
		for _, cf := range cs {
			cf()
		}
	}
}

func BenchmarkBitmapShortLived(b *testing.B) {

	f := NewSharedBuffer(fsm.NewBitmapFreeSpaceManagerWithWordSize(64), management.NewMemoryBuffer())

	maxRequest := 64
	batchSize := 64
	for i := 0; i < b.N/batchSize; i++ {
		cs := []func(){}
		for j := 0; j < batchSize; j++ {
			_, cf := f.Make(b.N % maxRequest)
			cs = append(cs, cf)
		}
		for _, cf := range cs {
			cf()
		}
	}
}

func BenchmarkITreeShortLived(b *testing.B) {
	f := NewSharedBuffer(fsm.NewIntervalTreeFreeSpaceManager(), management.NewMemoryBuffer())
	maxRequest := 64
	batchSize := 64
	for i := 0; i < b.N/batchSize; i++ {
		cs := []func(){}
		for j := 0; j < batchSize; j++ {
			_, cf := f.Make(b.N % maxRequest)
			cs = append(cs, cf)
		}
		for _, cf := range cs {
			cf()
		}
	}
}
