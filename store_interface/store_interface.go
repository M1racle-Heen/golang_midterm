package store_interface

import "sync"

type StoreInterface interface {
	Get(key interface{}, wg *sync.WaitGroup, myMutex *sync.Mutex)
	Put(key interface{}, value interface{}, wg *sync.WaitGroup, myMutex *sync.Mutex)
}
