package main

import "fmt"

func main() {
	a := 1
	var b *int = &a
	fmt.Printf("'value': %v, *b: %v, type: %T \n", b, *b, b)
	*b = 2
	fmt.Println(`a:`, a)

	fmt.Print("enter value of *b: ")
	fmt.Scan(&*b)
	fmt.Println(a)
	fmt.Println("---")
	
	fmt.Print("enter another value of a: ")
	fmt.Scan(b)
	fmt.Println(a)
	fmt.Println("---")

	d := 10
	fmt.Print("new value of d: ")
	makeZero(d)
	fmt.Println(d)
	makeZeroFromMemory(&d)
	fmt.Println(d)
	makeTenFromMemory2(d)
	fmt.Println(d)

}

func makeZero(x int) {
	fmt.Println("start makeZero")
	x = 0
	fmt.Println("end makeZero")
}

func makeZeroFromMemory(x *int) {
	fmt.Println("start makeZeroFromMemory")
	*x = 0
}

func makeTenFromMemory2(x int) {
	fmt.Println("start makeZeroFromMemory2")
	*&x = 10
}