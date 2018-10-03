package yarl

import (
	"net/http"
)

// RedirectPolicy represents policy for handling redirections. It works like
// http.Client.CheckRedirect
type RedirectPolicy func(req *http.Request, via []*http.Request) error

func maxRedirect(max int) RedirectPolicy {
	return func(req *http.Request, via []*http.Request) error {
		if len(via) > max {
			return http.ErrUseLastResponse
		}

		return nil
	}
}

// MaxRedirect sets the maxium redirects of the request
func (req *Request) MaxRedirect(max int) *Request {
	req.redirectPolicies = append(req.redirectPolicies, maxRedirect(max))
	return req
}

// RedirectPolicy adds a custom redirect policy to the request.
func (req *Request) RedirectPolicy(p RedirectPolicy) *Request {
	req.redirectPolicies = append(req.redirectPolicies, p)
	return req
}
