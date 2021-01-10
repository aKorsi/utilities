package utilities

import (
	"fmt"
	"math/rand"
	"time"
)

func RandCode(low, high, length int) string {
	rand.Seed(time.Now().UnixNano())
	code := ""
	for i := 0; i < length; i++ {
		code += fmt.Sprint(low + rand.Intn(high-low))
	}

	return code
}
