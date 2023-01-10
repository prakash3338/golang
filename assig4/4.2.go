package main

import "fmt"

func AddFunction(a *int, b *int) int {
	return *a + *b
}

func SubFunction(a *int, b *int) int {
	return *a - *b
}

func MulFunction(a *int, b *int) int {
	return *a * *b
}

func DivFunction(a *int, b *int) (int, int) {
	return *a / *b, *a % *b
}

func main() {

	var a, b int

	fmt.Println("enter first number : ")
	fmt.Scanln(&a)
	fmt.Println("enter second number : ")
	fmt.Scanln(&b)

	fmt.Println(AddFunction(&a, &b))
	fmt.Println(SubFunction(&a, &b))
	fmt.Println(MulFunction(&a, &b))
	fmt.Println(DivFunction(&a, &b))

}
