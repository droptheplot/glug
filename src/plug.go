package glug

// Plug is a function type we should use to make a pipeline for request.
type Plug func(Conn) Conn
