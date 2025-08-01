package cache

type Cache interface {
	Get(key string) (*float64, error)
	Set(key string, value float64) error
}
