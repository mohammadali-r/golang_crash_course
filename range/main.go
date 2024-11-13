package main

import "fmt"

func main() {
	ids := []int{23, 21, 43, 12, 54, 14, 12, 53, 63}

	// loop on range
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// loop on range
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// sum ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum =", sum)

	// range with map
	emails := map[string]string{"john": "john@gmail.com", "alex": "alex@gmail.com", "peter": "peter@gmail.com", "sara": "sara@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Key: " + k)
	}
}
