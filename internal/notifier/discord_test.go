package notifier_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lbAntoine/ssh-portfolio/internal/notifier"
)

func TestSendDiscord_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()
	if err := notifier.SendDiscord(srv.URL, "Alice", "alice@test.com", "hello"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestSendDiscord_Non2xxReturnsError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer srv.Close()
	if err := notifier.SendDiscord(srv.URL, "Alice", "alice@test.com", "hello"); err == nil {
		t.Error("expected error for non-2xx status")
	}
}

func TestSendDiscord_InvalidURLReturnsError(t *testing.T) {
	if err := notifier.SendDiscord("://bad-url", "Alice", "alice@test.com", "hello"); err == nil {
		t.Error("expected error for invalid URL")
	}
}

func TestSendDiscord_EmbedTitleContainsName(t *testing.T) {
	var body []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()
	if err := notifier.SendDiscord(srv.URL, "Alice", "alice@test.com", "hello"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(string(body), "Alice from ssh-portfolio") {
		t.Errorf("expected embed title to contain 'Alice from ssh-portfolio', got: %s", body)
	}
}

func TestSendDiscord_PayloadContainsEmail(t *testing.T) {
	var body []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()
	if err := notifier.SendDiscord(srv.URL, "Alice", "alice@test.com", "hello"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(string(body), "alice@test.com") {
		t.Errorf("expected payload to contain email, got: %s", body)
	}
}

func TestSendDiscord_ContentTypeIsJSON(t *testing.T) {
	var ct string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ct = r.Header.Get("Content-Type")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()
	_ = notifier.SendDiscord(srv.URL, "Alice", "alice@test.com", "hello")
	if !strings.HasPrefix(ct, "application/json") {
		t.Errorf("expected application/json content-type, got: %s", ct)
	}
}
