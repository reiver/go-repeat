package repeat

import (
	"math/rand"
	"time"
)

func DelayFunc(from time.Duration, to time.Duration, fn func()bool) func()bool {
	return func() bool {
		return Delay(from, to, fn)
	}
}

// Delay waits between 'from' and 'to' and then calls 'fn'.
func Delay(from time.Duration, to time.Duration, fn func()bool) bool {
	randomness := newrandomness()
	return delay(from, to, randomness, fn)
}

func delay(from time.Duration, to time.Duration, randomness *rand.Rand, fn func()bool) bool {
	if nil == randomness {
		panic(errNilRandomness)
	}
	if nil == fn {
		panic(errNilFunc)
	}

	sleep(from, to, randomness)
	return fn()
}
