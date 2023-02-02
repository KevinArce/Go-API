package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	t.Run("home page", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		HomeHandler(response, request)

		got := response.Body.String()
		want := "Hello World!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
