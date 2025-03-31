package main

import (
	"booking-app/calculator"
	"fmt"
)

func greet(name string) string {
	return "hello " + name
}

func main() {
	message := greet("Hrithik")
	fmt.Println(message)
	sum, sub, mul, div := calculator.Calculate(2, 4)
	fmt.Println("sum is ", sum)
	fmt.Println("subtraction is ", sub)
	fmt.Println("multiplication is ", mul)
	fmt.Println("division is: ", div)

}
