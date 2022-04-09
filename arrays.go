package main

import "fmt"

func main() {
	var array2 [5]int
	array2 = [5]int{20, 20, 30, 40, 50}
	fmt.Println(array2)
	array2 = [5]int{1: 10, 3: 30}
	fmt.Print(array2)

	var array [4][2]int
	// Use an array literal to declare and initialize a two
	// dimensional integer array.
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Print(array)
	// Declare and initialize index 1 and 3 of the outer array.
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	fmt.Print(array)
	// Declare and initialize individual elements of the outer
	// and inner array.
	array = [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Print(array)

	var array3 [2][2]int
	array3 = [2][2]int{{1, 2}, {3, 4}}
	fmt.Print(array3)
	//Passing array address as the argument to the foo function which takes the pointer as argument
	foo(&array2)
}

func foo(array *[5]int) {
	fmt.Println("Inside Foo")
	fmt.Println(array)
}
