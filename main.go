package main

import (
	"fmt"
	"sync"

	"github.com/M1racle-Heen/golang_midterm/store_interface/store_struct"
)

func main() {
	store()
}

var (
	storeMap   map[string]string
	wg         sync.WaitGroup
	m          sync.Mutex
	valueToPut map[string]string
)

func store() {
	storeMap = map[string]string{"1": "Cola", "2": "Bacon", "3": "Egg", "4": "Apple"}
	store := store_struct.CreateStoreStruct(storeMap)
	valueToPut = map[string]string{"3": "apple", "5": "Bread", "6": "Sugar", "7": "Soda", "8": "Banana", "9": "Swag"}
	wg.Add(12)
	for k, v := range valueToPut {
		go func() {
			store.Get(k, &wg, &m)
		}()
		store.Put(k, v, &wg, &m)
	}
	wg.Wait()

	fmt.Println(store)
}
