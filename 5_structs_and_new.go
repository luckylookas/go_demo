package main

import "fmt"

type number int32

type person struct {
	name string
	age number
	address address
	clothes
}

type address struct {
	street string
	house number
}

type clothes struct {
	piece string

}
/*
"zero values":
false for booleans,
0 for integers,
0.0 for floats,
"" for strings,
nil for pointers, functions, interfaces, slices, channels, and maps.
*/

func main() {
	p1 := person{name: "Lukas", age: 29, address: address{street: "Daheim", house: 1}, clothes: clothes{piece: "pants"} }

	//pointers
	p2 := &person{name: "Michael", age: 28, address: address{street: "Daheim", house: 1}, clothes: clothes{piece: "nothing"}}
	//create a pointer to a "zero" person
	p3 := new(person)

	fmt.Println(p1.name, "is ", p1.age, " and lives in ", p1.address.street, ", wearing ", p1.piece, " whenever he can.")
	fmt.Println(p2.name, "is ", p2.age, " and lives in ", p2.address.street, ", wearing ", p2.piece, " whenever he can.")
	fmt.Println(p3.name, "is ", p3.age, " and lives in ", p3.address.street, ", wearing ", p3.piece, " whenever he can.")
}
