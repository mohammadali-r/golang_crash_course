package main

import (
	"fmt"
	"reflect"
)

func main() {
	// define map
	emails := make(map[string]string)

	// assign kv
	emails["john"] = "john@gmail.com"
	emails["alex"] = "alex@gmail.com"
	emails["peter"] = "peter@gmail.com"
	emails["sara"] = "sara@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["peter"])

	// declare map and assign kv in single line
	ids := map[int]string{1: "lily", 2: "bob", 3: "jasmin", 4: "justin"}

	fmt.Println(ids)
	fmt.Println(len(ids))
	fmt.Println(ids[2])

	// type check
	fmt.Printf("Type of emails: %T\n", emails)
	fmt.Printf("Type of ids: %T\n", ids)

	// with reflect
	fmt.Println("emails reflect:", reflect.TypeOf(emails).Kind())
	fmt.Println("ids reflect:", reflect.TypeOf(ids).Kind())

	// delete from map
	delete(ids, 3)
	fmt.Println(ids)
}
