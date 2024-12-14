package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// first_name string
	// last_name  string
	// gender     string
	// city       string
	// age        int

	first_name, last_name, gender, city string
	age                                 int
}

// two type method: 1. value receiver (that doesn't change the values), 2. pointer receiver (that changes the values)

// greet method (value receiver)
func (p Person) greet() string {
	return "My name is " + p.first_name + " " + p.last_name + " and I am " + strconv.Itoa(p.age)
}

// grow_up method (pointer receiver)
func (p *Person) grow_up() {
	p.age++
}

// get_married method (pointer receiver)
func (p *Person) get_married(spouse_last_name string) {
	if p.gender == "male" {
		return
	} else {
		p.last_name = spouse_last_name
	}
}

func main() {
	p1 := Person{first_name: "alex", last_name: "williams", gender: "male", city: "hamborg", age: 25}
	p2 := Person{"sara", "johnson", "female", "new york", 23}
	fmt.Println(p1)
	fmt.Println(p2)
	p1.grow_up()
	p1.get_married("hugo")
	fmt.Println(p1.greet())
	p2.grow_up()
	p2.get_married("jonas")
	fmt.Println(p2.greet())
}
