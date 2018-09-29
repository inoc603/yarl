package body

import "io"

type Body interface {
	// Type returns the body type
	Type() string

	io.ReadCloser
}
