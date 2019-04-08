package main

import "fmt"

func main() {
	// ARRAYS
	fmt.Println("ARRAY")
	str := [2]string{"A", "B"}
	fmt.Println(str[0])

	// "array" is a type, not a pointer, so unlike in C, fmt.Println(*str) gives a compiler error
	fmt.Println("ARRAYS ARE NOT POINTERS")
	str1 := str
	str1[0] = "C"
	fmt.Println(str1[0])
	fmt.Println(str[0])

	fmt.Println("SLICES")
	// SLICES
	strSlice := str[0:1]
	fmt.Println(len(strSlice), ": ", strSlice[0])
	strSlice = str[1:2]
	fmt.Println(len(strSlice), ": ", strSlice[0])

	fmt.Println("SLICES ARE POINTERS")
	slice2 := strSlice
	slice2[0] = "D"
	fmt.Println("strSlice: ", strSlice[0])
	fmt.Println("slice2: ", slice2[0])

}
