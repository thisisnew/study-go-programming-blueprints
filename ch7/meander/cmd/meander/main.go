package main

import (
	"encoding/json"
	"net/http"
	"study-go-programming-blueprints/ch7/meander"
)

func main() {
	http.HandleFunc("/journeys", func(writer http.ResponseWriter, request *http.Request) {
		respond(writer, request, meander.Journeys)
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	return json.NewEncoder(w).Encode(data)
}
