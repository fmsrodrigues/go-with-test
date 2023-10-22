package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	http_server "applications/http-server"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := http_server.NewInMemoryPlayerStore()
	server := http_server.PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}
