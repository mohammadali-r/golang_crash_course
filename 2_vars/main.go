package main
import "fmt"

func main(){
	// how to check the type of a var
	// var name = "MohammadAli"
	// fmt.Println(name)
	// fmt.Printf("%T\n", name)

	// how to declare a constant
	const lamp = true
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
}