package main

import (
	"fmt"
	"math/rand"
	"time"
)

//go run -race 17_go_race.go

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

func main() {
	//if the initial value vom randomDuration is very small, the goroutine could reach r.Reset, before the main goroutine has executed the assignment t = ...
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		t.Reset(randomDuration())
	})
	time.Sleep(5 * time.Second)
}
