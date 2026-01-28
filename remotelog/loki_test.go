package remotelog

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sipcapture/heplify-server/config"
)

func withConfig(t *testing.T, fn func()) {
	t.Helper()
	original := config.Setting
	defer func() {
		config.Setting = original
	}()
	fn()
}

func TestLokiSetupSendsOrgIDHeader(t *testing.T) {
	withConfig(t, func() {
		headerCh := make(chan string, 1)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != getPath {
				t.Errorf("unexpected path %s", r.URL.Path)
			}
			headerCh <- r.Header.Get("X-Scope-OrgID")
			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		config.Setting.LokiURL = server.URL
		config.Setting.LokiOrgID = "heplify-org-id"

		loki := &Loki{}
		if err := loki.setup(); err != nil {
			t.Fatalf("setup failed: %v", err)
		}

		select {
		case got := <-headerCh:
			if got != "heplify-org-id" {
				t.Errorf("expected X-Scope-OrgID header %q, got %q", "heplify-org-id", got)
			}
		default:
			t.Fatal("expected label check request")
		}
	})
}

func TestLokiSetupFailsOnNon2xx(t *testing.T) {
	withConfig(t, func() {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusForbidden)
		}))
		defer server.Close()

		config.Setting.LokiURL = server.URL

		loki := &Loki{}
		if err := loki.setup(); err == nil {
			t.Fatal("expected setup to fail on non-2xx status")
		}
	})
}

func TestLokiSendSendsOrgIDHeaderAndFailsOnNon2xx(t *testing.T) {
	withConfig(t, func() {
		headerCh := make(chan string, 1)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != postPathOne {
				t.Errorf("unexpected path %s", r.URL.Path)
			}
			headerCh <- r.Header.Get("X-Scope-OrgID")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("nope"))
		}))
		defer server.Close()

		config.Setting.LokiOrgID = "heplify-org-id"

		loki := &Loki{URL: server.URL + postPathOne}
		if _, err := loki.send(context.Background(), []byte("payload")); err == nil {
			t.Fatal("expected send to fail on non-2xx status")
		}

		select {
		case got := <-headerCh:
			if got != "heplify-org-id" {
				t.Errorf("expected X-Scope-OrgID header %q, got %q", "heplify-org-id", got)
			}
		default:
			t.Fatal("expected push request")
		}
	})
}
