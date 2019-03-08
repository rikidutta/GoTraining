package main

import "fmt"

func main() {
	name := "riki"
	switch name {
	case "riki":
		fmt.Println("hei", name)
		fallthrough
	case "dutta":
		fmt.Println("hei dutta")
		fallthrough
	default:
		fmt.Println("someone else")
	}
	fmt.Println("-------")

	var mark int
	fmt.Print("enter your mark: ")
	fmt.Scan(&mark)
	switch {
	case mark >= 60:
		fmt.Println("1st div")
		fallthrough
	case mark >= 30:
		fmt.Println("pass!")
	case mark < 30:
		fmt.Println("fail")
	default:
		fmt.Println("unable to analyse the mark")
	}
	fmt.Println("-------")

	pass := "green"
	switch pass {
	case "yellow", "green":
		fmt.Println("you can play")
	case "red":
		fmt.Println("not allowed to play")
	default:
		fmt.Println("no valid pass at all")
	}
}