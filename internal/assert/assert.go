// package assert provides simple testig wrapper around testify/assert
package assert

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// A is a simple wrapper around assert.Assertions to provide easier sub testing
type A struct {
	*require.Assertions
	t *testing.T
}

func New(t *testing.T) *A {
	return &A{require.New(t), t}
}

func (a *A) Run(name string, f func(a *A)) {
	a.t.Run(name, func(t *testing.T) {
		f(New(t))
	})
}
