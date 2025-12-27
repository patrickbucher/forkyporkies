Given a list of GitHub repositories (`username/repository`), output a table of all users who forked at least one of those repositories and the number of non-originator commits on those forks.

Usage:

    GITHUB_TOKEN=... go run cmd/forkyporkies.go repos.txt

Example for `repos.txt`:

    patrickbucher/average-test-doubles
    patrickbucher/separation-of-concerns
    patrickbucher/connect-four-debugging
    patrickbucher/currency-converter

Example Output:
    
    Author                             average-test-doubles                separation-of-concerns                connect-four-debugging                    currency-converter
    randomdude37                                          6                                     1                                     1                                     2
    anotheruser                                           2                                     1                                     2                                     1
    foobarquxbaz                                          0                                     0                                     0                                     🐷
    johnny1337                                            1                                     1                                     2                                     3
    somebodyready                                         1                                     1                                     3                                     2
    Michi-Jackson                                         3                                     2                                     5                                     0
    Pingu1993                                            🐷                                     2                                     1                                     1

🐷 means: no fork for this repository