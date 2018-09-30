package yarl

import (
	"fmt"
	"testing"
)

func TestTmp(t *testing.T) {
	resp, err := Get("http://whatever/v1.24/containers/json").
		UnixSocket("/var/run/docker.sock").
		Do()

	if err == nil {
		fmt.Println(resp.BodyString())
	} else {
		fmt.Println(err)
	}
}

func TestXX(t *testing.T) {
}
