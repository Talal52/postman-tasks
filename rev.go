package main

import "fmt"

func reverse(a int) {
	for a > 0 {
		n := a % 10
		a = a / 10
		fmt.Print(n)
	}
}
func main() {
	fmt.Print("Enter any Number: ")
	var input int
	fmt.Scanln(&input)
	reverse(input)
}