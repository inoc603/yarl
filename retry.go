package yarl

import (
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

func (req *Request) doWithRetry() *Response {
	if req.retry < 0 {
		req.retry = 1
	}

	var failed []*Response
	var res *Response
	defer func() { res.FailedAttempts = failed }()

	for i := 0; i < req.retry; i++ {
		resp, err := req.doRaw()
		if err == nil && req.validator(resp) {
			return res
		}

		failed = append(failed, resp)

		if i == req.retry-1 {
			return res
		}

		select {
		case <-time.After(req.interval):
			continue
		case <-req.req.Context().Done():
			return res
		}
	}

	return res
}
