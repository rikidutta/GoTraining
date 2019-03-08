package main

import "fmt"

func main() {
	slice1 := []int{1, 3}
	slice1[0] = 5
	fmt.Println("slice1:", slice1)
	slice1 = append(slice1, 7)
	fmt.Println("slice1:", slice1)
	fmt.Println("-----")

	var slice2 = make([]int, 2)
	slice2[0] = 3
	fmt.Println("slice2", slice2)
	slice2 = append(slice2, 7)
	fmt.Println(slice2)
	fmt.Println("-----")

	var slice3 = make([]int, 0)
	for i := 0; i < 10; i++ {
		slice3 = append(slice3, i)
	}

	fmt.Println("slice3", slice3)
	fmt.Println("[2:5]", slice3[2:5])
	fmt.Println("[2:]", slice3[2:])
	fmt.Println("[:5]", slice3[:5])
	fmt.Print("deleting elements of slice3 2nd element:", append(slice3[:1], slice3[2:]...))
	fmt.Println("-----")


	fmt.Println("riki[1]:", "riki"[1], ", string:", string("riki"[1]))
	fmt.Println("-----")

	slice4 := make([]string, 3, 4)
	slice4 = append(slice4, "c")
	slice4 = append(slice4, "d")
	slice4 = append(slice4, "e")
	slice4 = append(slice4, "f")
	fmt.Println("slice4:", slice4, "len:", len(slice4), "capacity:", cap(slice4))
	fmt.Println("slice4[3]:", slice4[3])
	fmt.Println("-----")

	slice5 := make([][]string, 0)

	subSlice1 := make([]string, 0)
	subSlice1 = append(subSlice1, "a", "b", "c")
	subSlice2 := make([]string, 0)
	subSlice2 = append(subSlice2, slice4...)
	subSlice3 := make([]string, 0)
	subSlice3 = append(subSlice3, []string{"x", "y", "z"}...)

	slice5 = append(slice5, subSlice1, subSlice2, subSlice3)
	fmt.Println("slice5:", slice5)
	fmt.Printf("type of slice5: %T & that of slice5[i]: %T \n", slice5, slice5[0])
	fmt.Println("-----")

	slice6 := make([]int, 1)
	slice6[0] = 10
	slice6[0]++
	fmt.Println("slice6[0]:", slice6)

	nilSlice := make([]int, 0)
	var nilSlice2 []string
	fmt.Println(nilSlice == nil)
	fmt.Println(nilSlice2 == nil)
	fmt.Println(nil)
}