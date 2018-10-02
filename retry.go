package yarl

import (
	"net/http"
	"time"
)

func (req *Request) Retry(attempts int, interval time.Duration) *Request {
	return req
}

func (req *Request) succeeded(c int) bool {
	for _, code := range req.successCode {
		if c == code {
			return true
		}
	}
	return false
}

func (req *Request) doWithRetry(r *http.Request) *Response {
	if req.retry <= 0 {
		req.retry = 1
	}

	var failed []*Response
	res := &Response{}
	defer func() { res.FailedAttempts = failed }()

	for i := 0; i < req.retry; i++ {
		raw, err := req.client.Do(r)
		resp := &Response{
			Raw: raw,
			err: err,
		}

		if err == nil && req.validator(resp) {
			res = resp
			return res
		}

		failed = append(failed, resp)

		if i == req.retry-1 {
			return res
		}

		select {
		case <-time.After(req.interval):
			continue
		case <-r.Context().Done():
			return res
		}
	}

	return res
}
