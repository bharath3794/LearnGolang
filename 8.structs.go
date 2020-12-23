package main

import (
	"fmt"
)

// Declaring a struct variable. 
// If we need multiple struct variables we need to declare multiple struct varibales as below 
// for each struct variable we want
var mystruct1 struct{
	x int
	y string
	z bool
}

// Creating a struct type
// By creating a new struct type we can create multiple varibales based on this struct type
type mystruct struct{
	x int
	y string
	z bool
}

// Adding a struct as a field on another custom-defined type
// Accessing elements in nested structs is a bit tedious
// So we should use a different concept in GO called Anonymous struct fields
type nestedStruct struct{
	a int
	b bool
	c mystruct
}

// Anonymous struct fields
// we can now directly access the elements of inner struct (mystruct)
// as if they are fields of outer struct (anonyStruct i.e. anonyStruct.x, anonyStruct.y etc.;)
type anonyStruct struct{
	a int
	b bool
	mystruct
}


func main(){
	mystruct1.x = 1
	mystruct1.y = "Bharath"
	mystruct1.z = true
	fmt.Printf("mystruct1 = %v\nGo Representation %#[1]v\n", mystruct1)

// Accessing struct variable elements
	fmt.Println("mystruct1.x =", mystruct1.x, "; mystruct1.y =", mystruct1.y, 
		"; mystruct1.z =", mystruct1.z)
	fmt.Println("--------------------------------")
// Declaring a varibale on custom-defined type mystruct and using struct literals
// In declaration like below, the order of field values are important and
// should be matched with order of custom-defined type fields
// This type of declaration of field values is only good if you know value of every field beforehandedly
// Otherwise, if you only want to set some field values initially 
// its better to use mystruct{x:1, y:"Bharath", z:true} syntax
	mystruct2 := mystruct{1, "Bharath", true} 
	fmt.Printf("mystruct2 = %v\nGo Representation %#[1]v\n", mystruct2)

	// Accessing struct variable elements
	fmt.Println("mystruct2.x =", mystruct2.x, "; mystruct2.y =", mystruct2.y, 
		"; mystruct2.z =", mystruct2.z)
	fmt.Println("--------------------------------")
// Other Method to declare field values using struct literal
// In declaration like below, the order of field values is not important
// Mixture of both declarations like this (mystruct{1, y: "Bharath", z: true}) would not work
	mystruct21 := mystruct{y: "Bharath", x: 1, z: true}
	fmt.Printf("mystruct21 = %v\nGo Representation %#[1]v\n", mystruct21)
	fmt.Println("--------------------------------")

// Leaving some field values in struct literal while declaration like below
// This syntax works mystruct{x: 1, y: "Bharath"}
// but this doesn't work mystruct{1, "B"} or mystruct{1, z:true} etc.;
	mystruct22 := mystruct{x: 1}
	fmt.Printf("mystruct22 = %v\nGo Representation %#[1]v\n", mystruct22)
	fmt.Println("--------------------------------")

// our varibale mystruct2 of type mystruct is passed as argument to showInfo() function
	mystruct3 := showInfo(mystruct2)
	fmt.Println("Passed mystruct2 argument is returned and stored in mystruct3 and is =",mystruct3)
	fmt.Println("--------------------------------")

	mystruct4 := mystruct{1, "Bharath", true}
	fmt.Println("Before calling modified(&mystruct4), mystruct4 =", mystruct4)
	modified(&mystruct4)
	fmt.Println("After calling modified(&mystruct4), mystruct4 =", mystruct4)
	fmt.Println("--------------------------------")

// Declaring nestedstruct varibale using struct literal 
// If we want we can also omit declaring some field values and those will be set to their zero values
// For ex: nestedStruct{a:1} = {1, false, {0, "", false}}
	nestedStruct1 := nestedStruct{1, true, mystruct{1, "Bharath", true}}
	fmt.Printf("nestedStruct1 = %v\nGo Representation %#[1]v\n", nestedStruct1)
// Accessing or setting elements on nested structs is a bit tedious
// So we should use a different concept in GO called Anonymous struct fields
	fmt.Println("nestedStruct1.a =", nestedStruct1.a, "nestedStruct1.b =", nestedStruct1.b, 
		"nestedStruct1.c =", nestedStruct1.c)
// If we further access nestedStruct1.c values which is itslef a struct
	fmt.Println("nestedStruct1.c.x =", nestedStruct1.c.x, "nestedStruct1.c.y =", nestedStruct1.c.y,
	"nestedStruct1.c.z =", nestedStruct1.c.z)
	fmt.Println("--------------------------------")

// Declaring anonyStruct variable using a struct literal
	anonyStruct1 := anonyStruct{1, true, mystruct{1, "Bharath", true}}
	fmt.Printf("anonyStruct1 = %v\nGo Representation %#[1]v\n", anonyStruct1)
// Accessing or setting elements on nested structs is now easy with anonymous struct fields
	fmt.Println("anonyStruct1.a =", anonyStruct1.a, "anonyStruct1.b =", anonyStruct1.b, 
		"anonyStruct1.mystruct =", anonyStruct1.mystruct)
// If we want to further access anonyStruct1.mystruct values which is itslef a struct
// we can now directly access the elements of inner struct (mystruct)
// as if they are fields of outer struct (anonyStruct i.e. anonyStruct.x, anonyStruct.y etc.;)
	fmt.Println("anonyStruct1.x =", anonyStruct1.x, "anonyStruct1.y =", anonyStruct1.y,
	"anonyStruct1.z =", anonyStruct1.z)
	fmt.Println("--------------------------------")

// How to copy structs type of one variable to other varibale and How it internally works?
// To understand this, go to copyStructs() function where we performed both normal copy and
// deep copy
	copyStructs()
	fmt.Println("--------------------------------")
// How to assign a struct field in map
	fmt.Println("One way of Assigning a struct field in map")
	temp1 := map[string]mystruct{"item1":mystruct{1, "bharath", true}}
	fmt.Printf("Before assigning a struct field in map, temp1 = %#v\n", temp1)
// The below code will return an error saying: 
// Error: cannot assign to struct field temp1["item1"].z in map

//	temp1["item1"].z = false

// One way of doing it is:
	fmt.Println("By checking if key is in map or not")
	if temp2, ok := temp1["item1"]; ok {
		temp2.z = false
		temp1["item1"] = temp2
	}
	fmt.Printf("After assigning a struct field in map, temp1 = %#v\n", temp1)
	fmt.Println("--------------------------------")
// The other way of doing it is: make your map a map of pointers to structs.
// The disadvantage with this method is you can't see struct values when printed, instead we
// see pointer address to this struct.
	fmt.Println("Other way of Assigning a struct field in map")
	fmt.Println("Using pointer to structs in a map")
	temp2 := map[string]*mystruct{"item1": &mystruct{1, "bharath", true}}
	fmt.Printf("Before assigning a struct field in map, temp2 = %#v\n", temp2)
	fmt.Println("Assigning to a struct field in a map, temp2['item1'].z = false")
	temp2["item1"].z = false
	fmt.Printf("After assigning a struct field in map, temp2 = %#v\n", temp2)
// If in map[string]mystruct, mystruct itself is a map, bool = mystruct{map[string]int, bool};
// We can directly assign map values in mystruct, and no need of doing as above for bool
	fmt.Println("--------------------------------")
	type exstruct struct {
		m map[string]int
		b bool
	} 
	temp3 := map[string]exstruct{"v1":exstruct{map[string]int{"v2":0}, false}}
	fmt.Printf("Before assigning to a map value in a struct field in map, temp3 = %#v\n", temp3)
// From above discussions, we can't perform assignment of temp3["v1"].b = true and lead to error
// But we can directly assign to map values in struct;
// For ex, temp3["v1"].m["v2"] = 1 is perfectly valid
	fmt.Println("Assigning to a map value in a struct field which is itself in map,\ntemp3['v1'].m['v2'] = 1")
	temp3["v1"].m["v2"] = 1
	fmt.Printf("After assigning to a map value in a struct field in map, temp3 = %#v\n", temp3)


}

// function on custom-defined type
// i.e. custome-defined type is passed as argument to this function
func showInfo(m mystruct) mystruct {
	fmt.Println("mystruct2 passed as argument to showInfo() and is =", m)
	return m
}

// function 
func modified(m *mystruct){
// we don't need to use pointer (i.e. *m) to set x, y, z fields. 
// If we pass a pointer, Go is able to understand directly while accessing fields of any struct
// as if it is a direct access rather a pointer access.
// But even here if we want to use pointer to access those fields, we need to use
// (*m).x, (*m).y etc., All these are equivalent to accessing fields directly using m.x, m.y etc.;
// We shouldn't use *m.x, *m.y etc. because Go considers it as m.x has a pointer and 
// then looks for pointer value in m.x which is not the case in our example.
	m.x = 2
	m.y = "Talluri"
	m.z = false
	fmt.Println("In modified() function and is =",m)
}


func copyStructs(){
	type point2d struct{
		x float64
		y float64
	}

	type points struct{
		p []point2d
	}

	v := points{[]point2d{point2d{2, 3}, point2d{12, 30}, point2d{40, 50}, point2d{5, 1}, 
	       point2d{12, 10}, point2d{3, 4}, point2d{6, 2}}}
	fmt.Printf("Go Representation = %#v\n", v)
	fmt.Println("v =", v)
	fmt.Println("------Performing Normal Copy (normalCopy := v)------")
	normalCopy := v
	fmt.Printf("normalCopy = %v\n", normalCopy)
	fmt.Println("(&normalCopy == &v) =", &normalCopy == &v)
	fmt.Println("(&normalCopy.p == &v.p) =", &normalCopy.p == &v.p)
	fmt.Println("(&normalCopy.p[0] == &v.p[0]) =", &normalCopy.p[0] == &v.p[0])
	fmt.Println("------Performing Deep Copy using copy() function------")
    temp := make([]point2d, len(v.p))
    copy(temp, v.p)
	deepCopy := points{temp}
	fmt.Printf("deepCopy = %v\n", deepCopy)
	fmt.Println("(&deepCopy == &v) =", &deepCopy == &v)
	fmt.Println("(&deepCopy.p == &v.p) =", &deepCopy.p == &v.p)
	fmt.Println("(&deepCopy.p[0] == &v.p[0]) =", &deepCopy.p[0] == &v.p[0])
}