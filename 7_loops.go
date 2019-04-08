package main

import "fmt"

func main() {
	slice := []string{"A", "B", "NO!"}

	//range returns index, value - but we do not care about the index
	for _, str := range slice {
		fmt.Println(str)
	}

}
