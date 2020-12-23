package main

import (
	"fmt"
)

func main() {
	a, b, c, d := 3, 5.3, "Bharath", true
	variadic(a, b, "Venkatasiva", "Bharath", "Talluri")
	variadic(a, b)
	variadicTypes(a, b, c, d)

	// Passing Slice to a variadic function which only takes int values
	slice := []int{1, 2, 3, 4}
	variadicSlice(slice...)
	/* But passing array will not work and raises panic */
	// arr := [4]int{1, 2, 3, 4}
	// variadicSlice(arr...)
}

// non-variadic arguments x1, y1 are compulsory and 
// variadic arguments t1 are not compulsory and are stored as slice
func variadic(x1 int, y1 float64, t1 ...string){
	fmt.Println("int Value x1 =", x1, ", float value y1 =", y1, ", string values t1 =", t1)
}

// This Variadic function can take all types as parameters because of empty interface interface{}
// Empty interface has no methods defined in it
// So in Go, every type including built-in and custom defined types satisfy this property of no methods
// So every type can be part of this empty interface.
func variadicTypes(types ...interface{}) {
	fmt.Println("Passed Arguments to this Variadic function =", types)
}

// This variadic function only takes multiple int values as arguments
// But we can pass a slice of int values as argument to this variadic function directly
// This function will understand this is a slice of int values,
// eventhough it is not an int value directly
func variadicSlice(numbers ...int){
	fmt.Println("Passed Slice of int values =", numbers)
}