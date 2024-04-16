package random

import "math/rand"
import "time"

func Get() int {
	// User must pass in number of integers to generate

	err := error(nil)
	if err != nil { // Maybe they didn't pass an integer
		panic(err)
	}
	start := time.Now()
	rand.Seed(start.UnixNano())
	return rand.Intn(256)
}
