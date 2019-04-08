package main

import "fmt"

func main() {
	s := ""
	if len("") > 0 {
		fmt.Println(s)
	} else {
		fmt.Println("na geeeeh")
	}

	switch s {
	case " ":
		fallthrough
	case "":
		fmt.Println("nothing")
	default:
		fmt.Println("default")
	}
}
