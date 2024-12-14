package main

import "fmt"

func main() {
	// how to declare arrays in some ways
	// and assign values to them
	var cars [3]string
	cars[0] = "Mercedes Benz"
	cars[1] = "BMW"
	cars[2] = "Audi"
	fruits := []string{"apple", "orange", "bannana", "grape", "cherry"}

	fmt.Println(cars)
	fmt.Println(fruits)
	fmt.Println(cars[2])
	fmt.Println(fruits[1])
	// slices stuff
	fmt.Println(len(fruits))
	fmt.Println(fruits[1:3])
	new_array := fruits[1:3]
	fmt.Printf("%T\n", new_array)
}
