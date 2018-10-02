package yarl

import (
	"fmt"
	"testing"
)

func TestTmp(t *testing.T) {
	resp := Get("http://whatever/v1.24/containers/json").
		UnixSocket("/var/run/docker.sock").
		Do()

	if resp.Error() == nil {
		fmt.Println(resp.BodyString())
	} else {
		fmt.Println(resp.Error())
	}
}

func TestXX(t *testing.T) {
}
