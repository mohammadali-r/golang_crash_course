package main

// how to import many packages
import (
	// how to import the package that written by yourself
	"example/hello/3_packages/strutils"
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Sqrt(9))
	// how to call a funcs from any package
	fmt.Println(strutils.Reverse("olleh"))
}
