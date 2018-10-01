package yarl

import (
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/inoc603/yarl/internal/body"
)

func (req *Request) Set(k string, v interface{}) *Request {
	req.err = req.body.SetItem(k, v)
	return req
}

func (req *Request) Body(v interface{}) *Request {
	req.err = req.body.Set(v)
	return req
}

func (req *Request) JSON() *Request {
	req.body = body.NewJSON()
	return req
}

func (req *Request) Multipart() *Request {
	req.body = body.NewMultipart()
	return req
}

// File adds a file from the given path to multipart field, with optional custom
// filedname. Calling this will automatically sets the body type to multipart.
// It will be considered an error if the body is of another type and not empty.
func (req *Request) File(path string, field ...string) *Request {
	f, err := os.Open(path)
	if err != nil {
		req.err = errors.Wrapf(err, "open file %s", path)
		return req
	}

	fieldName := filepath.Base(path)
	if len(field) > 0 {
		fieldName = field[0]
	}

	return req.FileFromReader(f, path, fieldName)
}

// FileFromReader adds a file to the multipart from the given reader. Its behavior
// is the same as File.
// TODO: Maybe an io.ReadCloser?
func (req *Request) FileFromReader(r io.Reader, name string, field string) *Request {
	if _, ok := req.body.(*body.Multipart); !ok {
		if !req.body.IsEmpty() {
			req.err = errors.Errorf("there's already data in %s body", req.body.Type())
			return req
		}
	}

	if err := req.body.SetItem(field, &body.File{name, r}); err != nil {
		req.err = err
		return req
	}

	return req
}
