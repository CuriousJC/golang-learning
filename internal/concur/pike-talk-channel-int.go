package concur

import (
	"fmt"
	"time"
)

func incrementInt(num int, c chan int) {
	for num < 20 {
		num++
		c <- num //send the number into the channel
		fmt.Printf("Sent number %d into the channel\n", num)
		time.Sleep(500 * time.Millisecond)
	}
}

func BoringCounter() {
	numLoop := 3
	c := make(chan int, 1) //make a channel that can hold integers.  buffer size of 3
	go incrementInt(1, c)  //go and execute and I'm not going to wait for you.  just "go" do this
	for i := 1; i < numLoop; i++ {
		fmt.Printf("Loop Counter is %d, Received counter value is: %d\n", i, <-c) //receive from the channel
		time.Sleep(1 * time.Second)
	}

	fmt.Println("You're boring; I'm leaving.")
}
