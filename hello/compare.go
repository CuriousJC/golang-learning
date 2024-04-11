package main

import "fmt"
import "os"
import "strconv"
import "math/rand"
import crand "crypto/rand"
import "time"

func main() {
	// User must pass in number of integers to generate
	if len(os.Args) < 2 {
		println("Usage:\n")
		println("  go run compare.go <numberOfInts>")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil { // Maybe they didn't pass an integer
		panic(err)
	}

	// Phase 1 — Using math/rand
	// Initialize the byte slice
	b1 := make([]byte, n)
	// Get the time
	start := time.Now()
	//time.Sleep(5)
	// Initialize the random number generator
	rand.Seed(start.UnixNano())
	// Generate the pseudo-random numbers
	for i := 0; i < n; i++ {
		b1[i] = byte(rand.Intn(256)) // Where the magic happens!
	}
	// Compute the elapsed time
	elapsed := time.Now().Sub(start)
	// In case n is very large, only print a few numbers
	for i := 0; i < 5; i++ {
		println(b1[i])
	}
	fmt.Printf("Time to generate %v pseudo-random numbers with math/rand: %v\n", n, elapsed)

	// Phase 2 — Using crypto/rand
	// Initialize the byte slice
	b2 := make([]byte, n)
	// Get the time (Note: no need to seed the random number generator)
	start = time.Now()
	//time.Sleep(5)
	// Generate the pseudo-random numbers
	_, err = crand.Read(b2) // Where the magic happens!
	// Compute the elapsed time
	elapsed = time.Now().Sub(start)
	// exit if error
	if err != nil {
		panic(err)
	}
	// In case n is very large, only print a few numbers
	for i := 0; i < 5; i++ {
		println(b2[i])
	}
	fmt.Printf("Time to generate %v pseudo-random numbers with crypto/rand: %v\n", n, elapsed)
}
