package yarl

import "time"

func (req *Request) Retry(attempts int, interval time.Duration) *Request {
	req.retry = attempts
	req.interval = interval
	return req
}

func (req *Request) Timeout(timeout time.Duration) *Request {
	req.client.Timeout = timeout
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
