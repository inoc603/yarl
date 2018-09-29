package yarl

import (
	"io"
	"net/http"
	"time"

	"github.com/inoc603/yarl/internal/body"
)

type Request struct {
	retry    int
	interval time.Duration

	client  http.Client
	req     *http.Request
	cookies []*http.Cookie

	body body.Body

	output io.Writer

	successCode []int
	err         error
}

func newReq(method, url string) *Request {
	r := &Request{
		successCode: []int{200},
	}

	r.req, r.err = http.NewRequest(method, url, r.body)
	return r
}

func (req *Request) DoStd() (*http.Response, error) {
	return req.client.Do(req.req)
}
