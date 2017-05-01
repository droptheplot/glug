package glug

import (
	"strings"
)

type node struct {
	path     string
	isParam  bool
	methods  map[string][]Plug
	children map[string]*node
}

func (curr *node) graft(method string, path string, plugs []Plug) {
	parts := strings.Split(path, "/")[1:]
	depth := len(parts)
	prev := curr

	for index, part := range parts {
		if child, ok := prev.children[part]; ok {
			curr = child
		} else {
			isParam := false

			if part[0] == ':' {
				isParam = true
			}

			curr = &node{
				path:     part,
				children: make(map[string]*node),
				isParam:  isParam,
			}
		}

		if depth == index+1 {
			curr.methods = make(map[string][]Plug)
			curr.methods[method] = plugs
		}

		prev.children[part] = curr
		prev = curr
	}
}
