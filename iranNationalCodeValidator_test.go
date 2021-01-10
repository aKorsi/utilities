package utilities

import (
	"testing"
)

func TestIranNationalCodeValidator(t *testing.T) {
	if IranNationalCodeValidator("0066829028") == false {
		t.Error("not pass!")
	}
}
