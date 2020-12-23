package main

import (
	"fmt"
	"reflect"
)


func main(){
	a := []interface{}{1, true, "bharath"}
	typeSwitch(a)
	b := []bool{true, false, true}
	typeSwitch(b)
	fmt.Println("-----------Type Switch Use Case Example-----------")
	v1 := []interface{}{1, true, "bharath", []interface{}{2, false, "Talluri", 
		 []interface{}{3, true, "Venkata"}}}
	fmt.Printf("v1 = %v, \nGo Representation = %#[1]v\n", v1)
	fmt.Println("reflect.TypeOf(a[3]) == reflect.TypeOf([]interface{}{}):", 
				 reflect.TypeOf(v1[3]) == reflect.TypeOf([]interface{}{}))
	rslt := embedAccess(v1, 3, 3)
	fmt.Printf("embedAccess(v, 3, 3) = %v, of type %[1]T\n", rslt)

 }



func typeSwitch(x interface{}) {
	switch v := x.(type) {
	case int:
	    fmt.Println("int:", v)
	case float64:
	    fmt.Println("float64:", v)
	case string:
	    fmt.Println("string:", v)
	case bool:
	    fmt.Println("bool:", v)
	case []int:
		fmt.Println("[]int:", v)
	case []float64:
		fmt.Println("[]float64:", v)
	case []string:
		fmt.Println("[]string:", v)
	case []bool:
		fmt.Println("[]bool:", v)
	case []interface{}:
		fmt.Println("[]interface{}:", v)
	default:
	    fmt.Println("unknown")
	} 
}

func typeSwitchSample(x interface{}) []interface{}{
	var value []interface{}
	switch v := x.(type) {
	case []interface{}:
		value = v
	default:
	    fmt.Println("unknown")
	}
	return value
}

func embedAccess(root []interface{}, items ...int) []interface{}{
	iterVal := root
	for _, val := range items {
		iterVal = typeSwitchSample(iterVal[val])
	}
	return iterVal
}