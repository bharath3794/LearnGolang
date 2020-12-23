package main

import (
	"fmt"
	"10.Dependencies/setget"
	"log"
)

func main(){
	event1 := setget.Event{}
	err := event1.SetDay(3)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("event1.day =", event1.Day())
	err = event1.SetMonth(7)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("event1.name =", event1.Month())
	err = event1.SetYear(1994)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("event1.year =", event1.Year())
	event1.SetName("Bharath")
	fmt.Println("event1.name =", event1.Name())
	event1.SetTitle("Birthday")
	fmt.Println("event1.title =", event1.Title())
	fmt.Println("-------------------------------")
	fmt.Printf("event1 = %v\nGo Representation = %#[1]v\n", event1)
}