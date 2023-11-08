package backend

type Backend interface {
	Put(url string, data []byte) error
	Get(url string) ([]byte, error)
}
