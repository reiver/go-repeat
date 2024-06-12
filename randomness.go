package repeat

import (
	"math/rand"
	"time"
)

func newrandomness() *rand.Rand {
	return rand.New( rand.NewSource(time.Now().Unix() + rand.Int63n(0xFFFFFFFF) ) )
}
