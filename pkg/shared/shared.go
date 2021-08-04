package shared

import (
	"math/rand"
	"time"

	"github.com/willy182/evermos-flashsale/pkg/helper"
)

// GenerateOrderID function for random number
func GenerateOrderID(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	charsLength := len(helper.NUMBERS)
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = helper.NUMBERS[rand.Intn(charsLength)]
	}
	return string(result)
}
