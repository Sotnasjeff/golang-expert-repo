package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started")
	defer log.Println("Request finished")
	select {
	case <-time.After(5 * time.Second):
		log.Println("Request proccessed successfully")
		w.Write([]byte("Request proccessed successfully"))
	case <-ctx.Done():
		log.Println("Request cancelled by user")
	}
}
