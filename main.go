package main

// Initialize a simple ping pong web server
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})
	http.ListenAndServe(":8080", nil)
}