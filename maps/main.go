package main

import "fmt"

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

	// delete from map
	delete(ids, 3)
	fmt.Println(ids)
}
