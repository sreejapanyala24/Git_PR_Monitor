package main

type PRPayload struct {
	Action string `json:"action"`

	PullRequest struct {
		Number    int    `json:"number"`
		Title     string `json:"title"`
		CreatedAt string `json:"created_at"`
		URL       string `json:"html_url"`

		User struct {
			Login string `json:"login"`
		} `json:"user"`
	} `json:"pull_request"`
}

type PR struct {
	Number    int    `json:"number"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	URL       string `json:"url"`
	Author    string `json:"author"`
}
