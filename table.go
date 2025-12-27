package forkyporkies

type Table struct {
	Repos []string
	Forks map[string]Entry
}

type Entry struct {
	Author  string
	Commits map[string]uint
}
