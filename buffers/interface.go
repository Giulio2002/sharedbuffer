package management

type freeFunc func()

type Buffer interface {
	Get(n int) ([]byte, freeFunc)
}
