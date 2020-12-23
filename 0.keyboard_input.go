package main

import(
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter something: ")
	input, err := reader.ReadString('\n')
	if err != nil{
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	fmt.Printf("input = %v of type %[1]T\n", input)
	inpFloat, err := strconv.ParseFloat(input, 64)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("input = %v converted to type %[1]T\n", inpFloat)
	inpInt, err := strconv.ParseInt(input, 10, 64)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("input = %v converted to type %[1]T", inpInt)
}