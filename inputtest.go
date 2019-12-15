package main

import "fmt"

func main(){
	// creating arrays of 6 elements of type int,
	// and put elements 1,2,3,4,5 and 6 inside of it in that particular order.

	var array1 [6]int = [6]int {1, 2, 3, 4, 5, 6} // classical way
	var array2 = [6]int {1, 2, 3, 4, 5, 6} // a less verbose way

	fmt.Println("array1:", array1)
	fmt.Println("array2:", array2)
}