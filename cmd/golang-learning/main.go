package main

import "fmt"
import "github.com/curiousjc/golang-learning/internal/hello"
import "github.com/curiousjc/golang-learning/pkg/random"
import "github.com/curiousjc/golang-learning/pkg/goodbye"

func main() {
	fmt.Println("Hello!!!")

	fmt.Println(hello.FromAModule())

	hello.Myhello()

	hello.Mycompare(555555)

	fmt.Println("We have a random number coming from pkg and it is...")
	fmt.Println(random.Get())

	fmt.Println(goodbye.Goodbye())
}
