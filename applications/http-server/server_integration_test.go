package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	http_server "applications/http-server"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := http_server.NewInMemoryPlayerStore()
	server := http_server.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("Get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("Get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, http_server.JsonContentType)

		got := getLeagueFromResponse(t, response.Body)
		want := []http_server.Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})
}
