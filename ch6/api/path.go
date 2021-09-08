package main

import "strings"

const PathSeperator = "/"

type Path struct {
	Path string
	ID   string
}

func NewPath(p string) *Path {
	var id string
	p = strings.Trim(p, PathSeperator)
	s := strings.Split(p, PathSeperator)
	if len(s) > 1 {
		id = s[len(s)-1]
		p = strings.Join(s[:len(s)-1], PathSeperator)
	}
	return &Path{
		Path: p,
		ID:   id,
	}
}

func (p *Path) HasID() bool {
	return len(p.ID) > 0
}
