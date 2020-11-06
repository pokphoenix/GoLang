package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

/*
  main redeclared in this block error will occur 
  when you have multiple .go files in a single package (i.e. package main) 
  then you can only have one main() function.
  that why i have to create folder lesson1 , lesson2 for solve this issue
*/
func main() { 
	scanner := bufio.NewScanner(os.Stdin)  
	fmt.Printf("Type something: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You typed: %q \n",input)

	fmt.Printf("Type the year you were born: ")
	scanner.Scan()
	
	/*
	inputYear := scanner.Text()
	fmt.Printf("You will be %d years old at the end of 2020", 2020 - inputYear ) 
		//above code will occur error
		//cannot use 2020(type untype int) as type string
		//use strconv to solve this issue 
		//see at below code
	*/

	inputYear,_ := strconv.ParseInt(scanner.Text(),10,64)
	fmt.Printf("You will be %d years old at the end of 2020", 2020 - inputYear ) 

	/* 
	   if type string instead of birth year 
	   system will return 2020 because strconv return empty 
    */

}

