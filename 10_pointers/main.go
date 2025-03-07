package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	b := &a
	fmt.Println(a, b)

	// type check with format specifier
	fmt.Printf("%T\n", b)

	// type check with reflect
	fmt.Println("a:", reflect.TypeOf(a).Kind())
	fmt.Println("b:", reflect.TypeOf(b).Kind())

	// read value from address
	fmt.Println("Values:")
	fmt.Println("*b:", *b)
	fmt.Println("*&a:", *&a) // actually reference to the a value

	// change value with pointer
	*b = 20
	println(a)
}
