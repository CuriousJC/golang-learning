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
	go undyingBoring("boring!")
	fmt.Println("I'm listening.")
	time.Sleep(5 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}
