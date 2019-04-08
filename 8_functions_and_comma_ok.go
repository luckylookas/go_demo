package main

import (
	"errors"
	"fmt"
	"strings"
)

func noUmlautsGoddammit(arg string) (string, error) {
	if (strings.Contains(arg, "ä") || strings.Contains(arg, "ö") || strings.Contains(arg, "ü")) {
		return "", errors.New("no umlauts goddammit")
	} else {
		return arg, nil
	}
}

type TrimmableString string


func (this *TrimmableString) trim () *TrimmableString {
	*this = TrimmableString(strings.Trim(string(*this), " "))
	return this
}

func main() {
	slice := []string{"A", "     ", "B"}

	for _, value := range slice {
		if value, err := noUmlautsGoddammit(value); err == nil {
			fmt.Println(TrimmableString(value).trim())
		} else {
			fmt.Println(err.Error())
		}
	}
}
