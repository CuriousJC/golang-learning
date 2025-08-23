package concur

import (
	"fmt"
	"time"
)

const waitChannelLoop = 10

type Message struct {
	str  string
	wait chan bool
}

func boringWait(msg string) <-chan Message {

	//create my generator channel
	c := make(chan Message)

	//create my waitForIt chanel
	waitForIt := make(chan bool) // shared by all messages from this generator

	go func() {
		for i := 0; ; i++ {

			if msg == "Bob" {
				time.Sleep(1000 * time.Millisecond) // Bob falls asleep twice
			}
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt} //send our message to the channel
			<-waitForIt
			// Wait until allowed to proceed which doesn't come in until the main loop
			//  has been unblocked for every message and then produces a "true" value to all the message waitForIts
		}
	}()
	return c
}

// fanIn takes ... channels as input and returns a single channel
// that is getting everything from multiple inputs and consolidating/multiplexing them
func fanInWait(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i]
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}

func WaitChannel() {

	//create a channel that has both my Bob message channel and my Sue message channel
	c := fanInWait(boringWait("Bob"), boringWait("Sue"))

	//Loop through printing my messages
	for i := 0; i < waitChannelLoop; i++ {
		msg1 := <-c //receive msg1 from my channel
		fmt.Println(msg1.str)

		msg2 := <-c //receive msg2 from my channel
		fmt.Println(msg2.str)

		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're boring; I'm leaving.")
}
