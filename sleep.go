package repeat

import (
	"math/rand"
	"time"
)

// sleepDuration returns a time-duration between 'from' and 'to'.
func sleepDuration(from time.Duration, to time.Duration, randomness *rand.Rand) time.Duration {

	if to < from {
		from, to = to, from
	}

	var difference int64 = to.Nanoseconds() - from.Nanoseconds()

	var delta int64 = randomness.Int63n(difference)

	var duration time.Duration = from + (time.Nanosecond * time.Duration(delta))

	return duration
}

// sleep sleeps for a time-duration randomnly chosen between 'from' and 'to'.
func sleep(from time.Duration, to time.Duration, randomness *rand.Rand) {

	var duration time.Duration = sleepDuration(from, to, randomness)

	time.Sleep(duration)
}

