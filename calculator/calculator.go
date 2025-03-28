package calculator

import (
	"fmt"
)

func CliCalculator() {
	var num1, num2 int
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter operator (+, - ,/, * ): ")
	fmt.Scan(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	switch operator {
	case "+":
		fmt.Printf("sum of these numbers is this %d: ", num1+num2)
	case "-":
		fmt.Printf("subtraction is %d :", num1-num2)
	case "*":
		fmt.Printf("multiplication is: %d", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("divide is: %d", num1/num2)
		} else {
			fmt.Printf("wrong input lil bro")
		}
	default:
		fmt.Printf("invalid arguments")
	}

}

func Calculate(num1 int, num2 int) (int, int, int, int) {
	var sum int = num1 + num2
	var sub int = num1 - num2
	var mul int = num1 * num2
	var div int = num1 / num2
	return sum, sub, mul, div
}
