package main

import "fmt"

func main() {
	for i := 190; i < 200; i++ {
		fmt.Println(i, string(i), []byte(string(i)))
	}
	fmt.Printf("%T %T \n", 'r', "r")

	var marks [3] int
	for j := 0; j < 3; j++ {
		var num int
		fmt.Print(j + 1, ". enter mark: ")
		fmt.Scan(&num)
		marks[j] = num
	}
	fmt.Println(marks)
}