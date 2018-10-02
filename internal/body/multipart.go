package body

import (
	"bytes"
	"io"
	"mime/multipart"

	"github.com/pkg/errors"
)

type Multipart struct {
	written bool
	buffer  []byte
	writer  *multipart.Writer
}

type File struct {
	Name    string
	Content io.Reader
}

func NewMultipart() *Multipart {
	body := &Multipart{}
	body.writer = multipart.NewWriter(bytes.NewBuffer(body.buffer))
	return body
}

func (body *Multipart) Type() string {
	return "multipart"
}

func (body *Multipart) Encode() io.ReadCloser {
	return &ReadCloser{bytes.NewBuffer(body.buffer)}
}

func (body *Multipart) SetItem(k string, v interface{}) error {
	f, ok := v.(*File)
	if !ok {
		return errors.Errorf("invalid type for multipart file")
	}

	body.written = true

	w, err := body.writer.CreateFormField(k)
	if err != nil {
		return errors.Wrap(err, "create from field")
	}

	_, err = io.Copy(w, f.Content)
	return errors.Wrap(err, "copy content")
}

func (body *Multipart) Set(v interface{}) error {
	return errors.Errorf("Set is not implemented on multipart body")
}

func (body *Multipart) IsEmpty() bool {
	return !body.written
}
