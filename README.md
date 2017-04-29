# Glug

[![GoDoc](https://godoc.org/github.com/droptheplot/glug?status.svg)](https://godoc.org/github.com/droptheplot/glug)
[![Go Report Card](https://goreportcard.com/badge/github.com/droptheplot/glug)](https://goreportcard.com/report/github.com/droptheplot/glug)

Inspired by [Plug](https://github.com/elixir-lang/plug) and [Rack](https://github.com/rack/rack) this package provides simple router which allows you to aggregate functions into pipelines for each endpoint. Pipelines are built with:

##### Conn (struct)

```go
type Conn struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params  url.Values
}
```

Contains all request related data through plugs in pipeline.

##### Plug (function)

```go
type Plug func(Conn) Conn
```

Should modify `Conn` or `Halt()` pipeline.

## Getting Started

```shell
go get -u github.com/droptheplot/glug
```

## Usage

```go
package main

import (
	"github.com/droptheplot/glug"
	"fmt"
	"net/http"
)

func Root(conn glug.Conn) glug.Conn {
	fmt.Fprintf(conn.Writer, "Everyone can access this page.")
	return conn
}

func BlogIndex(conn glug.Conn) glug.Conn {
	fmt.Fprintf(conn.Writer, "Nothing to see here!")
	return conn
}

func Auth(conn glug.Conn) glug.Conn {
	http.Redirect(conn.Writer, conn.Request, "/", http.StatusFound)
	return conn.Halt()
}

func main() {
	r := glug.New()
	r.HandleFunc("GET", "/", Root)
	r.HandleFunc("GET", "/blog", Auth, BlogIndex)
	http.ListenAndServe(":3000", r)
}
```
