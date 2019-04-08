package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func square_multiply(value int, out chan<- float64) {
	out<- float64(value) * float64(value)
}


func square_pow(value int, out chan<- float64) {
	out <- math.Pow(float64(value),2.0)
}

func square_chan_multiply(in <-chan int, out chan<- float64) {
	value := <-in
	out<- float64(value) * float64(value)
}

func square_chan_pow(in <-chan int, out chan<- float64) {
	out <- math.Pow(float64(<-in),2.0)
}

func main() {
	output := make(chan float64)
	defer close(output)

	fmt.Println("race two implementations for the fastest result")
	go square_multiply(1901023213124, output)
	go square_pow(1901023213124, output)

	fmt.Println(<-output)

	fmt.Println("race two implementations for the fastest result with a timeout and know who won")

	input_mul, input_pow := make(chan int), make(chan int)
	output_mul, output_pow := make(chan float64), make(chan  float64)

	go square_chan_multiply(input_mul, output_mul)
	go square_chan_pow(input_pow, output_pow)

	input_mul<- 1901023213124
	input_pow <- 1901023213124
	close(input_pow)
	close(input_mul)

	select {
		case <-time.After(50 * time.Nanosecond):
			fmt.Println("both functions took longer than 50 ns.")
			break;
		case value := <-output_pow:
			fmt.Println("pow was faster: ", value)
			break;
		case value := <- output_mul:
			fmt.Println("multiply was faster: ", value)
			break;
	}

	fmt.Println("read from channels until they are closed")

	stringChannel := make(chan string)
	go func(in <-chan string) {
		for value := range in {
			fmt.Println(value)
		}
	}(stringChannel)
	stringChannel<-"a"
	stringChannel<-"b"
	stringChannel<-"c"

	fmt.Println("buffered channels")

	close(stringChannel)
	stringChannel = make(chan string, 4)

	stringChannel<-"a"
	stringChannel<-"b"
	stringChannel<-"c"
	stringChannel<-"c"

	go func(in <-chan string) {
		for value := range in {
			fmt.Println(value)
		}
	}(stringChannel)

	time.Sleep(250 * time.Millisecond)

	fmt.Println("wait for go routines to finish")
	close(stringChannel)
	stringChannel = make(chan string, 4)
	wg := sync.WaitGroup{}

	stringChannel<-"a"
	stringChannel<-"b"
	stringChannel<-"c"
	stringChannel<-"c"
	close(stringChannel)
	wg.Add(1)
	go func(in <-chan string) {
		for value := range in {
			fmt.Println(value)
		}
		wg.Done()
	}(stringChannel)
	wg.Wait()
}
