package yarl

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/inoc603/yarl/internal/pipe"
	"github.com/pkg/errors"
)

type Response struct {
	// Raw is the underlying http.Response
	Raw *http.Response

	// bodyCache caches the response body for latter use
	bodyCache []byte
}

func newResponse(raw *http.Response) *Response {
	return &Response{
		Raw: raw,
	}
}

func (resp *Response) StatusCode() int {
	return resp.Raw.StatusCode
}

func (resp *Response) contentType() string {
	return resp.Raw.Header.Get("Content-Type")
}

// Body returns the body as io.Reader
func (resp *Response) Body() io.Reader {
	if resp.bodyCache != nil {
		return bytes.NewBuffer(resp.bodyCache)
	}

	p := pipe.New(func(b []byte) {
		resp.bodyCache = append(resp.bodyCache, b...)
	})

	go func() {
		_, err := io.Copy(p.Writer(), resp.Raw.Body)
		p.CloseRead(err)
		resp.Raw.Body.Close()
	}()

	return p.Reader()
}

// BodyBytes returns the body as bytes
func (resp *Response) BodyBytes() ([]byte, error) {
	return ioutil.ReadAll(resp.Body())
}

// BodyString returns the body as string
func (resp *Response) BodyString() (string, error) {
	b, err := resp.BodyBytes()
	return string(b), err
}

// BodyMarshal marshalls the body content to the given interface according to the
// content type.
func (resp *Response) BodyMarshal(v interface{}) error {
	if strings.Contains(resp.contentType(), "application/json") {
		return json.NewDecoder(resp.Body()).Decode(v)
	}
	return errors.Errorf("unknown body type %s", resp.contentType())
}

// BodyJSON marshalls the body content to the given interface as JSON, regardless
// of the Content-Type header in the response.
func (resp *Response) BodyJSON(v interface{}) error {
	return json.NewDecoder(resp.Body()).Decode(v)
}
