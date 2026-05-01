# PR Monitor (Go Webhook Service)

## Overview

This project is a Go-based service that monitors GitHub Pull Request events using webhooks and maintains a list of currently open PRs.

The service listens for `pull_request` events, updates internal state, and generates a readable output file with active PR details.

---

## Features

* Real-time PR tracking using GitHub webhooks
* Handles PR lifecycle events:

    * `opened`
    * `reopened`
    * `closed`
* Maintains current PR state in JSON
* Generates a human-readable text file
* Webhook signature validation for security
* Local testing using ngrok

---

## Architecture

```text
GitHub → Webhook → Go Server → State (JSON) → Output (TXT)
```

Flow:

1. GitHub sends a webhook event (`pull_request`)
2. Go server receives the request at `/webhook`
3. Signature is validated using a secret
4. Based on action:

    * opened/reopened → add PR
    * closed → remove PR
5. JSON state is updated
6. TXT file is regenerated

---

## Project Structure

```
GIT_PR_Monitor
├── main.go        # Starts HTTP server
├── handler.go     # Webhook handler
├── state.go       # JSON state management
├── writer.go      # TXT file generation
├── security.go    # Signature validation
├── models.go      # Data structures
├── open_prs.json  # Stored state (auto-generated)
└── pull_requests.txt # Output file (auto-generated)
```

---

## Key Concepts

### Webhooks

GitHub sends real-time HTTP POST requests when PR events occur.

### State Management

`open_prs.json` stores the current open PRs.

### Output

`pull_requests.txt` is generated from JSON for readability.

---

## Security

Webhook requests are validated using:

* `X-Hub-Signature-256`
* Shared secret (`WEBHOOK_SECRET`)

---

## Setup & Run

### 1. Clone repository

```
git clone <repo-url>
cd <repo-name>
```

---

### 2. Set environment variable

#### Linux / WSL / Mac

```
export WEBHOOK_SECRET="your-secret"
```

#### Windows (PowerShell)

```
$env:WEBHOOK_SECRET="your-secret"
```

---

### 3. Run the server

```
go run .
```

Server starts on:

```
http://localhost:9090
```

---

### 4. Start ngrok

```
ngrok http 9090
```

You will get a public URL like:

```
https://abc123.ngrok.io
```

---

### 5. Configure GitHub Webhook

Go to:

```
Repo → Settings → Webhooks → Add webhook
```

Set:

* Payload URL:

  ```
  https://<ngrok-url>/webhook
  ```
* Content type:

  ```
  application/json
  ```
* Secret:

  ```
  same as WEBHOOK_SECRET
  ```
* Events:

  ```
  Pull requests
  ```

---

## Testing

1. Create a new Pull Request → should be added
2. Close the PR → should be removed
3. Verify files:

```
open_prs.json
pull_requests.txt
```

---

## Example Output

```
CURRENT OPEN PULL REQUESTS
==========================

PR Number: 2
Title: add city
Created At: 2026-05-01T18:50:13Z
PR Link: https://github.com/sreejapanyala24/Git_PR_Monitor_Test/pull/2
Raised By: sreejapanyala24
----------------------------------------
```

---

## Notes

* Only open PRs are tracked
* Closed PRs are removed immediately
* TXT file is regenerated on every event

---

