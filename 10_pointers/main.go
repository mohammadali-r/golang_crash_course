package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Println(a, b)

	fmt.Printf("%T\n", b)

	// read value from address
	println(*b)
	println(*&a) // actually reference to the a value

	// change value with pointer
	*b = 20
	println(a)
}
