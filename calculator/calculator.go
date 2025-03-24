package calculator

import (
	"fmt"
)

func CliCalculator() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter operator (+, - ,/, * ): ")
	fmt.Scan(&operator)

	fmt.Print("Enter first number: ")
	fmt.Scan(&num2)

	switch operator {
	case "+":
		fmt.Print("sum of these numbers is this: ", num1+num2)
	case "-":
		fmt.Printf("subtraction is :", num1-num2)
	case "*":
		fmt.Printf("multiplication is: ", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("divide is: ", num1/num2)
		} else {
			fmt.Printf("wrong input lil bro")
		}
	default:
		fmt.Printf("invalid arguments")
	}

}
