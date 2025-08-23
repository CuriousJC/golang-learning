package concur

import (
	"fmt"
	"math/rand"
	"time"
)

const numLoop = 10

func channelBoring(msg string, c chan string) {
	for i := 0; ; i++ { //undying loop
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func BoringChannel() {
	c := make(chan string)
	go channelBoring("channel-text-is-boring!", c)
	for i := 0; i < numLoop; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	go channelBoring("channel-text-is-still-boring!", c)
	for i := 0; i < numLoop; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func BoringChannels() {
	c1 := make(chan string)
	go channelBoring("channel-text-is-boring!", c1)
	for i := 0; i < numLoop; i++ {
		fmt.Printf("You say: %q\n", <-c1)
	}

	c2 := make(chan string)
	go channelBoring("channel-text-is-still-boring!", c2)
	for i := 0; i < numLoop; i++ {
		fmt.Printf("You say: %q\n", <-c2)
	}
	fmt.Println("You're boring; I'm leaving.")
}

// channelBoringClean now listens for a quit signal
func channelBoringClean(msg string, c chan string, quit chan struct{}) {
	for i := 0; ; i++ {
		select {
		case c <- fmt.Sprintf("%s %d", msg, i):
			// sent successfully, sleep a bit
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		case <-quit:
			// received quit signal, stop the goroutine
			fmt.Printf("%s quitting!\n", msg)
			return
		}
	}
}

func BoringChannelsClean() {
	c1 := make(chan string)
	quit1 := make(chan struct{}) // channel to signal goroutine to stop

	// start first boring goroutine
	go channelBoringClean("channel-text-is-boring!", c1, quit1)
	for i := 0; i < numLoop; i++ {
		fmt.Printf("You say: %q\n", <-c1)
	}

	// signal first goroutine to stop
	close(quit1) // closing a channel broadcasts to all listeners

	// start second boring goroutine
	c2 := make(chan string)
	quit2 := make(chan struct{})
	go channelBoringClean("channel-text-is-still-boring!", c2, quit2)
	for i := 0; i < numLoop; i++ {
		fmt.Printf("You say: %q\n", <-c2)
	}

	// stop second goroutine
	close(quit2)

	fmt.Println("You're boring; I'm leaving.")
}

func BoringChannelsOneQuit() {
	c1 := make(chan string)
	quit := make(chan struct{}) // channel to signal goroutine to stop

	// start first boring goroutine
	go channelBoringClean("channel-text-is-boring!", c1, quit)
	for i := 0; i < numLoop; i++ {
		fmt.Printf("You say: %q\n", <-c1)
	}

	// start second boring goroutine
	c2 := make(chan string)
	go channelBoringClean("channel-text-is-still-boring!", c2, quit)
	for i := 0; i < numLoop; i++ {
		fmt.Printf("You say: %q\n", <-c2)
		fmt.Printf("You say: %q\n", <-c1)
	}

	close(quit) // closing a channel broadcasts to all listeners

	fmt.Println("You're boring; I'm leaving.")
}
