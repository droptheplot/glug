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

// Halt will stop execution of plugs.
func (conn Conn) Halt() Conn {
	conn.halted = true
	return conn
}

// newConn returns new Conn struct.
func newConn(w http.ResponseWriter, r *http.Request, p url.Values) Conn {
	return Conn{Writer: w, Request: r, Params: p, halted: false}
}
