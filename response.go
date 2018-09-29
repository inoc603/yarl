package yarl

import "net/http"

type Response struct {
	Raw *http.Response
}

func (resp *Response) StatusCode() int {
	return resp.Raw.StatusCode
}
