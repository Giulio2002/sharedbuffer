package management

type Buffer interface {
	Get(startPos, n int) []byte
}
