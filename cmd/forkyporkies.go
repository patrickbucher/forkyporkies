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
		fmt.Fprintf(os.Stderr, "usage: %s REPO_FILE\n", os.Args[0])
		os.Exit(1)
	}
	repos, err := fp.Parse(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse repo file: %v\n", err)
		os.Exit(1)
	}
	t := fetchForkCommits(repos, token)
	fmt.Println(t)
}

func fetchForkCommits(repos []fp.Repo, token string) fp.Table {
	t := fp.Table{
		Repos: make([]string, 0),
		Forks: make(map[string]fp.Entry, 0),
	}
	for _, r := range repos {
		t.Repos = append(t.Repos, r.String())
		forks, err := r.GetForks(token)
		if err != nil {
			fmt.Fprintf(os.Stderr, "get fork for %s: %v\n", r, err)
			continue
		}
		for _, f := range forks {
			fr := fp.FromFork(f)
			login := f.Owner.Login
			commits, err := fr.GetCommits(token)
			if err != nil {
				fmt.Fprintf(os.Stderr, "get commits for %v: %v", fr, err)
				continue
			}
			if _, ok := t.Forks[login]; !ok {
				t.Forks[login] = fp.Entry{
					Author:  login,
					Commits: make(map[string]uint, 0),
				}
			}
			var nCommits uint = 0
			for _, c := range commits {
				if c.Author.Login != r.Owner {
					nCommits++
				}
			}
			t.Forks[login].Commits[r.String()] = nCommits
		}
	}
	return t
}
