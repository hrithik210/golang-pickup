package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func greet(p person) string {
	return "Hello " + p.name + "you are" + fmt.Sprint(p.age) + "years old"
}

func main() {
	var p person = person{
		name: "Bombardino crocodillo",
		age:  1,
	}
	fmt.Println(greet(p))
}
