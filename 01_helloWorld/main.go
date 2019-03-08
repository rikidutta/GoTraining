package main

import "fmt"

func main() {
	// fmt.Println("Hello world!")
	// fmt.Printf("%d - %b - %x \n", 27, 27, 27)
	// fmt.Printf("%#x \t", 27)
	// fmt.Printf("%#X", 27)
	for i := 2900; i < 3000; i++ {
		fmt.Printf("%q \n", i)
	}
}
