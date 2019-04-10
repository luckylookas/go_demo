package main

import (
	"fmt"
	"strconv"
)

type Duck interface {
	Quack()
}

type RegularDuck struct {
	Name string
}

func (this RegularDuck) Quack() {
	fmt.Println("Quack")
}

func main() {

	regular := RegularDuck{Name: "Roy"}

	var duck Duck = regular

	//assert that duck is indeed a RegularDuck, only works with interfaces - this is how you get the actual type back from an interface
	var asserted RegularDuck = duck.(RegularDuck)

	asserted.Quack()

	var i int = 10
	//type conversion from number to number
	var f float32 = float32(i)

	fmt.Println(f)

	//type conversion for strings work with the strconv package
	fmt.Println(strconv.ParseFloat("1.1", 32))

}
