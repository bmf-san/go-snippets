package main

import (
	"errors"
	"fmt"
	"sync"
)

type Data struct {
	value string
}

type SyncMap struct {
	s sync.Map
}

func NewSyncMap() SyncMap {
	return SyncMap{}
}

func (s *SyncMap) Store(key int, value Data) {
	s.s.Store(key, value)
}

func (s *SyncMap) Load(key int) (Data, error) {
	v, ok := s.s.Load(key)
	if !ok {
		return Data{}, errors.New("Not Found")
	}

	t, ok := v.(Data)
	if !ok {
		return Data{}, errors.New("Type is invalid")
	}

	return t, nil
}

func main() {
	// Sync.Map supports multi-threaded concurrent read and write, which is better than the previous lock map performance.
	s := NewSyncMap()
	s.Store(1, Data{
		value: "Hello World",
	})

	t, err := s.Load(1)
	fmt.Println(t, err)
}
