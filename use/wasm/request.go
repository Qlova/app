package wasm

import (
	"bytes"
	"io"

	"qlova.org/seed/client/clientrpc"
)

//Request holds the metadata about an incomming request from the client.
type Request struct {
	path string

	response io.Writer

	headers map[string]string
}

func (cr Request) SetHeader(key, value string) {
	cr.headers[key] = value
}

func (cr Request) Header(key string) string {
	return cr.headers[key]
}

func (cr Request) Path() string {
	return cr.path
}

func (cr Request) Writer() io.Writer {
	return cr.response
}

//NewRequest returns a new request from the given values.
func NewRequest(buffer *bytes.Buffer) Request {
	return Request{
		path:     "/",
		response: buffer,
		headers:  make(map[string]string),
	}
}

//Arg returns the named query value with the given name.
func (cr Request) Arg(name string) string { return "" }

//Set sets the value of a cookie associated with requests by this client.
func (cr Request) Set(c clientrpc.Cookie, value string) {}

//Get gets the value of a cookie associated with requests by this client.
func (cr Request) Get(c clientrpc.Cookie) string { return "" }
