package yarl

import (
	"testing"

	"github.com/inoc603/yarl/internal/assert"
)

func TestHeader(t *testing.T) {
	a := assert.New(t)

	a.Run("SingleHeader", func(a *assert.A) {
		r := Get(exampleURL).
			Header("h1", "v1").
			Header("h2", "v2")

		a.Equal("v1", r.header.Get("h1"))
		a.Equal("v2", r.header.Get("h2"))
	})

	a.Run("StructHeader", func(a *assert.A) {
		r := Get(exampleURL).Headers(struct {
			String string `header:"string"`
			Int    int    `header:"int"`
		}{"s", 1})

		h := r.header

		a.Equal("s", h.Get("string"))
		a.Equal("1", h.Get("int"))
	})

	a.Run("MapHeader", func(a *assert.A) {
		r := Get(exampleURL).Headers(map[string]interface{}{
			"string": "s",
			"int":    1,
		})

		h := r.header

		a.Equal("s", h.Get("string"))
		a.Equal("1", h.Get("int"))
	})
}
