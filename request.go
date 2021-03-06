package yarl

import (
	"context"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/inoc603/yarl/internal/body"
	"github.com/pkg/errors"
	"golang.org/x/net/publicsuffix"
)

type Request struct {
	method  string
	url     *url.URL
	client  *http.Client
	cookies []*http.Cookie
	header  http.Header
	ctx     context.Context

	validator ResponseValidator

	body body.Body

	redirectPolicies []RedirectPolicy

	retryMax      int
	retryInterval time.Duration

	successCode []int
	err         error

	shared bool
}

func newReq(method string) *Request {
	jar, _ := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})

	r := &Request{
		method:      method,
		successCode: []int{200},
		header:      make(http.Header),
		client:      &http.Client{Jar: jar},
		validator:   func(_ *Response) bool { return true },
	}

	return r
}

func (req *Request) hasError() bool {
	return req.err != nil
}

// BasePath sets a common base path for all requests from this instance
func (req *Request) BasePath(p string) *Request {
	u := copyURL(req.url)
	u.Path = p
	req.url = u
	return req
}

// URL sets the request url.
func (req *Request) URL(rawURL string) *Request {
	u, err := url.Parse(rawURL)
	if err != nil {
		req.err = err
		return req
	}

	req.url = u
	return req
}

func (req *Request) Copy() *Request {
	cp := &Request{
		method:  req.method,
		url:     copyURL(req.url),
		cookies: req.cookies,
		header:  req.header,

		validator: req.validator,

		redirectPolicies: req.redirectPolicies,

		retryMax:      req.retryMax,
		retryInterval: req.retryInterval,
		successCode:   req.successCode,
	}

	if req.client != nil {
		// TODO: copy client
		cp.client = req.client
	}

	return cp
}

// Host sets a common host for all requests from this instance
func (req *Request) Host(h string) *Request {
	u := copyURL(req.url)
	u.Host = h
	req.url = u
	return req
}

// WithContext sets a context for the request. Context are valid through
// retries and redirects.
func (req *Request) WithContext(ctx context.Context) *Request {
	req.ctx = ctx
	return req
}

// Timeout sets timeout on the the http client
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
	req.validator = func(resp *Response) bool {
		if resp.StatusCode() > code {
			resp.err = errors.Errorf("status code %d", resp.StatusCode())
			return false
		}

		return true
	}
	return req
}

// Validator sets a custom response validator
func (req *Request) Validator(v ResponseValidator) *Request {
	req.validator = v
	return req
}

// Do makes the request and returns a reponse.
func (req *Request) Do() *Response {
	var body io.ReadCloser
	if req.body != nil {
		body = req.body.Encode()
	}

	r, err := http.NewRequest(req.method, req.url.String(), body)
	if err != nil {
		return &Response{err: errors.Wrap(err, "create request")}
	}

	for k := range req.header {
		r.Header.Add(k, req.header.Get(k))
	}

	for _, c := range req.cookies {
		r.AddCookie(c)
	}

	if req.ctx != nil {
		r = r.WithContext(req.ctx)
	}

	// Apply redirect policy if any. When no policy is given, the client
	// will use the default policy of http.Client, which stops at 10 redirections
	if len(req.redirectPolicies) > 0 {
		req.client.CheckRedirect = func(r *http.Request, via []*http.Request) error {
			for _, p := range req.redirectPolicies {
				if err := p(r, via); err != nil {
					return err
				}
			}
			return nil
		}
	}

	return req.doWithRetry(r)
}
