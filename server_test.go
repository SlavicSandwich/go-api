package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func NewGetScoreRequest(t testing.TB, name string) (*http.Request, error) {
	t.Helper()
	return http.NewRequest(http.MethodGet, "/players/"+name, nil)
}

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := NewGetScoreRequest(t, "Pepper")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := NewGetScoreRequest(t, "Floyd")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}
