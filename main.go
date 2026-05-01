package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", HandleWebhook)

	log.Println("PR monitoring server listening on :9090")

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Server failed to start", err)
	}
}
