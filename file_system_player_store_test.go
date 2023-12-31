package poker_test

import (
	"os"
	"testing"

	poker "github.com/comerc/go-app-via-tests"
)

func TestFileSystemPlayerStore(t *testing.T) {
	t.Run("works with an empty file", func(t *testing.T) {
		file, removeFile := mustCreateTempFile(t, "")
		defer removeFile()
		_, err := poker.NewFileSystemPlayerStore(file)
		AssertNoError(t, err)
	})
	file, removeFile := mustCreateTempFile(t, `[{"Name":"Cleo","Wins":10},{"Name":"Chris","Wins": 33}]`)
	defer removeFile()
	store, err := poker.NewFileSystemPlayerStore(file)
	AssertNoError(t, err)
	t.Run("league sorted", func(t *testing.T) {
		got := store.GetLeague()
		want := poker.League{
			{Name: "Chris", Wins: 33}, {Name: "Cleo", Wins: 10},
		}
		AssertLeague(t, got, want)
	})
	t.Run("get player score for exist player", func(t *testing.T) {
		got := store.GetPlayerScore("Chris")
		want := 33
		AssertScoreEquals(t, got, want)
	})
	t.Run("store wins for existing players", func(t *testing.T) {
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		want := 34
		AssertScoreEquals(t, got, want)
	})
	t.Run("get player score for unknown player", func(t *testing.T) {
		got := store.GetPlayerScore("Pepper")
		want := 0
		AssertScoreEquals(t, got, want)
	})
	t.Run("store wins for new players", func(t *testing.T) {
		store.RecordWin("Pepper")
		got := store.GetPlayerScore("Pepper")
		want := 1
		AssertScoreEquals(t, got, want)
	})
}

func mustCreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "temp_"+poker.DBFileName)
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
