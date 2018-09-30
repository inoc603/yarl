package yarl

import (
	"time"
)

func (req *Request) Retry(attempts int, interval time.Duration) *Request {
	req.retry = attempts
	req.interval = interval
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

type ResponseError struct {
	Response *Response
	Err      error
}

func (e *ResponseError) Error() string {
	return e.Err.Error()
}

type Executor interface {
	Do(*Response) []*ResponseError
}

type RetryExecutor struct {
	Attempt  int
	Interval time.Duration
	// Context       context.Context
}

func (exe *RetryExecutor) Do(req *Request) []*ResponseError {
	if exe.Attempt < 0 {
		exe.Attempt = 1
	}

	var res []*ResponseError

	for i := 0; i < exe.Attempt; i++ {
		resp, err := req.do()
		res = append(res, &ResponseError{resp, err})
		if err == nil && req.validator(resp) {
			return res
		}

		select {
		case <-time.After(exe.Interval):
			continue
		case <-req.req.Context().Done():
			return res
		}
	}

	return res
}
