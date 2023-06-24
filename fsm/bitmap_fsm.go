package fsm

import "github.com/Giulio2002/sharedpool/types"

const defaultWordSize = 128

type bitmapFreeSpaceManager struct {
	bitmap   *types.Bitmap
	wordSize int
}

func NewBitmapFreeSpaceManager() FreeSpaceManager {
	return &bitmapFreeSpaceManager{
		bitmap:   types.NewBitmap(),
		wordSize: defaultWordSize,
	}
}

func NewBitmapFreeSpaceManagerWithWordSize(wordSize int) FreeSpaceManager {
	return &bitmapFreeSpaceManager{
		bitmap:   types.NewBitmap(),
		wordSize: wordSize,
	}
}

func (b *bitmapFreeSpaceManager) MarkBusy(startPos, count int) {
	currentWord := startPos / b.wordSize
	wordCount := (count + b.wordSize - 1) / b.wordSize
	for i := currentWord; i < currentWord+wordCount; i++ {
		b.bitmap.Set(i, true)
	}
}

func (b *bitmapFreeSpaceManager) MarkFree(startPos, count int) {
	currentWord := startPos / b.wordSize
	wordCount := (count + b.wordSize - 1) / b.wordSize
	for i := currentWord; i < currentWord+wordCount; i++ {
		b.bitmap.Set(i, false)
	}
}

func (b *bitmapFreeSpaceManager) FirstFreeIndex(count int) int {
	wordCount := (count + b.wordSize - 1) / b.wordSize
	// just find first word count successive 0s bits
	bitsCount := 0
	firstIndex := 0

	for current := 0; bitsCount != wordCount; current++ {
		if !b.bitmap.Get(current) {
			bitsCount++
			if bitsCount == 1 {
				firstIndex = current
			}
			continue
		}
		bitsCount = 0
	}
	return firstIndex * b.wordSize
}
