package main

import (
	"fmt"
	"sort"
)

type employee struct{
	name string
	salary int
	age int
}


func main(){
/*
* In new versions of Go, after 1.8 we can do it in easy way with sort.Slice method.
* See below implementation
*/
	fmt.Println("----------------Implementation in newer versions of GO (>1.8)----------------")
	fmt.Println("-----Sorting based on Name-----")
	employees := []employee{employee{"Bharath", 2000, 26}, employee{"Adhi", 3000, 12}, employee{"Teja", 1000, 28}}
	fmt.Println("Before Sorted =", employees)
	sort.Slice(employees, func(i, j int) bool {
  		return employees[i].name < employees[j].name
	})
	fmt.Println("After sorting based on Name =", employees)
	fmt.Println("-----Sorting based on Salary-----")
	employees = []employee{employee{"Bharath", 2000, 26}, employee{"Adhi", 3000, 12}, employee{"Teja", 1000, 28}}
	fmt.Println("Before Sorted =", employees)
	sort.Slice(employees, func(i, j int) bool {
  		return employees[i].salary < employees[j].salary
	})
	fmt.Println("After sorting based on Salary =", employees)
	fmt.Println("-----Sorting based on Age-----")
	employees = []employee{employee{"Bharath", 2000, 26}, employee{"Adhi", 3000, 12}, employee{"Teja", 1000, 28}}
	fmt.Println("Before Sorted =", employees)
	sort.Slice(employees, func(i, j int) bool {
  		return employees[i].age < employees[j].age
	})
	fmt.Println("After sorting based on Salary =", employees)
fmt.Println("-----Sorting based on multiple Keys, first with Name, second with salary if Names are equal", 
	"and then with age if both Names and salary are equal-----")
	employees = []employee{employee{"Bharath", 2000, 26}, employee{"Adhi", 3000, 12}, 
						   employee{"Teja", 2000, 28}, employee{"Teja", 2000, 25}, 
						   employee{"Bharath", 1000, 26}}
	fmt.Println("Before Sorted =", employees)
	sort.Slice(employees, func(i, j int) bool {
		if employees[i].name < employees[j].name {
			return true
		} else if employees[i].name > employees[j].name {
			return false
		}
		if employees[i].salary < employees[j].salary {
			return true
		} else if employees[i].salary > employees[j].salary {
			return false
		}
  		return employees[i].age < employees[j].age
	})
	fmt.Println("After sorting based on Salary =", employees)
	fmt.Println("-----Another way of doing it, More shorter version of the same above method-----")
	employees = []employee{employee{"Bharath", 2000, 26}, employee{"Adhi", 3000, 12}, 
						   employee{"Teja", 2000, 28}, employee{"Teja", 2000, 25}, 
						   employee{"Bharath", 1000, 26}}
	fmt.Println("Before Sorted =", employees)
	sort.Slice(employees, func(i, j int) bool {
		if employees[i].name != employees[j].name {
			return employees[i].name < employees[j].name
		} 
		if employees[i].salary != employees[j].salary {
			return employees[i].salary < employees[j].salary
		} 
  		return employees[i].age < employees[j].age
	})
	fmt.Println("After sorting based on Salary =", employees)
	fmt.Println("-----One more way of doing it, using Switch Cases-----")
	employees = []employee{employee{"Bharath", 2000, 26}, employee{"Adhi", 3000, 12}, 
						   employee{"Teja", 2000, 28}, employee{"Teja", 2000, 25}, 
						   employee{"Bharath", 1000, 26}}
	fmt.Println("Before Sorted =", employees)
	sort.Slice(employees, func(i, j int) bool {
		switch {
		case employees[i].name != employees[j].name:
			return employees[i].name < employees[j].name
		case employees[i].salary != employees[j].salary:
			return employees[i].salary < employees[j].salary
		default:
			return employees[i].age < employees[j].age
		}
	})
	fmt.Println("After sorting based on Salary =", employees)
	// A Special Requirement
	fmt.Println("-----If some slice has already been sorted based on certain key",
		"but you still want to sort part of this slice with one other key if some conidition is met-----")
	// Below is already sorted by Name, if you then want to sort it by salary for whom names are equal and
	// if both are equal then you want to sort it by age
	employees = []employee{employee{"Adhi", 3000, 12}, employee{"Bharath", 1000, 26}, 
						   employee{"Bharath", 1000, 22}, employee{"Teja", 4000, 28}, 
						   employee{"Teja", 2000, 28}}
	fmt.Println("Before Sorted =", employees)
	sort.Slice(employees, func(i, j int) bool {
		if employees[i].name == employees[j].name {
			if employees[i].salary != employees[j].salary {
				return employees[i].salary < employees[j].salary
			} else { //means salary's are equal
				return employees[i].age < employees[j].age
			}
		} 
  		return false
	})
	fmt.Println("After sorting based on Salary =", employees)
	fmt.Println("-----Sorting a 2-D Slice-----")
	arr1 := [][]int{[]int{2, 5}, []int{1, 4}, []int{3, 6}}
	fmt.Println("Before Sorted =", arr1)
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i][0] < arr1[j][0]
	})
	fmt.Println("After sorting based on 1st column element in each row =", arr1)
	fmt.Println("--------------------------------------------------------")

/*
* If employees is an array instead of slice we can convert array to slice using array[:]
and pass it tosort.Slice() method ex: sort.Slice(employees[:], func(i, j int) bool {})
*/
	// Convert an array to a slice easy way
	fmt.Println("Converting an array to a slice")
	arr := [3]int{1, 2, 3}
	slice8 := arr[:]
	fmt.Printf("arr = %v of type %[1]T is converted to slice8 = %v of type %[2]T\n", arr, slice8)



	olderVersionSort()
}



/*
* Implementation in older versions of GO (<1.8)
*/
// Below methods are dependencies for olderVersionSort() function
type nameSorter []employee

func (a nameSorter) Len() int {return len(a)}
func (a nameSorter) Less(i, j int) bool{return a[i].name < a[j].name}
func (a nameSorter) Swap(i, j int) {a[i], a[j] = a[j], a[i]}



type salarySorter []employee

func (a salarySorter) Len() int {return len(a)}
func (a salarySorter) Less(i, j int) bool{return a[i].salary < a[j].salary}
func (a salarySorter) Swap(i, j int) {a[i], a[j] = a[j], a[i]}



type ageSorter []employee

func (a ageSorter) Len() int {return len(a)}
func (a ageSorter) Less(i, j int) bool{return a[i].age < a[j].age}
func (a ageSorter) Swap(i, j int) {a[i], a[j] = a[j], a[i]}

func olderVersionSort(){
	

	fmt.Println("----------------Implementation in older versions of GO (<1.8)----------------")
	fmt.Println("-----Sorting based on Name-----")
	employees := []employee{employee{"Bharath", 2000, 26}, employee{"Adhi", 3000, 12}, employee{"Teja", 1000, 28}}
	fmt.Println("Before Sorted =", employees)
	sort.Sort(nameSorter(employees))
	fmt.Println("After sorting based on Name =", employees)
	fmt.Println("-----Sorting based on Salary-----")
	employees = []employee{employee{"Bharath", 2000, 26}, employee{"Adhi", 3000, 12}, employee{"Teja", 1000, 28}}
	fmt.Println("Before Sorted =", employees)
	sort.Sort(salarySorter(employees))
	fmt.Println("After sorting based on Salary =", employees)
	fmt.Println("-----Sorting based on Age-----")
	employees = []employee{employee{"Bharath", 2000, 26}, employee{"Adhi", 3000, 12}, employee{"Teja", 1000, 28}}
	fmt.Println("Before Sorted =", employees)
	sort.Sort(ageSorter(employees))
	fmt.Println("After sorting based on Salary =", employees)
}