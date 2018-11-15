package yarl

import (
	"fmt"
	"net/url"
)

const (
	POST    = "POST"
	GET     = "GET"
	HEAD    = "HEAD"
	PUT     = "PUT"
	DELETE  = "DELETE"
	PATCH   = "PATCH"
	OPTIONS = "OPTIONS"
)

func copyURL(u *url.URL) *url.URL {
	return &url.URL{
		Scheme:     u.Scheme,
		Opaque:     u.Opaque,
		User:       u.User,
		Host:       u.Host,
		Path:       u.Path,
		RawPath:    u.RawPath,
		ForceQuery: u.ForceQuery,
		RawQuery:   u.RawQuery,
		Fragment:   u.Fragment,
	}
}

func New(url string) *Request {
	return Get(url)
}

func Get(url string) *Request {
	return newReq(GET).URL(url)
}

func Put(url string) *Request {
	return newReq(PUT).URL(url)
}

func Post(url string) *Request {
	return newReq(POST).URL(url)
}

func Delete(url string) *Request {
	return newReq(DELETE).URL(url)
}

func Patch(url string) *Request {
	return newReq(PATCH).URL(url)
}

func (req *Request) setURL(path string, args ...interface{}) *Request {
	u, err := url.Parse(fmt.Sprintf(path, args...))
	if err != nil {
		req.err = err
		return req
	}

	if u.Scheme == "http" || u.Scheme == "https" {
		req.url = u
		return req
	}

	req.url.Path += u.Path

	return req
}

func (req *Request) Get(url string, args ...interface{}) *Request {
	req.method = GET
	return req.setURL(url, args...)
}

func (req *Request) Put(url string, args ...interface{}) *Request {
	req.method = PUT
	return req.setURL(url, args...)
}

func (req *Request) Post(url string, args ...interface{}) *Request {
	req.method = POST
	return req.setURL(url, args...)
}

func (req *Request) Delete(url string, args ...interface{}) *Request {
	req.method = DELETE
	return req.setURL(url, args...)
}

func (req *Request) Patch(url string, args ...interface{}) *Request {
	req.method = PATCH
	return req.setURL(url, args...)
}

func (req *Request) Method(method string) *Request {
	req.method = method
	return req
}
