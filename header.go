package yarl

import (
	"net/http"

	"github.com/pkg/errors"
)

func (req *Request) Header(k, v string) *Request {
	req.header.Add(k, v)
	return req
}

func (req *Request) Headers(v interface{}) *Request {
	switch headers := v.(type) {
	case http.Header:
		req.header = headers
	default:
		req.err = errors.Errorf("unsupported headers type")
	}

	return req
}

func (req *Request) Cookie(c *http.Cookie) *Request {
	req.cookies = append(req.cookies, c)
	return req
}
