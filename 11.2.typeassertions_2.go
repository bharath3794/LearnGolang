package main

import "fmt"

func main(){
// interfaceSlice would be a slice with each element is of type empty interface = interface{}
// Empty interface has no methods defined in it
// So in Go, every type including built-in and custom defined types satisfy this property of no methods
// So every type can be part of this empty interface slice.
// The disadvantage of values of type empty interface would be we can't call any method on them unless
// we perform type assertion and convert to its underlying type which looses the 
// whole purpose of using interface concept
// Thus, these values of empty interface type are more useful for just storing elements of 
// multiple data types in a single slice or any built-in data structure.
// The only operation we can perform on these empty interface elements is we are able to print them
// To perform the methods or other operations based on their underlying type of these 
// empty interface elements, we need to first perform type assertion and 
// then call their respective methods
	interfaceSlice := []interface{}{1, "Bharath", true, []int{2, 3, 4}}
	fmt.Printf("interfaceSlice = %v\nGO Representation = %#[1]v\n",interfaceSlice)
	fmt.Println("----------------------------------------------")
// Accessing elements using "range" in for loop
	fmt.Println("Accessing elements using 'range' in for loop")
	for i, v := range interfaceSlice{
		fmt.Println("Index=", i, "\tValue=", v)
	}

	fmt.Println("----------------------------------------------")


//Accessing index wise elements using for loop
	fmt.Println("Accessing index wise elements using for loop")
	for i:=0; i<len(interfaceSlice); i++{
		fmt.Println("Index=", i, "\tValue=", interfaceSlice[i])
	}

	fmt.Println("----------------------------------------------")

// How to access elements of slice []int{2, 3, 4} at index 3 i.e. interfaceSlice[3]
	fmt.Println("How to access elements of slice []int{2, 3, 4} at index 3 i.e. interfaceSlice[3]")
	fmt.Printf("interfaceSlice[3] = %v, GO Representation = %#[1]v, Type = %[1]T\n", interfaceSlice[3])

// Accessing element at index 0 of []int{2, 3, 4} slice.
// But this would not work because the type of slice []int{2, 3, 4} is actually interface even though
// above it is mentioned as []int. This is because we are adding slice []int{2, 3, 4} to a slice whose
// elements are of type interface{} i.e. empty interface. Empty interface has no methods defined in it
// So in Go, every type including built-in and custom defined types satisfy this property of no methods
// So every type can be part of this empty interface slice.
// To access element at index 0 of []int{2, 3, 4} slice,
// we need to first convert it to original type []int slice using type assertion
	// fmt.Println(interfaceSlice[3][0]) // panic: (type interface {} does not support indexing) 


// We can't even range over interface type. Below code results in
// panic: cannot range over interfaceSlice[3] (type interface {})
	// for _, v := range interfaceSlice[3]{ 
	// 	fmt.Println(v)
	// }


// Converting []int{2, 3, 4} interface type to original type []int slice
	fmt.Println("Type Assertion: Converting []int{2, 3, 4} interface type to original type []int slice")
	slice := interfaceSlice[3].([]int)
	fmt.Printf("slice = %v, GO Representation = %#[1]v, Type = %[1]T\n", slice)
	fmt.Printf("slice[0] = %v\n", slice[0])

// No need of storing the result from type assertion to a variable if you don't need it.
// We can directly use the result
	fmt.Println("Directly using the type assertion result")
	fmt.Println("interfaceSlice[3].([]int)[0] =", interfaceSlice[3].([]int)[0])
fmt.Println("Accessing type assertion slice=interfaceSlice[3].([]int) elements using 'range' in for loop")
	for i, v := range slice{
		fmt.Println("Index=", i, "\tValue=", v)
	}

}