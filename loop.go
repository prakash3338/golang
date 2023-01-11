package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//1
	fmt.Println("First Question:")
	for i := 1; i <= 100; i++ {
		fmt.Println(i)

	}

	//2
	fmt.Println("second question:")
	i := 0
	for i < 50 {
		if i%2 == 1 {
			fmt.Println(i)
		}
		i++
	}
	//3
	fmt.Println("thrid question:")
	i = 0
	for i <= 50 {
		if i%2 == 0 {
			fmt.Println(i)
		}
		i++
	}

	//4
	fmt.Println("Fouth  question:")
	for i := 50; i < 106; i++ {
		fmt.Println(i, i/6)
	}
	//5
	fmt.Println("fifth question:")
	output := "Golang tutorial"
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	if input == output {
		fmt.Println("Welcome")
	} else {
		fmt.Println("end")
	}
	//6
	fmt.Println("Sixth question:")
	for i := 1; i <= 80; i++ {
		if i%2 == 0 && i%4 == 0 {
			fmt.Println("Golang tutorial")
		} else if i%2 == 0 {
			fmt.Println("Golang")
		} else if i%4 == 0 {
			fmt.Println("tutorial")
		}
	}
}
