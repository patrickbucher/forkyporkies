package main

import (
	"fmt"
	"os"

	fp "github.com/patrickbucher/forkyporkies"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Fprintln(os.Stderr, "GITHUB_TOKEN environment variable must be set")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s REPO_FILE", os.Args[0])
		os.Exit(1)
	}
	repos, err := fp.Parse(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse repo file: %v", err)
		os.Exit(1)
	}
	for _, r := range repos {
		forks, err := r.GetForks(token)
		if err != nil {
			fmt.Fprintf(os.Stderr, "get fork for %s: %v", r, err)
			continue
		}
		for _, f := range forks {
			fmt.Println(f)
		}
	}
}
