package yarl

import "net/http"

type RedirectPolicy func(req Request, via []*Request) error

func maxRedirect(max int) RedirectPolicy {
	return func(req Request, via []*Request) error {
		if len(via) >= max {
			return http.ErrUseLastResponse
		}

		return nil
	}
}

func (req *Request) MaxRedirect(max int) *Request {
	req.redirectPolicies = append(req.redirectPolicies, maxRedirect(max))
	return req
}
