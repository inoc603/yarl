package yarl

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/inoc603/yarl/internal/body"
)

type Request struct {
	retry    int
	interval time.Duration

	client  http.Client
	req     *http.Request
	cookies []*http.Cookie

	query url.Values

	validator ResponseValidator

	body body.Body

	output io.Writer

	successCode []int
	err         error
}

func newReq(method, url string) *Request {
	r := &Request{
		successCode: []int{200},
	}

	r.req, r.err = http.NewRequest(method, url, nil)
	if r.req != nil {
		r.query = r.req.URL.Query()
	}
	return r
}

func (req *Request) hasError() bool {
	return req.req == nil || req.err != nil
}

// WithContext sets a context for the request. Context are valid through
// retries and redirects.
func (req *Request) WithContext(ctx context.Context) *Request {
	if req.hasError() {
		return req
	}

	req.req = req.req.WithContext(ctx)
	return req
}

// Timeout sets timeout on the request
func (req *Request) Timeout(t time.Duration) *Request {
	if req.hasError() {
		return req
	}

	req.client.Timeout = t
	return req
}

// ResponseValidator tells whether the given response is valid.
type ResponseValidator func(*Response) bool

// MaxCode tells the client the validate the request with the given max status
// code. If status code of a response is larger than it, the response is
// considered to be failed.
func (req *Request) MaxCode(code int) *Request {
	if req.hasError() {
		return req
	}

	req.validator = func(resp *Response) bool {
		return resp.StatusCode() < code
	}
	return req
}

// Validator sets a custom response validator
func (req *Request) Validator(v ResponseValidator) *Request {
	if req.hasError() {
		return req
	}

	req.validator = v

	return req
}

// Do makes the request and returns a reponse.
func (req *Request) Do() (*Response, error) {
	req.req.URL.RawQuery = req.query.Encode()
	raw, err := req.client.Do(req.req)
	resp := &Response{
		Raw: raw,
	}
	return resp, err
}

func (req *Request) do() (*Response, error) {
	req.req.URL.RawQuery = req.query.Encode()
	raw, err := req.client.Do(req.req)
	resp := &Response{
		Raw: raw,
	}
	return resp, err
}

// DoMarshal makes the request and marshal the response body to the given
// interface according to the response content type. If the body can't be
// marshalled, the body content can still be used from the response. If the
// response is considered failed, the body will not be marshalled.
func (req *Request) DoMarshal(v interface{}) (*Response, error) {
	resp, err := req.Do()
	if err != nil {
		return resp, err
	}

	return resp, resp.BodyMarshal(v)
}

// DoSilent makes the request and only returns the error. Whether a response is
// valid or not can be deermined by MaxCode. The response body will not be read.
func (req *Request) DoSilent() error {
	return nil
}
