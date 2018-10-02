package body

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"

	"github.com/buger/jsonparser"
)

type JSON struct {
	buffer []byte
}

func NewJSON() *JSON {
	return &JSON{}
}

func (body *JSON) Type() string {
	return "json"
}

func (body *JSON) Encode() io.ReadCloser {
	return &ReadCloser{bytes.NewBuffer(body.buffer)}
}

func (body *JSON) SetItem(k string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	if body.buffer == nil {
		body.buffer = make([]byte, 0)
	}

	_, err = jsonparser.Set(body.buffer, data, strings.Split(k, ".")...)
	return err
}

func (body *JSON) Set(value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if body.buffer == nil {
		body.buffer = data
		return nil
	}

	return nil
}

func (body *JSON) IsEmpty() bool {
	return len(body.buffer) == 0
}
