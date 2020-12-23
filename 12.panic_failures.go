package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
// The defer keyword is defined before a function or method call with in some function or method
// It delays the execution of it till end of that some function or some method if they end gracefully, 
// otherwise if they lead to panic the function or method before which defer is used are executed
// and then conitnued to panic

// Here, defer reportPanic() waits for main() function to end gracefully and then it calls reportPanic()
// or it waits for a panic in below statements or function or method calls and delay the execution
// of panic till the execution of reportPanic() is completed
	defer reportPanic()
	// panic("Random Panic")
	readDirectory("waste1")
}



func readDirectory(folder string){
	fmt.Println(folder)
	contents, err := ioutil.ReadDir(folder)
	if err != nil{
		panic(err)
	}
	if len(contents) == 0{
		fmt.Println("--------Folder", folder, "is empty--------")
	}
	for _, v := range contents{
		filePath := folder+"/"+ v.Name()
		if v.IsDir() {
			readDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}



func reportPanic(){
// recover() function returns the value in the panic() i.e. panic() takes any type because it is an
// empty interface where as recover() returns any type because it returns empty interface 

// Here, if panic("Random error") is the panic() which takes the string "Random error"
// recover() function actually returns this "Random error" as an empty inteface type whose
// underlying type is a string. So p = interface{}{"Random error"}
	p := recover()
}
}
// As p is an error value converted to an empty interface interface{} type.
// Thus, the underlying type of p is error. We can't call the Error() method on this p even though
// its underlying type is error. We need to first convert this empty interface type of p to its 
// underlying type error to benefit from calling the functions like Error() on p
	if p == nil{
		return
	} else {
		// So we can directly print p like below or we can perform type assertion and then use it
		// for printing as shown in below comment
		fmt.Println("panic:", p) // equivalent to fmt.Println(p.(error)) or fmt.Println((p.(error)).Error())
		// Even this code is valid and prints same as above
		/* --------Code Block----------
		err, ok := p.(error)
		if ok{
			fmt.Println(err) // equivalent to fmt.Println(err.Error())
		} else { // means the above p is not of error type i.e. the original panic is not of error type
		// So we need to send p to original panic() for all other types i.e. not errors 
			panic(p)
		}
		*/
	}
}