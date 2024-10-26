package main

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{store: map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
	mux   sync.Mutex
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mux.Lock()
	defer i.mux.Unlock()
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mux.Lock()
	defer i.mux.Unlock()
	return i.store[name]
}
