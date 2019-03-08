package main

import "fmt"

const number = 10
const name = "riki"

const (
	zero = iota
	one = iota
	two = iota
)

func main() {
	fmt.Println(number)
	fmt.Printf("%d \t %b \n", 1 << (one), 1 << (one))
	fmt.Printf("%d \t %b \n", 2 << (one), 2 << (one))
	fmt.Printf("%d \t %b \n", 3 << (one), 3 << (one))
	fmt.Printf("%d \t %b \n", 4 << (one), 4 << (one))
	fmt.Printf("%d \t %b \n", 5 << (one), 5 << (one))
	fmt.Printf("%d \t %b \n", 6 << (one), 6 << (one))
	fmt.Printf("%T \n", number)
	fmt.Printf("%T \n", number + 5.3)
	// fmt.Printf("%T \n", name + 5.3)
}