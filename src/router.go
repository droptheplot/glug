package glug

import (
	// "github.com/davecgh/go-spew/spew"
	"net/http"
	"strings"
)

// Router is a http.Handler.
type Router struct {
	node node
}

// New will initialize new router.
func New() *Router {
	return &Router{node: node{children: make(map[string]*node)}}
}

// HandleFunc will add new endpoint to router.
func (router *Router) HandleFunc(method string, path string, plugs ...Plug) {
	if path == "/" {
		router.node.path = path
		router.node.methods = make(map[string][]Plug)
		router.node.methods[method] = plugs
	} else {
		router.node.graft(method, path, plugs)
	}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}

	logRequest(r.Method, r.URL.Path)

	r.ParseForm()

	curr := &router.node
	conn := newConn(w, r, r.Form)

	if r.URL.Path != "/" {
		parts := strings.Split(r.URL.Path, "/")[1:]
		depth := len(parts)

		for index, part := range parts {
			if curr.children[part] == nil {
				for _, child := range curr.children {
					if child.isParam == true {
						conn.Params.Add(child.path[1:], part)
						curr = child
						break
					}
				}
			} else {
				curr = curr.children[part]
			}

			if depth == index+1 {
				break
			}
		}
	}

	logParams(conn.Params)

	for _, plug := range curr.methods[r.Method] {
		result := plug(conn)

		if result.halted == false {
			logPlug(plug)
			conn = result
		} else {
			logHalt(plug)
			break
		}
	}
}
