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
	max := 6810 // highest BFSNR of a swiss town on 2022-08-12 according to: https://swisspost.opendatasoft.com/explore/dataset/politische-gemeinden_v2/table/?sort=bfsnr
	return rand.Intn(max-min+1) + min
}
