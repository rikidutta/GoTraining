package main

import "fmt"

func main() {
	var str bool
	if str {
		fmt.Println(str, " is true.")
	} else {
		fmt.Println("no variable")
	}

	if num := 10; !str {
		fmt.Println(num)
	}

	var mark int
	fmt.Print("enter mark: ")
	fmt.Scan(&mark)
	if mark >= 60 {
		fmt.Println("1st div")
	} else if mark >= 30 {
		fmt.Println("pass")
	} else {
		fmt.Println("fail")
	}

	fmt.Println(100 / 3, " && ", 100.0 / 3.0)

	var total int
	for num := 1; num < 1000; num++ {
		if num % 3 == 0 || num % 5 == 0 {
			total += num
		}
	}
	fmt.Println(total)
}