package main

import(
	"fmt"
	"os"
	"strconv"
	"log"
)

func main(){
	fmt.Println(os.Args[1:])
	for i, v := range os.Args[1:]{
		fmt.Printf("%v. value=%v of type %[2]T\n", i, v)
		vFloat, err := strconv.ParseFloat(v, 64)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Printf("%v. value=%v converted to type %[2]T\n", i, vFloat)
		vInt, _ := strconv.ParseInt(v, 10, 32)
		fmt.Printf("%v. value=%v converted to type %[2]T\n", i, vInt)
	}

}