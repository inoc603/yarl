package yarl

import (
	"net/http"
	"time"
)

func (req *Request) Retry(attempts int, interval time.Duration) *Request {
	req.retryMax = attempts
	req.retryInterval = interval
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
	if req.retryMax <= 0 {
		req.retryMax = 1
	}

	var failed []*Response
	res := &Response{}
	defer func() { res.FailedAttempts = failed }()

	for i := 0; i < req.retryMax; i++ {
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

		if i == req.retryMax-1 {
			res = resp
			return res
		}

		select {
		case <-time.After(req.retryInterval):
			continue
		case <-r.Context().Done():
			return res
		}
	}

	return res
}
