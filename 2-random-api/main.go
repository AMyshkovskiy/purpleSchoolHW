package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	newMux := http.NewServeMux()

	newMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := rand.Intn(6)
		for {
			if result != 0 {
				break
			} else {
				result = rand.Intn(6)
			}
		}
		w.Write([]byte(strconv.Itoa(result)))
		return
	})
	server := http.Server{
		Addr:    ":8082",
		Handler: newMux,
	}

	server.ListenAndServe()
}
