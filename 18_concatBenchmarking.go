package main

import (
	"fmt"
	"strings"
)
/*
go build concat
./concat -> profile.out
go tool pprof --pdf concat <profile.out> > file.pdf
*/

func main() {

}

func concatSprintf (a string, b string) string {
	return fmt.Sprintf("%s%s", a, b)
}

func concatBuilder (a string, b string) (string) {
	builder := strings.Builder{}
	builder.WriteString(a)
	builder.WriteString(b)
	return builder.String()
}

func concatBuilderGrow (a string, b string) (string) {
	builder := strings.Builder{}
	builder.Grow(len(a)+len(b))
	builder.WriteString(a)
	builder.WriteString(b)
	return builder.String()
}

func concatSlice (a string, b string) (string) {
	value := make([]byte, len(a) + len(b))
	copy(value, []byte(a))
	copy(value[len(a):], []byte(b))
	return string(value)
}

func concatAppend (a string, b string) (string) {
	value := []byte{}
	value = append(value, a...)
	value = append(value, b...)
	return string(value)
}
