package main

import "fmt"

// go vet 16_go_vet.go

func main() {
	var s string = "1"
	u := s
	fmt.Println(s)
	return
	u = s
}
