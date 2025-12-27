Given a list of GitHub repositories (`username/repository`), output a table of all users who forked at least one of those repositories and the number of non-originator commits on those forks.

Usage:

    GITHUB_TOKEN=... go run cmd/forkyporkies.go repos.txt

Example for `repos.txt`:

    patrickbucher/average-test-doubles
    patrickbucher/separation-of-concerns
    patrickbucher/connect-four-debugging
    patrickbucher/currency-converter

Output:
    
    Author                             average-test-doubles                separation-of-concerns                connect-four-debugging                    currency-converter
    pascalbossert                                         6                                     1                                     1                                     2
    ELDylanosaurus                                        2                                     1                                     2                                     1
    jbokito                                               0                                     0                                     0                                     🐷
    marvinbbzw                                            1                                     1                                     2                                     3
    marvinhegi                                            1                                     1                                     3                                     2
    Fabian-Scheuch                                        3                                     2                                     5                                     0
    TobiasClausen                                         🐷                                     2                                     1                                     1
    Siri-On-Off                                           3                                     3                                     6                                     0
    yannickict                                            1                                     1                                     1                                     1
    HeiligerG                                             2                                     1                                     3                                    11
    Marcoo007                                             0                                     1                                     3                                     0
    TheGamingSimon                                        0                                     1                                     1                                     🐷
    StevenMattmann                                        0                                     1                                     1                                     0
    Domenic77                                             3                                     0                                     1                                     0
    LennyAmstutz                                          1                                     2                                     2                                     0
    Paedimeier                                            🐷                                     0                                     🐷                                     🐷
    Mich-i                                                1                                     2                                     2                                     2
    SaskiaTe                                              1                                     1                                     3                                     1
    arturict                                              2                                     9                                     1                                     5
    simonMettler2000                                      1                                     1                                     1                                     2
    Ricci008                                              1                                     1                                     1                                     0
    NicolasRapedius                                       2                                     1                                     0                                     0
    BBZWRodriguez                                         🐷                                     0                                     1                                     0
    DanialNajafi                                          1                                     1                                     2                                     0
    YvesTroxler1                                          0                                     0                                     0                                     0
    Nico610                                               1                                     0                                     0                                     🐷
    NoahLangensee-GIT                                     4                                     1                                     1                                     0
    LEEEAAAAAA                                            0                                     🐷                                     0                                     0
    utigernils                                            1                                     🐷                                     4                                     4
    nandobrun                                             0                                     1                                     1                                     0
