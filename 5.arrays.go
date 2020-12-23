package main

import(
	"fmt"
)

func main(){
	// Below Declaration Assigns Zero Values to elements of array arr1
	var arr1 [3]int
	fmt.Println("arr1 =", arr1)
	// Short-Variable Declaration; Also assigns Zero Values to elements of array arr2
	// arr2 := [3]int // This doesn't work and leads to panic
	arr2 := [3]int{}
	fmt.Println("arr2 =", arr2)
	// Array Literals using Short-variable Declaration
	arr3 := [3]int{1, 2}
	fmt.Println("arr3 =", arr3)
	// With var declaration
	var arr4 =  [3]int{1, 2, 3}
	fmt.Println("arr4 =", arr4)


	// Accessing Array elements using range
	fmt.Println("Accessing Array elements using range")
	for indx, value := range arr4{
		fmt.Println(indx, value)
	}

	// Accessing array elements using index
	fmt.Println("Accessing array elements using index")
	for i:=0; i<len(arr4); i++{
		fmt.Println(arr4[i])
	}
	// Accessing array elements using index in for loop with break
	fmt.Println("Accessing array elements using index in for loop with break")
	j := 0
	for true{
		if j >= len(arr4){
			break
		}
		fmt.Println(arr4[j])
		j++
	}

	// Accessing Array elements using range and if you just needs index but not its value
	fmt.Println("Accessing only index of Array elements using range")
	for indx := range arr4{
		fmt.Println(indx)
	}
	// Accessing Array elements using range and if you just needs value but not its index
	fmt.Println("Accessing only value of Array elements using range")
	for _, value := range arr4{
		fmt.Println(value)
	}
	//In Go, function parameters are passed by value (i.e. the copy of it)
	// Array 'a' modified in modifyArray() doesn't change original array 'a'
	fmt.Println("Array 'a' passed to modifyArray() function")
	a := [3]int{1, 2, 3}
	fmt.Println("Passed Array, a =", a)
	modifyArray(a)
	fmt.Println("After calling modifyArray(), a =", a)
}


// Modifications made to array in below function will not affect original array and are not reflected
func modifyArray(c [3]int){
	c[2] = 4
}