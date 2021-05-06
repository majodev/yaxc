package server

type Backend interface {
	Get(key string) (string, error)
	Set(key, value string) error

	GetHash(key string) (string, error)
	SetHash(key, value string) error
}
