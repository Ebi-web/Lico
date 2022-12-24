package main

import (
	"net/http"

	"lineBot/Controllers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/webhook", Controllers.HandleWebhook)

	http.ListenAndServe(":8080", nil)
}
