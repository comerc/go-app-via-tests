package main

import (
	"encoding/json"
	"os"
	"sync"
)

const dbFileName = "game.db.json"

// FileSystemPlayerStore collects data about players in file.
type FileSystemPlayerStore struct {
	database *json.Encoder
	// A mutex is used to synchronize read/write access to the map
	lock   sync.RWMutex
	league League
}

// NewFileSystemPlayerStore initialises an empty player store.
func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	file.Seek(0, 0)
	league, _ := NewLeague(file)
	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		lock:     sync.RWMutex{},
		league:   league,
	}
}

// RecordWin will record a player's win.
func (f *FileSystemPlayerStore) RecordWin(name string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	player := f.league.Find(name)
	if player == nil {
		f.league = append(f.league, Player{name, 1})
	} else {
		player.Wins++
	}
	f.database.Encode(f.league)
}

// GetPlayerScore retrieves scores for a given player.
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	f.lock.RLock()
	defer f.lock.RUnlock()
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}
