package store_struct

type StoreStruct struct {
	store_map map[string]string
}

func CreateStoreStruct(store_map map[string]string) StoreStruct {
	if store_map == nil {
		panic("Your store is empty please re-enter it")
	}
	return StoreStruct{store_map: store_map}
}

func (s *StoreStruct) Get(key string) (string, error) {
	return s.store_map[key], nil
}

func (s *StoreStruct) Put(key string, value string) error {
	return nil
}
