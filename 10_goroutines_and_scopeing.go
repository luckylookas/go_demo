package main

import (
	"fmt"
	"strings"
	"time"
)

func print(str string) {
	fmt.Println(str)
}

func main() {
	slice := []string{"A  ", "     ", "B "}

	for _, str := range slice {
		go func() {
			fmt.Println(strings.Trim(str, " "))
		}()
	}

	time.Sleep(250 * time.Millisecond)
	/*prints B,B,B because goroutines are not threads, and are not scheduled by the OS
		the go runtime sees: the routine will do I/O, which is slow and yields execution -
	so it decides to finish the loop, which requires no i/o.
	and as the variable is closured, all routines point to the same address in memory -> all print (alsmost all the time) "B"
	*/

	for _, str := range slice {
		go func() {
			str := str //freaky
			fmt.Println(strings.Trim(str, " "))
		}()
	}

	time.Sleep(250 * time.Millisecond)


	for _, str := range slice {
		go func(str string) {
			fmt.Println(strings.Trim(str, " "))
		}(str)
	}

	time.Sleep(250 * time.Millisecond)


	for _, str := range slice {
		go print(str)
	}

	time.Sleep(250 * time.Millisecond)
}
