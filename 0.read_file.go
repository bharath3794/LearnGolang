package main

import(
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
)

func main(){
	file, err := os.Open("data.txt")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("file (data.txt) =", file)
	scanner := bufio.NewScanner(file)
// Scan() reads line by line and returns true if it read data successfully and 
// false if it did not in each line
	for scanner.Scan(){
		fmt.Println(scanner.Text())
		valueInt, err := strconv.ParseInt(scanner.Text(), 10, 64)
		fmt.Printf("%v converted to type %[1]T\n, err = %v", valueInt, err)
	}
	err = file.Close()
	if err != nil{
		log.Fatal(err)
	}
	if scanner.Err() != nil{
		log.Fatal(scanner.Err())
	}
}