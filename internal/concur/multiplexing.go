package concur

import (
	"fmt"
	"math/rand"
	"time"
)

const multiplexLoop = 10

// Generator: function that returns a channel
func boringMultiplexBias1(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)

			//base delay
			delay := time.Duration(rand.Intn(1e3)) * time.Millisecond

			// bias: make Bob ~twice as slow
			if msg == "Bob" {
				delay *= 5
			}

			time.Sleep(delay)
		}
	}()
	return c
}

// Generator: function that returns a channel
func boringMultiplex(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {

			if msg == "Bob" {
				time.Sleep(1000 * time.Millisecond) // Bob falls asleep twice
			}
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(1000 * time.Millisecond)
		}
	}()
	return c
}

// fanIn takes two channels as input and returns a single channel
// that is getting everything from multiple inputs and consolidating/multiplexing them
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func GeneratorMultiplex() {

	//send in two channels to be multiplexed into a single channel that we will receive from
	c := fanIn(boringMultiplex("Bob"), boringMultiplex("Sue"))

	// Consume from the generated channel
	for i := 0; i < multiplexLoop; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}
