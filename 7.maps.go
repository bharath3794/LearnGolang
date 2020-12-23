/*
Reference Links
https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
*/

package main

import (
	"fmt"
)

func main(){
// Declare Map. Still not created and set this varibale map1 to zero value of map = nil
	var map1 map[string]int
	fmt.Printf("map1 = %#v, bool value of (map1 == nil) is %v\n", map1, map1==nil)
// Declared map2 and created using make() function. 
// Now the variable map2 is not set to zero value of map i.e. nil
	var map2 = make(map[string]int)
	fmt.Printf("map2 = %#v, bool value of (map2 == nil) is %v\n", map2, map2==nil)
// Similar to above but short-variable declaration
	map3 := make(map[string]int)
	fmt.Printf("map3 = %#v, bool value of (map3 == nil) is %v\n", map3, map3==nil)

	// Map Literals
// Standard Declaration for map literal
	var map4 = map[string]int{"a": 1, "b": 2}
	fmt.Println("map4 =", map4)
// Short variable declaration for map literal
	map5 := map[string]int{"a": 1, "b": 2}
	fmt.Println("map5 =", map5)

// Map Literal with no values added;
// This declaration is equivalent to map3 := make(map[string]int)
	map6 := map[string]int{}
	fmt.Println("map6 =", map6)
	fmt.Printf("Go Representation map6 = %#v, bool value of (map6 == nil) is %v\n", map6, map6==nil)


	map7 := map[string]int{"a": 1, "b": 2}
// key "c" is still not added to map; See what happens if we access the "c" key
// If key is not available in map it sets the value of the key to its zero value of the value Type
// To know if key is present in map or not, we use "ok" variable which is a bool value
	value, ok := map7["c"]
	fmt.Printf("The value of key 'c' in map7=%v is %v\n", map7, value)
	fmt.Println("Is key 'c' present in map7?", ok)
// Deleting a key/value pair
	delete(map7, "b")
	fmt.Println("Key 'b' deleted and now map7 =", map7)


	// Accessing all key/value pairs in a map
	map8 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("-------------------------------------")
	for key, value := range map8{
		fmt.Println("key =", key, "value =", value)
	}
// If we just want keys, we can use blank identifier '_' or 
// just use "key" in for loop with out value.
	fmt.Println("-------------------------------------")
	for key := range map8{ // This is equivalent to for key, _ := range map8
		fmt.Println("key =", key)
	}
// You can't skip key if you just need value. So we should use blank identifier to just print value
	fmt.Println("-------------------------------------")
	for _, value := range map8{ // This is equivalent to for key, _ := range map8
		fmt.Println("Value =", value)
	}

	fmt.Println("-------------------------------------")
// Check if a key is present in the map and do some operations if present
// In GO, an if statement can contain both initialization and condition
// Below, if 'c' is present in map8, map8["c"] will be loaded to val and ok will be true
// If 'c' is not available in map8, ok would be false and val would be assigned to zero value of 
// type of the values in map8. As 'c' is available in map8, ok would be true and if block
// would be executed otherwise (i.e. if ok=false) 'else' block would be executed 
	if val, ok := map8["c"]; ok{
		fmt.Printf("'c' is present in map8=%v and it's value is %v\n", map8, val)
	} else {
		fmt.Println("'c' is not present in map8=%v\n", map8)
	}
}