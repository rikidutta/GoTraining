package main

import "fmt"

func main() {
	fmt.Println("Map starts:")
	defer fmt.Println("Map ends.")

	m1 := map[string]float64{}
	m1["orange"] = 5
	m1["apple"] = 8.5
	fmt.Println("m1:", m1)
	m1["orange"] = 4
	delete(m1, "apple")
	fmt.Println("m1:", m1)
	fmt.Println("-----")

	m2 := make(map[string]float64)
	fmt.Println("m2:", m2)
	if v, prs := m2["banana"]; prs {
		fmt.Println("Already exist! m2[banana]:", v)
	} else {
		fmt.Print(`creating m2["banana"]...`, "\nset new price for banana: ")
		var price float64
		fmt.Scan(&price)
		m2["banana"] = price
		fmt.Println(`created m2["banana"]:`, m2["banana"])
	}
	fmt.Println("-------")

	ChemicalElements := make(map[string]map[string]string)
	
	Hydrogen := make(map[string]string)
	Hydrogen["symbol"] = "H"
	Hydrogen["At no."] = "1"
	ChemicalElements["Hydrogen"] = Hydrogen

	createNewElement(ChemicalElements, "Helium", "He", "2")
	createNewElement(ChemicalElements, "Helium", "He", "2")

	fmt.Println("ChemicalElements:", ChemicalElements)
	fmt.Println("-------")

	for key, value := range ChemicalElements {
		if _, exist := ChemicalElements["Lithium"]; !exist {
			createNewElement(ChemicalElements, "Lithium", "Li", "3")
			defer fmt.Println("Lithium Added later.")
		}
		fmt.Println(key, ":", value)
	}
}

func createNewElement(ChemicalElements map[string]map[string]string, elementName string, symbol string, At_no string) {
	if _, prs := ChemicalElements[elementName]; prs {
		fmt.Println("Element with name", elementName, "already exist in", ChemicalElements)
	} else {
		element := make(map[string]string)
		element["symbol"] = symbol
		element["At no."] = At_no
		ChemicalElements[elementName] = element
		fmt.Println("Element", elementName, "created.")
	}
}
