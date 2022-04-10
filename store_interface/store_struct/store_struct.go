package store_struct

import (
	"fmt"
	"sync"
)

type storeStruct struct {
	store_map map[string]string
}

func CreateStoreStruct(store_map map[string]string) storeStruct {
	if store_map == nil {
		panic("|->|Your store is empty please re-enter it! |<-|")
	}
	return storeStruct{store_map: store_map}
}

func (s *storeStruct) Get(key string, wg *sync.WaitGroup, myMutex *sync.Mutex) {
	myMutex.Lock()
	defer myMutex.Unlock()
	defer wg.Done()
	if _, err := s.store_map[key]; !err {
		fmt.Printf("|->|We don't have an item with id %v |<-|\n" + key)
		return
	}
	fmt.Printf("|##|Item with id %v is %v |##|\n", key, s.store_map[key])
}

func (s *storeStruct) Put(key string, value string, wg *sync.WaitGroup, myMutex *sync.Mutex) {
	myMutex.Lock()
	defer myMutex.Unlock()
	defer wg.Done()
	if key == "" || value == "" {
		panic("|->|Please check your key and value, and re-enter them again|<-|")
	}
	for k, v := range s.store_map {
		if k == key {
			fmt.Printf("|><|We have updated item with ID %v from %v to %v.|><|\n", key, v, value)
			s.store_map[key] = value
			return
		}
	}
	//We have added an item, named bread with the number 8 to our shop
	fmt.Printf("|<>|We have added an item %v with ID %v to the store.|<>|\n", value, key)
	s.store_map[key] = value
}
