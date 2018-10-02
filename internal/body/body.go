package body

import "io"

type Body interface {
	// Type returns the body type
	Type() string

	IsEmpty() bool

	Encode() io.ReadCloser

	SetItem(k string, v interface{}) error

	Set(v interface{}) error
}

type ReadCloser struct {
	io.Reader
}

func (r *ReadCloser) Close() error {
	return nil
}
