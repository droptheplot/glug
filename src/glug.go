package glug

import (
	// "github.com/davecgh/go-spew/spew"
	"net/http"
	"net/url"
	"strings"
)

// Conn is a struct to carry all request related data through plugs.
type Conn struct {
	uuid    string
	halted  bool
	Writer  http.ResponseWriter
	Request *http.Request
	Params  url.Values
}

// Router is a http.Handler.
type Router struct {
	node node
}

// Plug is a function type we should use to make a pipeline for request.
type Plug func(Conn) Conn

type node struct {
	path     string
	methods  map[string][]Plug
	children map[string]*node
}

// Halt will stop execution of plugs.
func (conn Conn) Halt() Conn {
	conn.halted = true
	return conn
}

// newConn returns new Conn struct.
func newConn(w http.ResponseWriter, r *http.Request, p url.Values) Conn {
	return Conn{Writer: w, Request: r, Params: p, halted: false}
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
		var curr *node

		parts := strings.Split(path, "/")[1:]
		depth := len(parts)
		prev := &router.node

		for index, part := range parts {
			if child, ok := prev.children[part]; ok {
				curr = child
			} else {
				curr = &node{path: part, children: make(map[string]*node)}
			}

			if depth == index+1 {
				curr.methods = make(map[string][]Plug)
				curr.methods[method] = plugs
			}

			prev.children[part] = curr
			prev = curr
		}
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
			curr = curr.children[part]

			if depth == index+1 {
				break
			}
		}
	}

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
