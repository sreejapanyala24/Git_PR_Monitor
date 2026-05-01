package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {

	// Step 1: Ensure method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Step 2: Read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read request body", http.StatusInternalServerError)
		return
	}

	// Step 3: Parse JSON payload
	var payload PRPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	log.Println("Received action:", payload.Action)

	// Step 4: Handle PR actions
	switch payload.Action {

	case "opened", "reopened":
		log.Println("Adding PR:", payload.PullRequest.Number)
		AddPR(payload)

	case "closed":
		log.Println("Removing PR:", payload.PullRequest.Number)
		RemovePR(payload.PullRequest.Number)

	default:
		log.Println("Ignoring action:", payload.Action)
	}

	// Step 5: Rewrite text file
	err = RewriteTextFile()
	if err != nil {
		log.Println("Error writing file:", err)
	}

	// Step 6: Send success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook processed"))
}
