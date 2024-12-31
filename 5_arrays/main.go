package main

import (
	"fmt"
	"reflect"
)

func main() {
	// how to declare arrays in some ways
	// and assign values to them
	var cars [3]string // this is an array with fixed-size
	cars[0] = "Mercedes Benz"
	cars[1] = "BMW"
	cars[2] = "Audi"
	nums := [6]int{1, 43, 23, 76, 23, 431}

	fruits := []string{"apple", "orange", "bannana", "grape", "cherry"} // this is a slice with dynamic size
	foods := make([]string, 3, 6)                                       // another way to declare a slice with capacity
	foods[0] = "pizza"
	foods[1] = "salad"
	foods[2] = "sandwich"

	// arrays stuff
	fmt.Println("cars= ", cars)
	fmt.Println(cars[2])
	fmt.Println("nums= ", nums)
	fmt.Println(len(nums))
	slice := nums[1:4] // this create a slice too
	fmt.Println("slice of nums= ", slice)

	// slices stuff
	fruits = append(fruits, "watermelon")
	fmt.Println("fruits= ", fruits)
	fmt.Println(fruits[1])
	fmt.Println(len(fruits))
	fmt.Println(fruits[1:3])

	// type check
	fmt.Printf("Type of cars: %T\n", cars)
	fmt.Printf("Type of fruits: %T\n", fruits)
	fmt.Printf("Type of foods: %T\n", foods)

	// type check with reflect
	fmt.Println("Type check with reflect:")
	fmt.Println("cars:", reflect.TypeOf(cars).Kind())
	fmt.Println("fruits:", reflect.TypeOf(fruits).Kind())
	fmt.Println("foods:", reflect.TypeOf(foods).Kind())
}
