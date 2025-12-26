package forkyporkies

import (
	"fmt"
	"os"
	"strings"
)

func Parse(path string) ([]Repo, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %v", path, err)
	}
	lines := strings.Lines(string(data))
	repos := make([]Repo, 0)
	for l := range lines {
		repo := strings.Split(strings.TrimSpace(l), "/")
		if len(repo) != 2 {
			fmt.Fprintf(os.Stderr, "invalid repo identifier '%s'", l)
			continue
		}
		repos = append(repos, Repo{Owner: repo[0], Name: repo[1]})
	}
	return repos, nil
}
