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

// variadic function
func all_sum(nums ...int) {
	fmt.Print(nums, " ")
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}

func main() {
	fmt.Println(greeting("MohammadAli"))
	fmt.Println(get_sum(10, 15))
	all_sum(1, 3, 5, 3, 2, 8)

	nums := []int{54, 21, 64, 32, 1, 6, 3} // pass a slice to function as argument
	all_sum(nums...)
}
