package main

import "fmt"

// how to declare a custom function
func greeting(name string) string {
	return "Welcome " + name
}

// another function
func get_sum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("MohammadAli"))
	fmt.Println(get_sum(10, 15))
}
