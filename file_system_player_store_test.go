package main

import (
	"os"
	"testing"
)

func TestFileSystemPlayerStore(t *testing.T) {
	file, removeFile := createTempFile(t, `[{"Name":"Chris","Wins": 33},{"Name":"Cleo","Wins":10}]`)
	defer removeFile()
	store := NewFileSystemPlayerStore(file)
	t.Run("league run from a reader", func(t *testing.T) {
		got := store.GetLeague()
		want := League{
			{Name: "Chris", Wins: 33}, {Name: "Cleo", Wins: 10},
		}
		assertLeague(t, got, want)
	})
	t.Run("get player score for exist player", func(t *testing.T) {
		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEquals(t, got, want)
	})
	t.Run("store wins for existing players", func(t *testing.T) {
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
	})
	t.Run("get player score for unknown player", func(t *testing.T) {
		got := store.GetPlayerScore("Pepper")
		want := 0
		assertScoreEquals(t, got, want)
	})
	t.Run("store wins for new players", func(t *testing.T) {
		store.RecordWin("Pepper")
		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "temp_"+dbFileName)
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	tmpfile.Write([]byte(initialData))
	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}
	return tmpfile, removeFile
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
