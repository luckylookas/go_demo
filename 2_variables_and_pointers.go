package main

import "fmt"

func main() {
	str0:= "A"
	var str1 string = "B"
	var str2 *string = &str1
	var str3 = "C"
	var str4 = &str3

	fmt.Println(str0)
	fmt.Println(str1)
	fmt.Println(*str2)

	fmt.Println(str3)
	fmt.Println(*str4)

	*str2 = "D"
	fmt.Println(str1)
	fmt.Println(*str2)
}
