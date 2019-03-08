package main

import "fmt"

func main() {
	var arr [257] byte	
	for i := 0; i < 257; i++ {
		arr[i] = byte(i + 1)
	}
	fmt.Println(arr)
	fmt.Println(arr[256])

	for j := 0; j < 257; j++ {
		fmt.Printf("%b \n", arr[j])
	}
}

