package main

import (
	"container/ring"
)

type Snippet struct {
	URL  string
	Body string

	Error         error
	Output        string
	IsValidPython bool
}

var Ring = ring.New(10)

func SaveSnippet(s Snippet) {
	Ring.Value = &s
	Ring = Ring.Next()
}

func LoadSnippets() []*Snippet {
	out := []*Snippet{}
	Ring.Do(func(p interface{}) {
		if p != nil {
			out = append(out, p.(*Snippet))
		}
	})

	return out
}
