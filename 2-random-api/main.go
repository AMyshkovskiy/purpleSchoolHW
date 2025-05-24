package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	newMux := http.NewServeMux()

	newMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := rand.Intn(6) + 1
		w.Write([]byte(strconv.Itoa(result)))
		w.Header().Set("Content-Type", "text/plain")
		return
	})
	server := http.Server{
		Addr:    ":8081",
		Handler: newMux,
	}

	server.ListenAndServe()
}
