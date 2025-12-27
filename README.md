Given a list of GitHub repositories (`username/repository`), output a table of all users who forked at least one of those repositories and the number of non-originator commits on those forks.

API:

- [forks](https://docs.github.com/en/rest/repos/forks?apiVersion=2022-11-28)
- [commits](https://docs.github.com/en/rest/commits/commits?apiVersion=2022-11-28)


Usage:

    GITHUB_TOKEN=... go run cmd/forkyporkies.go repos.txt