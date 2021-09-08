package main

import (
	"context"
	"flag"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

var contextKeyAPIKey = &contextKey{"api-key"}

type contextKey struct {
	name string
}

func APIKey(ctx context.Context) (string, bool) {
	key, ok := ctx.Value(contextKeyAPIKey).(string)
	return key, ok
}

func isValidAPIkey(key string) bool {
	return key == "abc123"
}

func withAPIKey(fn http.HandlerFunc) http.HandlerFunc {
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

func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Location")
		fn(w, r)
	}
}

type Server struct {
	db *mgo.Session
}

func main() {
	var (
		addr  = flag.String("addr", ":8080", "endpoint address")
		mongo = flag.String("mongo", "localhost", "mongodb address")
	)
	log.Println("Dialing mongo", *mongo)
	db, err := mgo.Dial(*mongo)
	if err != nil {
		log.Fatalln("failed to connect to mongo:", err)
	}
	defer db.Close()
	s := &Server{
		db: db,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/polls/", withCORS(withAPIKey(s.handlePolls)))
	log.Println("Starting web server on", *addr)
	http.ListenAndServe(":8080", mux)
	log.Println("Stopping...")
}
