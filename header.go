package yarl

import (
	"net/http"

	"github.com/pkg/errors"
)

func (req *Request) Header(k, v string) *Request {
	if req.req == nil || req.err != nil {
		return req
	}

	req.req.Header.Set(k, v)
	return req
}

func (req *Request) Headers(v interface{}) *Request {
	if req.req == nil || req.err != nil {
		return req
	}

	switch headers := v.(type) {
	case http.Header:
		req.req.Header = headers
	default:
		req.err = errors.Errorf("unsupported headers type")
	}

	return req
}
