package concur

import (
	"fmt"
	"math/rand"
	"time"
)

const generatorLoop = 10

// Generator: function that returns a channel
func boringGenerator(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func GeneratorBoring() {
	// Generator returns a channel
	bob := boringGenerator("boring!")

	// Consume from the generated channel
	for i := 0; i < numLoop; i++ {
		fmt.Printf("You say: %q\n", <-bob)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func GeneratorService() {
	// Generator returns a channel
	bob := boringGenerator("bob!")
	sue := boringGenerator("sue!")

	// Consume from the generated channel
	for i := 0; i < generatorLoop; i++ {
		fmt.Printf("You say: %q\n", <-bob)
		fmt.Printf("You say: %q\n", <-sue)
	}
	fmt.Println("You're boring; I'm leaving.")
}
