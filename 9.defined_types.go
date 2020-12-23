package main

import (
	"fmt"
)


// custom-defined type dozens with underlying type int
type dozens int
// custom-defined type kgs with underlying type float64
type kgs float64



func main(){
	// declaring bananas as type dozens. Here, bananas = 2 dozens
	bananas := dozens(2)
	fmt.Printf("bananas = %v %[1]T\n", bananas)
	// Pointer of bananas
	bananasPointer := &bananas
	fmt.Printf("bananasPointer = %v\n", bananasPointer)
// Passing bananas directly to both price() and pricePointer() methods
	fmt.Println("Passing bananas directly to both price() and pricePointer() methods")
// Price is 20 per dozen
	cost := bananas.price(20)
	fmt.Printf("Price of bananas = %.2f Rupees\n", cost)
	cost =  bananas.pricePointer(20)
	fmt.Printf("Price of bananas (Pointer Method) = %.2f Rupees\n", cost)
// Passing bananasPointer to both price() and pricePointer() methods
	fmt.Println("Passing bananasPointer to both price() and pricePointer() methods")
	cost = bananasPointer.price(20)
	fmt.Printf("Price of bananas = %.2f Rupees\n", cost)
	cost =  bananasPointer.pricePointer(20)
	fmt.Printf("Price of bananas (Pointer Method) = %.2f Rupees\n", cost)
	fmt.Println("-------------------------------------------------------")


// declaring grapes as type kgs. Here, grapes = 2.5 kgs
	grapes := kgs(2.5)
	fmt.Printf("grapes = %v %[1]T\n", grapes)
	// Pointer of grapes
	grapesPointer := &grapes
	fmt.Printf("grapesPointer = %v\n", grapesPointer)
// Passing grapes directly to both price() and pricePointer() methods
	fmt.Println("Passing grapes directly to both price() and pricePointer() methods")
// Price is 40 per kg
	cost = grapes.price(40)
	fmt.Printf("Price of grapes = %.2f Rupees\n", cost)
	cost =  grapes.pricePointer(40)
	fmt.Printf("Price of grapes (Pointer Method) = %.2f Rupees\n", cost)
// Passing grapesPointer to both price() and pricePointer() methods
	fmt.Println("Passing grapesPointer to both price() and pricePointer() methods")
	cost = grapesPointer.price(40)
	fmt.Printf("Price of grapes = %.2f Rupees\n", cost)
	cost =  grapesPointer.pricePointer(40)
	fmt.Printf("Price of grapes (Pointer Method) = %.2f Rupees\n", cost)
	fmt.Println("-------------------------------------------------------")





}

// Here, (d dozens) is called receiever parameter to method price(). 
// Receiever Parameters are like self, this in other programming languages.
// In Go, we pass receiver parameters instead of self, this
func (d dozens) price(costPerDozen float64) float64{
	return float64(d)*costPerDozen
}

// We can define same method on different types.
// But we can't define same type on same method one is by passing directly and 
// other is by passing pointer of that type variable.
// So we can't use (d dozens) and (d *dozens) on same price() method
func (d *dozens) pricePointer(costPerDozen float64) float64{
	return float64(*d)*costPerDozen
}


func (k kgs) price(costPerKg float64) float64{
	return float64(k)*costPerKg
}

func (k *kgs) pricePointer(costPerKg float64) float64{
	return float64(*k)*costPerKg
}


