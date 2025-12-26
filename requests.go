package forkyporkies

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
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

type Fork struct{}
type Commit struct{}

func (r *Repo) GetForks(token string) ([]Fork, error) {
	body, err := get(fmt.Sprintf("repos/%s/%s/forks", r.Owner, r.Name), token)
	if err != nil {
		return nil, fmt.Errorf("get forks for %s/%s: %v", r.Owner, r.Name, err)
	}
	fmt.Println(body)
	// FIXME: unmarshal JSON
	return []Fork{}, nil
}

func (r *Repo) GetCommits(token string) ([]Commit, error) {
	body, err := get(fmt.Sprintf("repos/%s/%s/commits", r.Owner, r.Name), token)
	if err != nil {
		return nil, fmt.Errorf("get commits for %s/%s: %v", r.Owner, r.Name, err)
	}
	fmt.Println(body)
	// FIXME: unmarshal JSON
	return []Commit{}, nil
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
