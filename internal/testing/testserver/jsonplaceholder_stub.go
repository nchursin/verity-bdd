package testserver

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func StartJSONPlaceholderStub(t testing.TB) string {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/posts":
			_, _ = w.Write([]byte(`[{"id":1,"userId":1,"title":"sunt aut facere","body":"quia et suscipit"}]`))
		case r.Method == http.MethodGet && r.URL.Path == "/posts/1":
			_, _ = w.Write([]byte(`{"id":1,"userId":1,"title":"sunt aut facere","body":"quia et suscipit"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/posts/2":
			_, _ = w.Write([]byte(`{"id":2,"userId":1,"title":"qui est esse","body":"est rerum tempore"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/users":
			_, _ = w.Write([]byte(`[{"id":1,"name":"Leanne Graham","email":"leanne@example.com"}]`))
		case r.Method == http.MethodPost && r.URL.Path == "/posts":
			payload := map[string]any{}
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, `{"error":"failed to read body"}`, http.StatusBadRequest)
				return
			}
			if err := json.Unmarshal(body, &payload); err != nil {
				http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
				return
			}
			payload["id"] = 101
			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(payload)
		default:
			http.NotFound(w, r)
		}
	}))

	t.Cleanup(server.Close)

	return server.URL
}
