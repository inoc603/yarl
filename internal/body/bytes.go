package body

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
)

type BodyBytes struct {
	buffer []byte
}

func (body *BodyBytes) Type() string {
	return ""
}

func (body *BodyBytes) IsEmpty() bool {
	return len(body.buffer) == 0
}

func (body *BodyBytes) Encode() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(body.buffer))
}

func (body *BodyBytes) SetItem(k string, v interface{}) error {
	return nil
}

func (body *BodyBytes) Set(v interface{}) error {
	switch b := v.(type) {
	case io.Reader:
		data, err := ioutil.ReadAll(b)
		if err == nil {
			body.buffer = data
		}
		return err
	case string:
		body.buffer = []byte(b)
		return nil
	case []byte:
		body.buffer = b
		return nil
	default:
		return errors.Errorf("unsupported value type")
	}
}
