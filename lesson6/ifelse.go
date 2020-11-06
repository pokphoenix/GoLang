package main

import "fmt"

func main() {

	name := "pok"

	fmt.Println("Before if")
	if name == "pok" {
		fmt.Println("inside if")
	}
	fmt.Println("after if")

	age := 13
	if age >= 18 {
		fmt.Println("You 18+")
	} else if age >= 14 {
		fmt.Println("You can go with parent")
	} else {
		fmt.Println("You too young to enter")
		fmt.Printf("Wait %d years \n", 18-age)
	}

}
