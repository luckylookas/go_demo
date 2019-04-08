package main

import (
	"fmt"
	"strings"
)

type trimmable interface {
	trim() *trimmable
}

type TrimmableString string

func (this *TrimmableString) trim () *TrimmableString {
	*this = TrimmableString(strings.Trim(string(*this), " "))
	return this
}


func main() {
	slice := []TrimmableString{"A", "     ", "B"}

	for _, value := range slice {
		fmt.Println(*value.trim())
	}
}
