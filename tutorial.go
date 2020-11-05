package main

import "fmt" 

func main() {

	/*var name string = "Hello pok"*/
	var name2 string 
	name2= "pok"
	name2= "aha"

	var number uint = 260
	//var number2 uint8 = 260  //over flow
	var number3 uint16 = 260
	number3 += 5





	fmt.Println("Hello world!\n",name2,number,number3)

	/*
	fmt.Printf("%T",name2) //string
	fmt.Printf("%T",number) //uint
	fmt.Printf("%T",number3) //uint16
	fmt.Printf("%T",2) //int 
	*/


	//number4 := 555
	//fmt.Printf("%T",number4) //int


	fmt.Printf("Hello %T %v \n",10,10) //Hello int 10

	var out string = fmt.Sprintf("Number : %07d is cool \n",45)
	fmt.Println(out)
}
// run command on terminal :     go run tutorial.go
