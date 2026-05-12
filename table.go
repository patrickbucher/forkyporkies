package forkyporkies

import (
	"cmp"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
	"text/tabwriter"
)

type Table struct {
	Repos []string
	Forks map[string]Entry
}

type Entry struct {
	Author  string
	Commits map[string]int
}

func (t Table) Output() {
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	l := longest(t.Repos)
	format := "%s" + strings.Repeat(fmt.Sprintf("\t%%%ds", l), len(t.Repos)) + "\n"
	th := make([]string, 1)
	th[0] = "Author"
	for _, r := range t.Repos {
		th = append(th, strings.Split(r, "/")[1])
	}
	fmt.Fprintf(tw, format, toAny(th)...)
	entries := slices.SortedFunc(maps.Values(t.Forks), func(l, r Entry) int {
		return cmp.Compare(strings.ToLower(l.Author), strings.ToLower(r.Author))
	})
	for _, e := range entries {
		commits := make([]int, 0)
		for _, r := range t.Repos {
			if c, ok := e.Commits[r]; ok {
				commits = append(commits, c)
			} else {
				commits = append(commits, -1)
			}
		}
		tr := make([]string, 1)
		tr[0] = e.Author
		for _, c := range commits {
			s := fmt.Sprintf("%d", c)
			if c == -1 {
				s = "X"
			}
			tr = append(tr, s)
		}
		fmt.Fprintf(tw, format, toAny(tr)...)
	}
	tw.Flush()
}

func longest(xs []string) int {
	l := 0
	for _, x := range xs {
		if len(x) > l {
			l = len(x)
		}
	}
	return l
}

func toAny(xs []string) []any {
	ys := make([]any, len(xs))
	for i := 0; i < len(xs); i++ {
		ys[i] = xs[i]
	}
	return ys
}
