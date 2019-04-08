package main
/*
#cgo CFLAGS: -g -Wall
#include <stdio.h>
#include <stdlib.h>
#include <regex.h>
int regmatch(char* target, char* pattern, int cflags) {
	regex_t reg;
	regmatch_t match;
	regcomp(&reg, pattern, cflags);
	int execRet = regexec(&reg, target, 1, &match, 0);
	regfree(&reg);
	return execRet;
}
*/
 import "C"
 import (
	 "fmt"
	 "unsafe"
 )

func main() {
	name := C.CString("Hi")
	defer C.free(unsafe.Pointer(name))

	if result := C.regmatch(C.CString("aaabaa"), C.CString("abc"), C.int(0)); result == 0 {
		fmt.Println("match")
	} else {
		fmt.Println("no match")
	}
}