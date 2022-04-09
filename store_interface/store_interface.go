package store_interface

type store_interface interface {
	Get(key string) (interface{}, error)
	Put(key string, value interface{}) error
}
