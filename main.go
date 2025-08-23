package main

//cd c:\repos\golang-learning\
//go run cmd\golang-learning\main.go

import "fmt"
import "github.com/curiousjc/golang-learning/internal/hello"
import "github.com/curiousjc/golang-learning/pkg/random"
import "github.com/curiousjc/golang-learning/pkg/goodbye"
import "github.com/curiousjc/golang-learning/pkg/conversation"
import "github.com/curiousjc/golang-learning/pkg/fizzbuzz"
import "github.com/curiousjc/golang-learning/pkg/runestuff"
import "github.com/curiousjc/golang-learning/internal/recurse"
import "github.com/curiousjc/golang-learning/internal/concur"

func main() {
	quiet := true
	fmt.Printf("Hello, this is just throwaway stuff for playing around.  I am running in a quiet mode of: %t\n", quiet)

	concur.WaitChannel()

	if !quiet {

		fibonacci := recurse.Fibonacci(10)
		fmt.Println("Fibonacci result is: ", fibonacci)

		factorial := recurse.Factorial(5)
		fmt.Println("Factorial Result is: ", factorial)

		fmt.Println(hello.FromAModule())

		hello.Myhello()

		hello.Mycompare(555555)

		fmt.Println("We have a random number coming from pkg and it is...")
		fmt.Println(random.Get())

		fmt.Println(goodbye.Goodbye())

		fmt.Println(conversation.Get())

		fizzbuzz.Fizzbuzz()

		runestuff.Get()

	}

}
