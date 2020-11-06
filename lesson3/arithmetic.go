package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 int = 8
	var num2 int = 4
	answer := num1 + num2
	fmt.Printf("%d \n", answer) //12

	/*
		above code will no issue
		but when have 2 data type
		like as below
	*/
	var num3 float64 = 8
	/*
		answer := num3 / num2
		it will occur error
		mismatched type float64 and int
		use function to solve that issue
	*/
	answer2 := num3 / float64(num2)
	fmt.Printf("%f \n", answer2) //2.0000

	/*
		var num4 int = 9
		var num5 int = 4
		answer3 := (num4 / num5)
		fmt.Printf("%d \n", answer3) //2
	*/

	/*
		var num4 float32 = 9
		var num5 float32 = 4
		answer3 := (num4 / num5)
		fmt.Printf("%g \n", answer3) //2.25
	*/

	var num4 float32 = 9
	var num5 float32 = 4
	answer3 := (num4 / num5) + 5
	fmt.Printf("%g \n", answer3) //7.25

	/*
		10 % 4 = 2
		0 % 4 = 0
		4 % 3.2 //constant 3.2 truncated to integer
	*/

	/*
		var num6 int = 4
		var num7 int = 3.2
		answer = num6 % num7
		fmt.Printf("%d \n", answer) // panic: runtime error: integer divide by zero
	*/

	fmt.Printf("%g \n", math.Pow(2, 2)) //4

}
