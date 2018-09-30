package body

import "io"

type Body interface {
	// Type returns the body type
	Type() string

	Encode() io.Reader

	SetItem(k string, v interface{}) error

	Set(v interface{}) error
}
