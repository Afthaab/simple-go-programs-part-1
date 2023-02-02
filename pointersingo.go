package main

import "fmt"

func main() {
	a := 10
	var ptr = &a
	fmt.Println("the address :", ptr)
	fmt.Println("the value :", *ptr)

	//Any changes made in the pointer fo the variable will change the actual memory
	*ptr = *ptr * 2
	fmt.Println("the new value of a :", a)

	// you cannot intialize the pointer of dffrent tye compared to the variable
	// b := "Afthab"
	// var ptr2 *int
	// ptr2 = &b
	// fmt.Println("the new value of a :", *ptr2)

}
