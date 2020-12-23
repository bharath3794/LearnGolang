/*
Reference Links:
https://github.com/golang/go/wiki/SliceTricks

https://www.calhoun.io/why-are-slices-sometimes-altered-when-passed-by-value-in-go/
*/

package main

import(
	"fmt"
	"math/rand"
	"reflect"
)

func main(){
// Declaring Slice. Just creates it but don't add zero values to individual elements.
// For Zero Values on individual elements, 
// we need to create a slice with certain length using make() function
// Unlike arrays, If we just declared a slice, the slice variable itself is assigned with a zero value nil
// The len() on slice which has zero value 'nil' is zero; Here, len(slice1) = 0
	var slice1 []int
	fmt.Println("slice1 =", slice1)
	fmt.Printf("slice1  = %#v, len(slice1) = %v\n", slice1, len(slice1))
	fmt.Println("bool value of (slice1 == nil) is", slice1 == nil)
// Check if the variable has a zero value using IsZero function in reflect package
	fmt.Println("Check if a variable has a zero value using reflect package")
	temp := reflect.ValueOf(slice1)
	fmt.Println("Is slice1 has zero value of type slice?", temp.IsZero())
// Creating Slice using make() function; also adds Zero Values to given length of slice
	var slice2 = make([]int, 5) //equivalent to var slice2 []int; slice2 = make([]int, 5)
	fmt.Println("slice2 =", slice2)
// Short Variable Declaration and make
	slice3 := make([]int, 5)
	fmt.Println("slice3 =", slice3)
// Slice Literal using var declaration
	var slice4 = []int{1, 2, 3, 4, 5}
	fmt.Println("slice4 =", slice4)
// Slice Literal using Short-Variable declaration
	slice5 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice5 =", slice5)

// append() function on slice
// Create an empty slice
	slice6 := []int{} // This declaration creates slice6 and the zero vlaue of slice6 is not nil
	fmt.Println("slice6 =", slice6)
	fmt.Printf("slice6 = %#v\n", slice6)
	fmt.Println("bool value of (slice6 == nil) is", slice6 == nil)
// We assigned the returned value from append() to same slice variable which we passed to it
	slice6 = append(slice6, 1)
	slice6 = append(slice6, 2, 3, 4, 5)
	fmt.Println("Appended slice6 =", slice6)

// Append two slices. Here, temp2 is appended to temp1. The three dots (ex:temp2...) indicates
// that temp2 is a slice of int values
	fmt.Println("Appending two Slices")
	temp1 := []int{1, 2, 3}
	temp2 := []int{4, 5}
	fmt.Printf("temp2 = %v is appended to temp1 = %v. ", temp2, temp1)
	temp1 = append(temp1, temp2...) // equivalent to temp1 = append(temp1, []int{4, 5}...)
	fmt.Printf("Now temp1 = %v\n", temp1)
// Caution: Different behaviour while appending to different slice variable
// Create an empty slice
	slice := []int{}
	fmt.Println("slice =", slice)
// We assigned the returned value from append() to different slice variable 
// i.e. not which we passed to it
	slice7 := append(slice, 1)
	slice7 = append(slice, 2, 3, 4, 5)
	fmt.Println("Appended slice =", slice)
	fmt.Println("Appended to different slice =", slice7)

// Convert an array to a slice easy way
	fmt.Println("Converting an array to a slice")
	arr := [3]int{1, 2, 3}
	slice8 := arr[:]
	fmt.Printf("arr = %v of type %[1]T is converted to slice8 = %v of type %[2]T\n", arr, slice8)

// Create a random number slice with Integer values of certain range
// 
	fmt.Println("Generating slice using math/rand package")
	rnd := rand.Perm(10)
	fmt.Printf("rnd = %v of type %[1]T\n", rnd)

// In Go, function parameters are passed by value (i.e. the copy of it)
// When a slice is passed to a function, even though it's passed by value, 
// the pointer variable will refer to the same underlying array.
// Hence when a slice is passed to a function as parameter, changes made inside the function 
// are visible outside the function too.
// It can only modify the existing elements and reflect it in original slice;
// Can't increase the length and capacity of the slice thus appending elements are not reflected
// in original slice 
	fmt.Println("Slice passed to modifySlice() function")
	a := []int{1, 2, 3}
	fmt.Println("Passed Slice, a =", a)
	modifySlice(a)
	fmt.Println("After calling modifySlice(), a =", a)
// cap(slice) is used to know the capacity of the slice 
	fmt.Println("Capacity of a Slice")
	temp3 := []int{1, 2, 3}
	fmt.Println("temp3 =", temp3)
	fmt.Println("cap(temp3) =", cap(temp3))
// Creating a copy of slice
	fmt.Println("Create Copy of a Slice")
	temp4 := []int{1, 2, 3, 4, 5}
	temp4Copy := append([]int(nil), temp4...)
	fmt.Println("temp4 =", temp4, "temp4Copy =", temp4Copy)


// assigning two depenedent values to two variables while modifying them doesn't work in GO
/* For ex, 
			slice = [0, 1, 2, 3, 4, 5, 6]
			ele, slice = slice[3], append(slice[:3], slice[4:]...)
We think that the above command should give ele = 3, slice = [0, 1, 2, 4, 5, 6] with the thought that
the operations of single line assignments with multiple variables will be processed from left to right.
i.e. In above statement first slice[i] will be assigned to ele before performing append() and assigning
it to slice. This is pythonic way of thinking because in python this works exactly as discussed above.
* But in Go, this wouldn't work and infact the above statement gives the output as:
		ele = 4, slice = [0, 1, 2, 4, 5, 6]
Similar code in python works and is given below:
			lst = [0, 1, 2, 3, 4, 5, 6]
			ele, lst = lst[3], lst[:3]+lst[4:]
	Result: ele = 3, lst = [0, 1, 2, 4, 5, 6]
*/
	slice = []int{0, 1, 2, 3, 4, 5, 6}
	i := 3 // delete element at index 3
	fmt.Printf("----------Before Deleting the element at index %v----------\n", i)
	fmt.Println("We think ele should be =", slice[i], "as slice =", slice)
	fmt.Printf("Performing ele, slice := slice[%v], append(slice[:%[1]v], slice[%[1]v+1:]...)\n", i)
	ele, slice := slice[i], append(slice[:i], slice[i+1:]...)
	fmt.Printf("----------After Deleting the element at index %v----------\n", i)
	fmt.Println("But strangely ele =", ele, "and slice =", slice)



}

// Modifications made only to existing elements of the slice are reflected in original slice;
// appending elements are not reflected in original slice
func modifySlice(c []int){
	c[2] = 4
	c[0] = 0
	c = append(c, 5)
}