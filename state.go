package main

import (
	"encoding/json"
	"io/ioutil"
)

var stateFile = "open_prs.json"

func LoadState() map[int]PR {
	data, err := ioutil.ReadFile(stateFile)
	if err != nil {
		return make(map[int]PR)
	}

	var prs map[int]PR
	json.Unmarshal(data, &prs)

	return prs
}

func SaveState(prs map[int]PR) {
	data, _ := json.MarshalIndent(prs, "", "  ")
	ioutil.WriteFile(stateFile, data, 0644)
}

func AddPR(payload PRPayload) {
	prs := LoadState()

	prs[payload.PullRequest.Number] = PR{
		Number:    payload.PullRequest.Number,
		Title:     payload.PullRequest.Title,
		CreatedAt: payload.PullRequest.CreatedAt,
		URL:       payload.PullRequest.URL,
		Author:    payload.PullRequest.User.Login,
	}

	SaveState(prs)
}

func RemovePR(prNumber int) {
	prs := LoadState()
	delete(prs, prNumber)
	SaveState(prs)
}
