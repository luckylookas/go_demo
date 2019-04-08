package main

import "fmt"

func main() {
	fmt.Println("INITIALIZE WITH COMPOSITE LITERAL")
	strMapComposite := map[string]string{"one": "A", "two": "B"}
	fmt.Println(strMapComposite["one"])

	fmt.Println("INITIALIZE WITH MAKE")
	strMapMake := make(map[string]string, 2)
	strMapMake["one"] = "A"
	fmt.Println(strMapComposite["one"])

	fmt.Println("ALLOCATE WITH NEW allocates a 'Zero' map in memory. unfortunately a zero map can not be written to")
	strmapNew := new(map[string]string)

	(*strmapNew)["one"] = "A"
	fmt.Println((*strmapNew)["one"])
}
