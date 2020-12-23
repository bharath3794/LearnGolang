package main

import (
	"fmt"
	"strings"
)

func main(){
	s := "Hello"
// Printing whole string using Printf
	fmt.Println("-----Using '%s' in Printf()-----")
	fmt.Printf("s = %v; Go Representation = %#[1]v; Type = %[1]T\n", s)
// Length of a string
	fmt.Println("-----Length of a string-----")
	fmt.Println("len(s) =", len(s))
// Accessing a character in string using Index actually prints its ASCII value
	fmt.Println("-----Directly access an index in string; s[0]-----")
	fmt.Println("s =", s, "and s[0] =",s[0],"is the ASCII value of character at index 0 i.e. 'H'")
// How to access a character instead of its ASCII value?
	// One way, if you just need to print it
	fmt.Println("-----Using '%c' in Printf()-----")
	fmt.Printf("s[0] = %c\n", s[0])
	fmt.Println("-----Using string() conversion-----")
	// We can convert it to string using string() function
	fmt.Println("string(s[0]) =", string(s[0]))
	fmt.Println("-----Using slicing-----")
	// We can convert it to string using string() function
	fmt.Println("s[0:1] =", s[0:1])
// Iterating over for loop to access each character
	fmt.Println("-----Iterating over for loop and range-----")
	for _, v := range s{
		fmt.Printf("%c", v)
	}
	fmt.Println()
// String to slice using strings.Split(s string, sep string). It returns a slice of characters
	fmt.Println("-----strings.Split(s string, sep string)-----")
	slice := strings.Split(s, "")
	fmt.Printf("slice = %v, GO Representation = %#[1]v, Type = %[1]T\n", slice)


// All the above methods consider the characters as ASCII characters;
// What if the characters are Unicode?
// Go has a different data type to handle Unicode characters i.e. []rune
/*
* In UTF-8, ASCII characters are single-byte corresponding to the first 128 Unicode characters. 
Strings behave like slices of bytes. A rune is an integer value identifying a Unicode code point.
* Go doesn't really have a character type as such. 
byte is often used for ASCII characters, and rune is used for Unicode characters, 
but they are both just aliases for integer types (uint8 and int32)
*/
	fmt.Println("-----Handling Unicode Characters-----")
	s = "界"
	fmt.Printf("s = %v, Printf(\"%%c\", []rune(s)[0]) = %c\n", s, []rune(s)[0])
	fmt.Println("s =", s, ", Println(string([]rune(s)[0])) =", string([]rune(s)[0]))
	fmt.Println("-----converting string to rune-----")
	s = "Hello, 界"
	fmt.Printf("s = %v; Go Representation = %#[1]v; Type = %[1]T\n", s)
	rune := []rune(s)
	fmt.Printf("s = %v; Go Representation = %#[1]v; Type = %[1]T\n", rune)


// Concatenate each element of []string to a string using strings.Join(s []string, sep string)(string)
	fmt.Println("-----strings.Join(s []string, sep string)(string)-----")
	slice = []string{"H", "e", "l", "l", "o"}
	fmt.Println("slice =", slice, ", strings.Join(slice, "") =", strings.Join(slice, ""))
}