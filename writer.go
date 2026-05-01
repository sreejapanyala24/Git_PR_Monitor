package main

import (
	"fmt"
	"os"
)

var outputFile = "pull_requests.txt"

func RewriteTextFile() error {
	prs := LoadState()

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintln(file, "CURRENT OPEN PULL REQUESTS")
	fmt.Fprintln(file, "==========================")
	fmt.Fprintln(file)

	if len(prs) == 0 {
		fmt.Fprintln(file, "No open pull requests.")
		return nil
	}

	for _, pr := range prs {
		fmt.Fprintf(file, "PR Number: %d\n", pr.Number)
		fmt.Fprintf(file, "Title: %s\n", pr.Title)
		fmt.Fprintf(file, "Created At: %s\n", pr.CreatedAt)
		fmt.Fprintf(file, "PR Link: %s\n", pr.URL)
		fmt.Fprintf(file, "Raised By: %s\n", pr.Author)
		fmt.Fprintln(file, "----------------------------------------")
	}

	return nil
}
