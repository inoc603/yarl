package yarl

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime"
	"net/http"

	"github.com/inoc603/yarl/internal/pipe"
	"github.com/pkg/errors"
)

type Response struct {
	// Raw is the underlying http.Response
	Raw *http.Response

	err error

	// FailedAttempts keeps all previous failed attempts through retries
	FailedAttempts []*Response

	// bodyCache caches the response body for latter use
	bodyCache []byte
}

func newResponse(raw *http.Response) *Response {
	return &Response{
		Raw: raw,
	}
}

func (resp *Response) Error() error {
	return resp.err
}

func (resp *Response) StatusCode() int {
	if resp.Raw == nil {
		return 0
	}
	return resp.Raw.StatusCode
}

func (resp *Response) contentType() string {
	if resp.Raw == nil {
		return ""
	}
	return resp.Raw.Header.Get("Content-Type")
}

// Body returns the body as io.Reader
func (resp *Response) Body() io.Reader {
	if resp.Raw == nil {
		// TODO: Maybe we should not return nil reader here to avoid
		// possible panic?
		return nil
	}

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
	t, _, err := mime.ParseMediaType(resp.contentType())
	if err != nil {
		return errors.Wrap(err, "parse content-type")
	}

	switch t {
	case "application/json":
		return json.NewDecoder(resp.Body()).Decode(v)
	}
	return errors.Errorf("unknown body type %s", resp.contentType())
}

// BodyJSON marshalls the body content to the given interface as JSON, regardless
// of the Content-Type header in the response.
func (resp *Response) BodyJSON(v interface{}) error {
	return json.NewDecoder(resp.Body()).Decode(v)
}
