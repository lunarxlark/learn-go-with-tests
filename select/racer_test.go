package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayServer(20 * time.Millisecond)
		defer slowServer.Close()
		fastServer := makeDelayServer(0 * time.Millisecond)
		defer fastServer.Close()

		want := fastServer.URL

		got, err := Racer(slowServer.URL, fastServer.URL)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayServer(25 * time.Microsecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Second)
		if err != nil {
			t.Error("expected an error but didn't get one")
		}
	})

}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}