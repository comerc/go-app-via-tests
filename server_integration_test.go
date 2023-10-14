package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	poker "github.com/comerc/go-app-via-tests"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	file, removeFile := mustCreateTempFile(t, `[]`)
	defer removeFile()
	store, err := poker.NewFileSystemPlayerStore(file)
	AssertNoError(t, err)
	// store := NewInMemoryPlayerStore()
	server := mustMakePlayerServer(t, store, dummyGame)
	player := "Pepper"
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		AssertStatus(t, response, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "3")
	})
	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		AssertStatus(t, response, http.StatusOK)
		got := getLeagueFromResponse(t, response.Body)
		want := poker.League{{player, 3}}
		AssertLeague(t, got, want)
	})
}
