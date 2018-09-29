package yarl

import (
	"bytes"
	"io"
	"net/http"
)

type Response struct {
	Raw       *http.Response
	bodyCache bytes.Buffer
}

func newResponse() *Response {
	return &Response{}
}

func (resp *Response) StatusCode() int {
	return resp.Raw.StatusCode
}

func (resp *Response) Body() io.WriteCloser {
	return nil
}

// ToStruct marshalls the body content to the given interface according to the
// content type.
func (resp *Response) ToStruct(v interface{}) error {
	return nil
}
