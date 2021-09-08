package main

import (
	"context"
	"net/http"
)

var contextKeyAPIKey = &contextKey{"api-key"}

type contextKey struct {
	name string
}

func APIkey(ctx context.Context) (string, bool) {
	key, ok := ctx.Value(contextKeyAPIKey).(string)
	return key, ok
}

func isValidAPIkey(key string) bool {
	return key == "abc123"
}

func withAPIkey(fn http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		key := request.URL.Query().Get("key")
		if !isValidAPIkey(key) {
			respondErr(writer, request, http.StatusUnauthorized, "invalid API key")
			return
		}
		ctx := context.WithValue(request.Context(), contextKeyAPIKey, key)
		fn(writer, request.WithContext(ctx))
	}
}

func main() {

}
