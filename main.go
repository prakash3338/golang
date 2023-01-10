package main

import (
	"calculator/app"
	"fmt"
)

func main() {

	var a, b int

	fmt.Println("enter first number : ")
	fmt.Scanln(&a)
	fmt.Println("enter second number : ")
	fmt.Scanln(&b)

	quotient, remainder := app.DivFunction(a, b)

	fmt.Println("Addition of ", a, "and", b, ":", app.AddFunction(a, b))
	fmt.Println("Subtraction of ", a, "and", b, ":", app.SubFunction(a, b))
	fmt.Println("Multiplication ", a, "and", b, ":", app.MultiplyFunction(a, b))
	fmt.Println("Quotient of ", a, "and", b, ":", quotient)
	fmt.Println("Remainder of ", a, "and", b, ":", remainder)
}
