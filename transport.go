package yarl

import (
	"net/http"
	"net/url"

	"golang.org/x/net/proxy"

	"github.com/pkg/errors"
)

// Transport sets a custom transport for the http client. If the transport is
// not a *http.Transport, no further customization can be done to the transport
func (req *Request) Transport(r http.RoundTripper) *Request {
	if req.hasError() {
		return req
	}

	req.client.Transport = r
	return req
}

// Proxy sets the proxy used to perform the request. It supports http/https, and
// socks5 proxy.
func (req *Request) Proxy(proxyURL string) *Request {
	if req.hasError() {
		return req
	}

	transport, ok := req.client.Transport.(*http.Transport)
	if !ok {
		req.err = errors.Errorf("can't set proxy for custom transport")
	}

	u, err := url.Parse(proxyURL)
	if err != nil {
		req.err = errors.Wrap(err, "parse proxy url")
		return req
	}

	switch u.Scheme {
	case "http":
		transport.Proxy = http.ProxyURL(u)
	case "https":
		transport.Proxy = http.ProxyURL(u)
	case "socks5":
		dialer, err := proxy.SOCKS5("tcp", u.Host, nil, proxy.Direct)
		if err != nil {
			req.err = errors.Wrap(err, "create socks5 proxy")
			return req
		}

		transport.Dial = dialer.Dial
	default:
		req.err = errors.Errorf("unsupported proxy scheme %s", u.Scheme)
	}

	return req
}
