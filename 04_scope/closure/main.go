package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

var x int = 0

func incre() int {
	x++;
	return x
}

func main() {
	fmt.Println("inside main():")
	fmt.Println(incre())
	fmt.Println(incre())
	fmt.Println(x)
	{
		y := 10
		fmt.Println("y = ", y)
	}
	fmt.Println(main2())

	z := 10
	incre2 := func () int {
		z++
		return z
	}
	fmt.Println("incre2() gives: ", incre2())

	intfnMain := intfn()
	fmt.Println("1. intfn() gives: ", intfnMain())
	fmt.Println("2. intfn() gives: ", intfnMain())

	{
		res, _ := http.Get("https://www.google.com/")
		page, _ := ioutil.ReadAll(res.Body)
		res.Body.Close();
		fmt.Printf("%s", page);
	}
	fmt.Println("-----------------")
	fmt.Println("-----------------")
	fmt.Println("-----------------")
	fmt.Println("-----------------")
	fmt.Println("-----------------")
	fmt.Println("-----------------")
	fmt.Println("-----------------")
	{
		res, _ := http.Get("https://www.facebook.com/")
		page, _ := ioutil.ReadAll(res.Body)
		res.Body.Close();
		fmt.Printf("%s", page);
	}

}

func main2() int {
	return x
}

func intfn() func() int {
	num := 100
	fmt.Println(num)
	return func() int {
		num++
		return num
	}
}