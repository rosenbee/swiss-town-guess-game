package towndata

import (
	"math/rand"
	"time"
)

func GetRandomSwissTownBFSNumber() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 770
	return rand.Intn(max-min+1) + min
}
