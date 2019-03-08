package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i , " is ", check(i))
	}
	fmt.Println("-----------")

	for j := 0; j < 5; j++ {
		for k := 0; k < 3; k++ {
			fmt.Println(j + k)
		}
		if j != 4 {
			fmt.Println("---")
		}
	}
	fmt.Println("-----------")

	i:= 0
	for i < 5 {
		fmt.Println("i is ", i)
		i++
	}
	fmt.Println("-----------")

	j:= 0
	for {
		j++
		if j > 10 {
			break
		}
		if j % 2 == 0 {
			continue
		}
		fmt.Println("j is ", j)
	}
}

func check(num int) string {
	if num % 2 == 0 {
		return "even \n"
	} else {
		return "odd \n"
	}
}