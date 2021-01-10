package utilities

import "testing"

func TestCheckInternetConnection(t *testing.T) {
	err := CheckInternetConnection("")
	if err != nil {
		t.Error(err)
	}
}
