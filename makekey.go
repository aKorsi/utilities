package utilities

import (
	"math/rand"
	"time"
)

func MakeKey(key *[]byte) error {
	rand.Seed(time.Now().UnixNano())
	_, err := rand.Read(*key)
	return err
}
