package poker

import "sync"

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct {
	store map[string]int
	// A mutex is used to synchronize read/write access to the map
	lock sync.RWMutex
}

// NewInMemoryPlayerStore initialises an empty player store.
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.RWMutex{},
	}
}

// RecordWin will record a player's win.
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.store[name]++
}

// GetPlayerScore retrieves scores for a given player.
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() League {
	i.lock.RLock()
	defer i.lock.RUnlock()
	var league League
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
