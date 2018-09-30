package yarl

import "net/http"

func (req *Request) Cookie(c *http.Cookie) *Request {
	return req
}
