package main

import "fmt"

func main() {
	num := 10
	fmt.Printf("%v %d \t %v %v \n", "address of num in dec: ", &num, "address of num in hex: ", &num)

	var celcius float64
	fmt.Print("Temp in Celcius: ")
	fmt.Scan(&celcius)
	fmt.Println("Temp in F: ", celcius / 5 * 9 + 32)

	var mark1 float64
	var mark2 float64
	var mark3 float64
	fmt.Scan(&mark1)
	fmt.Scan(&mark2)
	fmt.Scan(&mark3)
	fmt.Print("Best of 3 sub: ", mark1 + mark2 + mark3) 
}