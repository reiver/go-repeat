package repeat

import (
	"math/rand"
	"time"
)

// Forever call 'fn' over and over and over again forever, but waits between 'from' and 'to' before calling 'fn' again.
func Forever(from time.Duration, to time.Duration, fn func()) {
	randomness := newrandomness()
	forever(from, to, randomness, fn)
}

func forever(from time.Duration, to time.Duration, randomness *rand.Rand, fn func()) {
	if nil == randomness {
		panic(errNilRandomness)
	}
	if nil == fn {
		panic(errNilFunc)
	}

	for {
		fn()
		sleep(from, to, randomness)
	}
}
