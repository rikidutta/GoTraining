package main

import "fmt"

var A bool
var B float64

func main() {
	a := 10

	fmt.Printf("%v \n", a)
	fmt.Printf("%#v \n", a)
	fmt.Printf("%T \n", a)
	fmt.Printf("%v \n", A)
	fmt.Printf("%v \n", B)
	fmt.Printf("%T \n", 0)
}