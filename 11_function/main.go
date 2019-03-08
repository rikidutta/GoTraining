package main

import "fmt"

func main() {
	func() {
		fmt.Println("Function section starts:")
	}();
	defer func() {
		fmt.Println("Function section ends.")
	}();

	fmt.Println(greet("riki", 22))
	fmt.Println("---")
	fmt.Println(greet2("ricky", 11))
	fmt.Println("---")
	defer fmt.Println(greetMulti("riki", 22))
	defer fmt.Println(greetMulti("riki2", 22))
	fmt.Println("-------")

	var marks = [] float64{97, 94, 93}
	fmt.Println("avg variadic:", avgVariadic(97, 94, 93))
	fmt.Println("avg variadic2:", avgVariadic(marks...))
	fmt.Println("avg slice:", avgSlice(marks))
	fmt.Println("-------")

	fmt.Println("adding 7 and 9:", retFunc("add")(7, 9))
	fmt.Println("-------")

	callbackFunc("riki", welcome)
	fmt.Println("-------")

	var n int
	fmt.Print("Enter number to get factorial value: ")
	fmt.Scan(&n)
	fmt.Println(factorial(n))
	fmt.Println("-------")

	fmt.Println("Max number:", max([] int{6, 5, 4, 3, 2, 5, 6, 7, 8, 10}))
	fmt.Println("----------")
}

func greet(name string, age int) string {
	fmt.Println("Hey!", name, ", You are ", age, " years old.")
	return fmt.Sprint("DoB: ", 2019 - age)
}

func greet2(name string, age int) (s string) {
	fmt.Println("Hey!", name, ", You are ", age, " years old.")
	s = fmt.Sprint("DoB: ", 2019 - age)
	return
}

func greetMulti(name string, age int) (string, string) {
	return fmt.Sprintln("Your Name:", name),
	fmt.Sprint("Your age:", age)
}

func avgVariadic(marks... float64) float64 {
	fmt.Println(marks)
	total := 0.0
	for _, v := range marks {		// _ : index, we donot want that here
		total += v
	}
	return total / float64(len(marks))
}

func avgSlice(marks [] float64) float64 {
	fmt.Println(marks)
	total := 0.0
	for _, v := range marks {		// _ : index, we donot want that here
		total += v
	}
	return total / float64(len(marks))
}

func retFunc(oparator string) func(a, b int) string {
	switch oparator {
	case "add":
		return func (a, b int) string {
			return fmt.Sprint(a + b)
		}
	case "multi":
		return func (a, b int) string {
			return fmt.Sprint(a * b)			
		}
	default:
		return func (a, b int) string {
			return fmt.Sprint("Operation not available")
		}
	}
}

func callbackFunc(name string, callback func(string)) {
	callback(name)
}

func welcome(name string) {
	fmt.Println("welcome!", name)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n - 1)
}

func max(numbers[] int) int {
	var largest int
	for _, number := range numbers {
		if number > largest {
			largest = number
		}
	}
	return largest
}