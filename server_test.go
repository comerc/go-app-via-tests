package poker

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLeague(t *testing.T) {
	t.Run("it returns 200 on /league", func(t *testing.T) {
		wantedLeague := League{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		store := &StubPlayerStore{
			nil,
			nil,
			wantedLeague,
		}
		server := NewPlayerServer(store)
		request := newLeagueRequest()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		AssertContentType(t, response, jsonContentType)
		got := getLeagueFromResponse(t, response.Body)
		AssertStatus(t, response.Code, http.StatusOK)
		AssertLeague(t, got, wantedLeague)
	})
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league League) {
	t.Helper()
	league, _ = NewLeague(body)
	return
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func TestStoreWins(t *testing.T) {
	t.Run("it records wins when POST", func(t *testing.T) {
		store := &StubPlayerStore{
			nil,
			nil,
			nil,
		}
		server := NewPlayerServer(store)
		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		AssertStatus(t, response.Code, http.StatusAccepted)
		AssertPlayerWin(t, store, player)
	})
}

func TestGETPlayers(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}
	server := NewPlayerServer(store)
	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		AssertStatus(t, response.Code, http.StatusNotFound)
	})
	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		AssertStatus(t, response.Code, http.StatusOK)
		got := response.Body.String()
		want := "20"
		AssertResponseBody(t, got, want)
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		AssertStatus(t, response.Code, http.StatusOK)
		got := response.Body.String()
		want := "10"
		AssertResponseBody(t, got, want)
	})
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}
