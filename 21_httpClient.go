package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	for i := 0; i< 200000; i++ {
		str := fmt.Sprintf("http://localhost:8080/api/%d/%d",  rand.Intn(1000000),  rand.Intn(5000))
		http.Post(str, "applicaion/json", nil)
	}
}
