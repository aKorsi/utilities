package utilities

import (
	"net/http"
)

func CheckInternetConnection(address string) error {
	if len(address) == 0 {
		address = "http://clients3.google.com/generate_204"
	}
	_, err := http.Get(address)
	if err != nil {
		return err
	}
	return nil
}
