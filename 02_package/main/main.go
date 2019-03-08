package main

import "fmt"
import "github.com/GoRiki/GoTraining/GoTraining/02_package/stringsUtil"

func main() {
	fmt.Println(stringsUtil.MyName)
	fmt.Println(stringsUtil.FullName2("riki1", "dutta"))
	fmt.Print(stringsUtil.FullName2("riki2", "dutta"))
	fmt.Print(stringsUtil.FullName2("riki3", "dutta"))
	fmt.Println(stringsUtil.FullName2("riki4", "dutta"))
	fmt.Print(stringsUtil.FullName2("riki5", "dutta"))
}