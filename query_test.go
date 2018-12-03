package yarl

import (
	"testing"

	"github.com/inoc603/yarl/internal/assert"
)

const exampleURL = "http://github.com/inoc603/yarl"

func TestQuery(t *testing.T) {
	a := assert.New(t)

	r := Get(exampleURL).Query("k", "v")
	a.Equal("v", r.url.Query().Get("k"))
}

func TestQueries(t *testing.T) {
	a := assert.New(t)

	a.Run("StructQueries", func(a *assert.A) {
		r := Get(exampleURL).Queries(struct {
			String string `query:"string"`
			Int    int    `query:"int"`
		}{"s", 1})

		q := r.url.Query()

		a.Equal("s", q.Get("string"))
		a.Equal("1", q.Get("int"))
	})

	a.Run("MapQueries", func(a *assert.A) {
		r := Get(exampleURL).Queries(map[string]interface{}{
			"string": "s",
			"int":    1,
		})

		q := r.url.Query()

		a.Equal("s", q.Get("string"))
		a.Equal("1", q.Get("int"))
	})
}
