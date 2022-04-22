package main

import "fmt"

//test3
func main() {
	var slices []int
	var slices2 []string
	var slices3 []string
	//Creating slice using make method
	slices = make([]int, 5)
	fmt.Println(slices)

	slices3 = make([]string, 3, 5) //Length of 3 and capacity of 5.
	//That is the slice has the access to 3 elements although the underlying array has 5 elements
	fmt.Println(slices3)

	//If nothing is given inside [] it will be slice and if something is given then it is array
	slices2 = []string{"1", "2", "3", "4"}
	fmt.Println(slices2)

	slices2 = append(slices2, "5")
	fmt.Println(slices2)

	var arr [5]int   //Declaring an Array
	var slice4 []int //Declaring an Slice
	arr = [5]int{1, 2, 3, 4, 5}
	slice4 = []int{1, 2, 3, 4, 5}
	fmt.Print(arr)
	fmt.Print(slice4)

	//Nil slices
	var slice5 []int
	var slice6 []int
	slice5 = make([]int, 0)
	fmt.Println(slice5)
	//slice5 = []int{}

	//Accessing slice elements. Same as Array.
	fmt.Println(slice4[2])

	//Taking slice of an Slice
	var subSlice5 []int
	slice6 = []int{1, 2, 3, 4, 5}
	subSlice5 = slice6[1:3]
	subSlice5[1] = 2
	//IMP: Here the slices and subslices share the same underlying array.
	//So changes in one may be reflected in another
	fmt.Println(subSlice5)
	// 	A slice can only access indexes up to its length. Trying to access an element outside
	// of its length will cause a runtime exception. The elements associated with a slice’s
	// capacity are only available for growth.

	slice6 = append(slice6, 30)
	fmt.Println(slice6)
	var slice7 []int
	slice7 = slice6[2:3:4]
	/*Again, the first value represents the starting index position of the element the new slice
	will start with—in this case, 2. The second value represents the starting index position
	(2) plus the number of elements you want to include (1); 2 plus 1 is 3, so the second
	value is 3. For setting capacity, you take the starting index position of 2, plus the number
	of elements you want to include in the capacity (2), and you get the value of 4.
	Length: j - i or 3 - 2 = 1
	Capacity: k - i or 4 - 2 = 2
	*/

	slice8 := slice6[2:3:10] //Since slice8 was undeclared. To define and declare simultaneously we use ':='
	fmt.Println(slice7)
	fmt.Println("slice 8 :")
	fmt.Println(slice8)

	slice9 := []string{"Apple", "Banana", "Kiwi", "Grapes"}
	slice10 := slice9[2:3:3]
	fmt.Println(slice10)
	//Now we append a new String in slice10. Since the (accessible in underlying array pointed by both slice10 and slice9 )capacity is just 1 for slice10.
	//New Underlying array is formed for slice10

	slice10 = append(slice10, "Blackberry")
	fmt.Println("slice10 ", slice10)

	slice9[2] = "Nanana"
	fmt.Println(slice9, "<- slice9 slice10 ->", slice10) //As now you can see the underlying array for slice9 and slice10 are different. Comment line 77 to see the change.

	//Appending 2 slices together
	slice11 := append(slice10, slice9...)
	fmt.Println("slice11", slice11)

	//Iterating over slices
	for index, value := range slice11 {
		fmt.Println(index, "->", value)
	}
	/*Since a slice is a collection, we can iterate over the elements. Go has a special keyword
	called range that we use in conjunction with the keyword for to iterate over slices.
	range returns two value index and value at the slice
	*/

	slice12 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for index, value := range slice12 {
		fmt.Println("Index ", index, " Value ", value, " Value Add ", &value, " Element Add ", &slice12[index])
	}
	// As shown the Value Add remains the same because range provides the copy of the values.

	for _, value := range slice12 {
		fmt.Println(" Value ", value)
	}
	//To ignore the index we use blank '_' identifier.

	for i := 0; i < len(slice10); i++ {
		fmt.Println(slice10[i])
	}
	//Iterating the slice using traditional for loop
	/* There are two special built-in functions called len and cap that work with arrays, slices, and channels. For slices, the len function returns the length of the slice, and
	the cap function returns the capacity*/

	//Multidimensional Slices
	slice13 := [][]int{{1}, {2, 3}, {4, 5}}
	fmt.Println(slice13)
	slice13[0] = append(slice13[0], 6)
	fmt.Println(slice13)

	fmt.Println(foo2(slice12)) //Passing the slice in the function doesn't require pointers.
	//New copy of slice is created but the underlying array is same.

}

func foo2(slice []int) []int {
	return slice
}
