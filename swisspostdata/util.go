package towndata

import (
	"math/rand"
	"time"
)

// Get random number between 1 and 770
// BFS numbers for Swiss towns are in that range
func GetRandomSwissTownBFSNumber() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 770
	return rand.Intn(max-min+1) + min
}
