// Also Refer to these links
/*
https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
https://research.swtch.com/interfaces

https://stackoverflow.com/questions/45652560/interfaces-and-pointer-receivers
https://play.golang.org/p/Y0fJcAISw1

*/

package main

import (
	"fmt"
)

type googleSearch struct{
	browser string
}

type bingSearch struct{
	browser string
}

type yahooSearch struct{
	browser string
}

type web interface{
	startSearch(string)
}



func main() {
	items := []string{"Chimmiribanda", "Martur"}
// "tools" is of type web whose underlying type is interface.
// Variable "tools" can now be assigned with values that belongs to types which have methods defined in
// the web interface. Here, web interface has startSearch(string) method and all defined
	var tools web

	tools = googleSearch{"Chrome"}
	// Passing tools directly to search() function
	search(tools, items)
	// Passing tools as a pointer (&tools) to search() function
	// search(&tools, items) //results in panic: *web is pointer to interface, not interface
	fmt.Println("-----------------------------------------------------------")
	tools = bingSearch{"Edge"}
	search(tools, items)
	fmt.Println("-----------------------------------------------------------")
	
// tools1 is of type web whose underlying type is interface.
// Variable tools1 can now be assigned with values that belongs to types which have methods defined in
// the web interface. Here, web interface has startSearch(string) method and all defined
// 3 types googleSearch, bingSearch, yahooSearch have this startSearch(string) method.
// Both googleSearch, bingSearch types are defined directly by value as receiver parameter 
// on startSearch(string) method, where as yahooSearch is defined by a pointer receiver on this method
// So we can assign any value of these 3 types to variable tools1 web interface type
	var tools1 web

// When the Receiver Parameter of a method is a direct value instead of a pointer receiver, 
// we can send both direct value and a pointer receiver.

// Below two declarations(tools1 = googleSearch{"Chrome"} and tools1 = &googleSearch{"Chrome"}) 
// works without any error
	tools1 = googleSearch{"Chrome"}
	fmt.Printf("Passing tools1 = %#v directly to startSearch() function\n", tools1)
	for _, item := range items{
		tools1.startSearch(item)
	}

	tools1 = &googleSearch{"Chrome"}
	fmt.Printf("Passing tools1 = %#v as a pointer (tools1=%[1]v) to startSearch() function\n", tools1)
	for _, item := range items{
		tools1.startSearch(item)
	}
	fmt.Println("-----------------------------------------------------------")
	tools1 = &yahooSearch{"Mozilla"}
	fmt.Printf("Passing tools1 = %#v as a pointer (tools1=%[1]v) to startSearch() function\n", tools1)
	for _, item := range items{
		tools1.startSearch(item)
	}
	fmt.Println("-----------------------------------------------------------")

// The below method will not work because of need for pointer receiver to startSearch Method
	// tools1 = yahooSearch{"Mozilla"}
	// fmt.Printf("Passing tools1 = %#v directly to startSearch() function\n", tools1)
	// for _, item := range items{
	// 	tools1.startSearch(item) //panic: startSearch method has pointer receiver
	// }


// This will not work either
// Because, when we call (&tools1).startSearch(item), it is expecting a pointer to yahooSearch type
// but not a pointer to interface type. Here, tools1 is of web interface type and &tools1 would be a
// pointer to that interface not either the interface itself or a pointer to yahooSearch type
// If we declare a a variable as yahooSearch type and use the pointer as below would work on it
	// tools1 = yahooSearch{"Mozilla"}
	// fmt.Printf("Passing tools1 = %#v as a pointer (tools1=%[1]v) to startSearch() function\n", tools1)
	// for _, item := range items{
	// 	(&tools1).startSearch(item) // panic: (type *web is pointer to interface, not interface)
	// }

// Here, we are declaring tools2 as yahooSearch type instead of web interface type. So this should work
	tools2 := yahooSearch{"Mozilla"}
	fmt.Printf("tools2 = %#v of type = %[1]T\n", tools2)
	fmt.Printf("Passing tools2 as a pointer (&tools2=%v) to search() function\n", &tools2)
	for _, item := range items{
		(&tools2).startSearch(item)
	}
	fmt.Println("-----------------------------------------------------------")

}

func (g googleSearch) startSearch(term string){
	fmt.Printf("Browser: %v; Search Engine: Google; Search term: %v\n", g.browser, term)
}


func (b bingSearch) startSearch(term string){
	fmt.Printf("Browser: %v; Search Engine: Bing; Search term: %v\n", b.browser, term)
}

// Receiver parameter of type yahooSearch should be send as a pointer
func (y *yahooSearch) startSearch(term string){
	fmt.Printf("Browser: %v; Search Engine: Yahoo; Search term: %v\n", y.browser, term)
}

func search(w web, items []string){
	for _, item := range items{
		w.startSearch(item)
	}
}


