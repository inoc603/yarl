package body

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetItem(t *testing.T) {
	a := assert.New(t)

	j := NewJSON()

	a.Nil(j.SetItem("string", "a"))

	expected, _ := json.Marshal(struct {
		String string `json:"string"`
	}{"a"})

	a.Equal(expected, j.buffer)
}
