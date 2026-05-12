package forkyporkies

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"time"
)

const GITHUB_API_URL = "https://api.github.com"
const GITHUB_API_VERSION = "2022-11-28"

type Repo struct {
	Owner string
	Name  string
}

func (r Repo) String() string {
	return fmt.Sprintf("%s/%s", r.Owner, r.Name)
}

func FromFork(f Fork) Repo {
	return Repo{
		Owner: f.Owner.Login,
		Name:  f.Name,
	}
}

type Fork struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Owner    struct {
		Login string `json:"login"`
	} `json:"owner"`
}

type Commit struct {
	Author struct {
		Login string `json:"login"`
	} `json:"author"`
	Commit struct {
		Committer struct {
			Date time.Time `json:"date"`
		} `json:"committer"`
	} `json:"commit"`
}

func (r *Repo) GetForks(token string) ([]Fork, error) {
	var forks []Fork
	for page := 1; ; page++ {
		body, err := get(fmt.Sprintf("repos/%s/%s/forks?page=%d", r.Owner, r.Name, page), token)
		if err != nil {
			return nil, fmt.Errorf("get forks for %s/%s: %v", r.Owner, r.Name, err)
		}
		var pageForks []Fork
		err = json.Unmarshal(body.Bytes(), &pageForks)
		if err != nil {
			return pageForks, fmt.Errorf("unmarshal forks: %v", err)
		}
		if len(pageForks) == 0 {
			break
		}
		forks = append(forks, pageForks...)
	}
	return forks, nil
}

func (r *Repo) GetCommits(token string) ([]Commit, error) {
	body, err := get(fmt.Sprintf("repos/%s/%s/commits", r.Owner, r.Name), token)
	if err != nil {
		return nil, fmt.Errorf("get commits for %s/%s: %v", r.Owner, r.Name, err)
	}
	var commits []Commit
	b := body.Bytes()
	err = json.Unmarshal(b, &commits)
	if err != nil {
		return commits, fmt.Errorf("unmarshal commits: %v", err)
	}
	slices.SortFunc(commits, func(l Commit, r Commit) int {
		return l.Commit.Committer.Date.Compare(r.Commit.Committer.Date)
	})
	return commits, nil
}

func get(resource, token string) (*bytes.Buffer, error) {
	url := fmt.Sprintf("%s/%s", GITHUB_API_URL, resource)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("get %s: %v", resource, err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", GITHUB_API_VERSION)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get %s: %v", resource, err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get %s: expected %d, got %d", resource, http.StatusOK, res.StatusCode)
	}
	defer res.Body.Close()
	buf := bytes.NewBufferString("")
	_, err = io.Copy(buf, res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body from response: %v", err)
	}
	return buf, nil
}
