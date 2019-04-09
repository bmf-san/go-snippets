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
	mu sync.Mutex
	m  map[int]Data
}

func NewSyncMap() SyncMap {
	return SyncMap{m: map[int]Data{}}
}

func (s *SyncMap) Store(key int, value Data) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.m[key] = value
}

func (s *SyncMap) Load(key int) (Data, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, ok := s.m[key]
	if !ok {
		return Data{}, errors.New("Not Found")
	}
	return t, nil
}

func main() {
	// Implement something like sync.Map with mutex
	s := NewSyncMap()
	s.Store(1, Data{
		value: "Hello World",
	})
	t, err := s.Load(1)
	fmt.Println(t, err)
}
