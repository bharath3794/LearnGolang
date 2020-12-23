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


type web interface{
	startSearch(string)
}

func main() {

// "tools" is of type web whose underlying type is interface.
// Variable "tools" can now be assigned with values that belongs to types which have methods defined in
// the web interface. Here, web interface has startSearch(string) method and all defined
	var tool web
	
	tool = googleSearch{"Chrome"}
	searchTerm := "Chimmiribanda"
	tool.startSearch(searchTerm)

// The below call tool.openIncognito() would not work because tool is of type web Interface
// and the method openIncognito() is not  part of web interface. So we can't call the method
// on interface type that is not part of that particular interface. 
// We need Type Assertions in these situations
	// tool.openIncognito() // panic: type web has no field or method openIncognito

// Type Assertion: As tool is web interface type, we need to first convert it to the type
// which we need to call that particular method on. The type to which we are 
// going to convert needs to be the receiver parameter for that method we need to call on.
// As openIncognito() has a receiver parameter googleSearch type we need to convert this 
// web interface type tool to googleSearch and call the openIncognito() method on this
// converted googleSearch type
	interfaceField := tool.(googleSearch)
	interfaceField.openIncognito()


// Type Assertion Failure results in panic
	var tool1 web = googleSearch{"Chrome"}
// This results in panic: interface conversion: main.web is main.googleSearch, not main.bingSearch
	// test := tool1.(bingSearch)
	// fmt.Println(test)

// To handle these cases and avoid panics, the type assertion method returns 
// an optional second variable which is of type bool. Conventionally it is stored in variable 'ok'
	test, ok := tool1.(bingSearch)
	fmt.Println("test =", test, "ok =", ok)
}


func (g googleSearch) startSearch(term string){
	fmt.Printf("Browser: %v; Search Engine: Google; Search term: %v\n", g.browser, term)
}


func (b bingSearch) startSearch(term string){
	fmt.Printf("Browser: %v; Search Engine: Bing; Search term: %v\n", b.browser, term)
}

// This method openIncognito() is not part of web interface. So we can't call this method
// by using a variable whose type is web interface
func (g googleSearch) openIncognito() {
	fmt.Printf("Browser: %v; Search Engine: Google; Incognito Short-cut key: Ctrl+Shift+n\n", g.browser)
}
