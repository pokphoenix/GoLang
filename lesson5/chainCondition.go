package main

import "fmt"

func main() {
	fmt.Printf("%t \n", 5 < 7 || 7 > 9)            // true
	fmt.Printf("%t \n", (true || false) && !false) // true

	x := 8
	val := (true || false) && !!false
	val2 := val || x > 9
	fmt.Printf("%t \n", val2) //false

}
