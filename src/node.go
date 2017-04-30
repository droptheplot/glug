package glug

type node struct {
	path     string
	isParam  bool
	methods  map[string][]Plug
	children map[string]*node
}
