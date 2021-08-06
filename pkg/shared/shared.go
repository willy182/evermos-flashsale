package shared

import (
	"math/rand"
	"strconv"
	"time"
)

// GenerateOrderID function for random number
func GenerateOrderID() string {
	rand.Seed(time.Now().UTC().UnixNano())

	randomString := rand.Int63n(time.Now().UnixNano())
	result := strconv.FormatInt(randomString, 10)
	return result
}
