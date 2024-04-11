package hello

import "fmt"

import "rsc.io/quote"

import "math/rand"

import "time"

func Myhello() {
	fmt.Println(quote.Go())

	now := time.Now()
	rand.Seed(now.UnixNano())
	println("Numbers seeded using current date/time:", now.Format(time.StampNano))
	for i := 0; i < 5; i++ {
		println(rand.Intn(10))
	}

	fmt.Println("And now we're done...")

	fmt.Println("or maybe not...")

	fmt.Println("seriously...")

	fmt.Println("how are you just showing up?")

	fmt.Println("And now we're done...")
}

func Mygoodbye() string {
	return "Hello from a module!"
}
