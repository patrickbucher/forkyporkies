Given a list of GitHub repositories (`username/repository`), output a table of all users who forked at least one of those repositories and the number of non-originator commits on those forks.

Usage:

    GITHUB_TOKEN=... go run cmd/forkyporkies.go repos.txt

Example for `repos.txt`:

    patrickbucher/average-test-doubles
    patrickbucher/separation-of-concerns
    patrickbucher/connect-four-debugging
    patrickbucher/currency-converter
    patrickbucher/performance-tests