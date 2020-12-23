package main

import "fmt"

type Myvar struct {
	x, y float64
}

func (v1 *Myvar) func1(f float64) {
	v1.x = v1.x * f
	v1.y = v1.y * f
}

func (v2 *Myvar) func2(f float64) (float64, float64) {
	return v2.x * f, v2.y * f
}
func func3(v3 *Myvar, f float64) {
	v3.x = v3.x * f
	v3.y = v3.y * f
}
func func4(v4 *Myvar, f float64) (float64, float64) {
	return v4.x * f, v4.y * f
}

func (v5 Myvar) func5(f float64) {
	v5.x = v5.x * f
	v5.y = v5.y * f
}
func (v6 Myvar) func6(f float64) (float64, float64) {
	return v6.x * f, v6.y * f
}
func func7(v7 Myvar, f float64) {
	v7.x = v7.x * f
	v7.y = v7.y * f
}
func func8(v8 Myvar, f float64) (float64, float64) {
	return v8.x * f, v8.y * f
}


func main() {
	v := Myvar{4, 8}
	fmt.Println("---------------------------------")
	v.func1(2)
	// &v.func1(2) //can't use &v
	fmt.Println(v.x, v.y)
	fmt.Println("")
	fmt.Println(v.func2(2))
	// fmt.Println(&v.func2(2)) // can't use &v
	fmt.Println("---------------------------------")
	// func3(v, 2) //can't use v
	func3(&v, 2)
	fmt.Println(v.x, v.y)
	fmt.Println("")
	// fmt.Println(func4(v, 2)) // can't use v
	fmt.Println(func4(&v, 2))
	fmt.Println("---------------------------------")
	// &v.func5(2) //can't use &v
	v.func5(2)
	fmt.Println(v.x, v.y)
	fmt.Println("")
	// fmt.Println(func6(&v, 2)) //can't use &v
	fmt.Println(v.func6(2))
	fmt.Println("---------------------------------")
	func7(v, 2)
	fmt.Println(v.x, v.y)
	fmt.Println("")
	fmt.Println(func8(v, 2))
	fmt.Println("---------------------------------")
	p := &Myvar{4, 8}
	p.func1(2)
	fmt.Println(p.x, p.y)
	fmt.Println("")
	fmt.Println(v.x, v.y) //p.x, p.y and v.x, v.y are different
	fmt.Println("")
	fmt.Println(p.func2(2))
	fmt.Println("---------------------------------")
	func3(p, 2)
	fmt.Println(p.x, p.y)
	fmt.Println("")
	fmt.Println(v.x, v.y) //p.x, p.y and v.x, v.y are different
	fmt.Println("")
	fmt.Println(func4(p, 2))
	fmt.Println("---------------------------------")
	p.func5(2)
	fmt.Println(p.x, p.y)
	fmt.Println("")
	v.func5(2) // even v.func() call will work
	fmt.Println(v.x, v.y) //p.x, p.y and v.x, v.y are different
	fmt.Println("")
	fmt.Println(p.func6(2))
	fmt.Println("---------------------------------")
	// func7(p, 2) // can't use p because of &Myvar instead we can use *p
	func7(*p, 2)
	fmt.Println(p.x, p.y)
	fmt.Println("")
	func7(v, 2) // even func(v, something) call will work
	fmt.Println(v.x, v.y) //p.x, p.y and v.x, v.y are different
	fmt.Println("")
	 // fmt.Println(func8(p, 2)) // can't use p because of &Myvar instead we can use *p
	fmt.Println(func8(*p, 2))
	fmt.Println("---------------------------------")
}
