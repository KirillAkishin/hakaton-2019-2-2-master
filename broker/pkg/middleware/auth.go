package middleware

import (
	"context"
	"net/http"
	"broker/pkg/session"
)

var (
	// урлы, которые не требуют авторизации и не учитывают её наличие никак
	noAuthUrls = map[string]struct{}{
		"/api/v1/login":    {},
		"/api/v1/register": {},
	}

	// урлы и методы, которые требуют наличие сессии.
	requireSessUrlsMethods = map[string]map[string]struct{}{
		"/api/v1/status": {
			"":             {},  // GET too
			http.MethodGet: {},
		},
		"/api/v1/deal": {
			http.MethodPost:   {},
		},
		"/api/v1/cancel": {
			http.MethodPost:   {},
		},
		"/api/v1/history": {
			"":             {}, // GET too
			http.MethodGet: {},
		},
	}
)

// выяснить, есть ли авторизация, если она требуется или учитывается, и принять соотв меры
func Auth(sm *session.SessionsManager, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := noAuthUrls[r.URL.Path]; ok {
			next.ServeHTTP(w, r)
			return
		}
		client, err := sm.Check(r)
		requireSess := false
		methods, found := requireSessUrlsMethods[r.URL.Path]
		if !found {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		_, requireSess = methods[r.Method]
		if err != nil && requireSess {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		ctx := context.WithValue(r.Context(), session.SessionKey, client)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
