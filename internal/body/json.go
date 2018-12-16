package body

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/tidwall/sjson"
)

type JSON struct {
	buffer []byte
}

func NewJSON() *JSON {
	return &JSON{
		buffer: []byte{},
	}
}

func (body *JSON) Type() string {
	return "application/json"
}

func (body *JSON) Encode() io.ReadCloser {
	return &ReadCloser{bytes.NewBuffer(body.buffer)}
}

func (body *JSON) SetItem(k string, v interface{}) error {
	modified, err := sjson.SetBytesOptions(body.buffer, k, v, &sjson.Options{
		Optimistic:     true,
		ReplaceInPlace: true,
	})
	body.buffer = modified
	return err
}

func (body *JSON) Set(value interface{}) error {
	// TODO: add to current data
	switch v := value.(type) {
	case string:
		body.buffer = []byte(v)
	case []byte:
		// TODO: validate json?
		body.buffer = v
	case io.Reader:
		data, err := ioutil.ReadAll(v)
		if err != nil {
			return err
		}

		body.buffer = data
	default:
		data, err := json.Marshal(value)
		if err != nil {
			return err
		}

		body.buffer = data
	}

	return nil
}

func (body *JSON) IsEmpty() bool {
	return len(body.buffer) == 0
}
