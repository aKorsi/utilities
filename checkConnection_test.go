package utilities

import (
	"testing"
	"time"
)

func TestCheckConnection(t *testing.T) {
	err := CheckConnection("127.0.0.1", "80", time.Second)
	if err != nil {
		t.Error(err)
	}
}
