package fizzbuzz

import (
	"fmt"
)

func Fizzbuzz() {

	fmt.Println("Fizzing and Buzzing!!!")

	var three int = 3
	var five int = 5

	for i := 1; i < 101; i++ {
		if isMultiple(i, three) && isMultiple(i, five) {
			fmt.Println(i, "FizzBuzz!")
			continue
		}
		if isMultiple(i, three) {
			fmt.Println(i, "Fizz!")
		}
		if isMultiple(i, five) {
			fmt.Println(i, "Buzz!")
		}
	}
}

func isMultiple(number, divisor int) bool {
	return number%divisor == 0
}
