package main

import "fmt"

func main() {
	// how to declare a variable
	var fruit string = "orange"
	fmt.Println(fruit)
	var number int = 26
	fmt.Println(number)

	// how to check the type of a var
	// var name = "MohammadAli"
	// fmt.Println(name)
	// fmt.Printf("%T\n", name)

	// how to declare a constant
	const num int = 15
	const lamp = true // type inference, Go can infer the type of the constant
	fmt.Println(num)
	fmt.Println(lamp)
	fmt.Printf("%T\n", lamp)

	// how to declare a vars in second way, just use it in funcs
	// family := "Rohanian"
	// fmt.Println(family)
	// fmt.Printf("%T\n", family)

	// how to declare two vars in single line
	name, family := "MohammadAli", "Rohanian"
	fmt.Println(name, family)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", family)

	fmt.Printf("%v\n", fruit)
	fmt.Printf("%v\n", number)
	fmt.Printf("%v\n", lamp)
	fmt.Printf("%v\n", name)
}
