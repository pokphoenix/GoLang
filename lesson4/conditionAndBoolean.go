package main

import "fmt"

func main() {
	/*
		## Operater ##
		<
		>
		<=
		>=
		==
		!=
	*/

	x := 5
	y := 6.5
	val := float64(x)+1.5 != y
	fmt.Printf("%t  \n", val) // false

	/*
		5 < 5    false
		5 == 5   true
		5 == 6.5  invalid operation: x == y (mismatched types int and float64)
		float64(5) != 6.5  true
		"tim" == "Tim"   false
	*/

	fmt.Printf("%t  \n", "a" < "b")      // true
	fmt.Printf("%t  \n", "tim" == "Tim") // false

}
