package concur

import (
	"fmt"
	"time"
)

func undyingBoring(msg string) {
	for i := 0; ; i++ {
		fmt.Printf("%s %d\n", msg, i)
		time.Sleep(time.Second)
	}
}

func Boring() {
	go undyingBoring("boring!") //go and execute and I'm not going to wait for you.  just "go" do this
	fmt.Println("I'm listening.")
	time.Sleep(5 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}
