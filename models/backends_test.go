package models

import (
	"fmt"
	"testing"
)

func TestControllers(t *testing.T) {
	err := DefaultBackend.UpdateKV("/skydns/sdp/examples/n8", "test", 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	//DefaultBackend.DeleteKey("/skydns/sdp/examples/n8")
}
