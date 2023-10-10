package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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

// file_system_store.go
func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}
	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}
	return nil
}

// NewFileSystemPlayerStore initialises an empty player store.
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}
	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}
	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		lock:     sync.RWMutex{},
		league:   league,
	}, nil
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
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}
