package pkg

import (
	crand "crypto/rand"
	"math"
	"math/big"
	rnd "math/rand"

	"github.com/seehuhn/mt19937"
)

func GetRand() *rnd.Rand {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rng := rnd.New(mt19937.New())
	rng.Seed(seed.Int64())
	return rng
}
