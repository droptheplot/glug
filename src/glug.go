package glug

import (
	"net/http"
	"net/url"
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
	children map[string][]Plug
}

// Plug is a function type we should use to make a pipeline for request.
type Plug func(Conn) Conn

// Halt will stop execution of plugs.
func (conn Conn) Halt() Conn {
	conn.halted = true
	return conn
}

// Init will initialize new router.
func Init() *Router {
	return &Router{children: make(map[string][]Plug)}
}

// HandleFunc will add new endpoint to router.
func (router *Router) HandleFunc(method string, path string, plugs ...Plug) {
	router.children[path] = append(router.children[path], plugs...)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/favicon.ico" {
		logRequest("GET", r.URL.Path)
	}

	r.ParseForm()

	conn := Conn{Writer: w, Request: r, Params: r.Form, halted: false}
	plugs := router.children[r.URL.Path]

	for _, plug := range plugs {
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
