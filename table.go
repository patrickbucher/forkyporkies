package forkyporkies

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Table struct {
	Repos []string
	Forks map[string]Entry
}

type Entry struct {
	Author  string
	Commits map[string]uint
}

func (t Table) Output() {
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	format := "%s"
	for i := 0; i < len(t.Repos); i++ {
		format += "\t%s"
	}
	format += "\n"
	th := make([]string, 1)
	th[0] = "Author"
	for _, r := range t.Repos {
		th = append(th, r)
	}
	fmt.Fprintf(tw, format, th)
	for _, e := range t.Forks {
		commits := make([]uint, 0)
		for _, r := range t.Repos {
			if c, ok := e.Commits[r]; ok {
				commits = append(commits, c)
			} else {
				commits = append(commits, 0)
			}
		}
		tr := make([]string, 1)
		tr[0] = e.Author
		for _, c := range commits {
			tr = append(tr, fmt.Sprintf("%d", c))
		}
		fmt.Fprintf(tw, format, tr)
	}
	tw.Flush()
}
