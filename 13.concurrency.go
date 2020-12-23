package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"time"
)

type page struct{
	url string
	size int
}

func main(){
	links := []string{"https://www.microsoft.com/", "https://bitly.com/", "https://www.google.com/", 
	"https://www.youtube.com/", "https://ucf.edu/", "https://my.ucf.edu/", "https://www.amazon.com/"}
	fmt.Println("--------------With out using go routine--------------")
	start := time.Now()
	for _, link := range links{
	responseSize(link)
	}
	fmt.Println("Time taken =", time.Since(start))


	fmt.Println("--------------Using go routine--------------")
	mychan := make(chan page)
	start = time.Now()
	for _, link := range links{
	go goresponseSize(link, mychan)
	}
	for i:=0; i<len(links); i++{
	rslt := <-mychan
	fmt.Println(rslt.url, ":", rslt.size)
	}
	fmt.Println("Time taken =", time.Since(start))
}

// This function uses go routine execution
func goresponseSize(url string, channel chan page) {
	response, err := http.Get(url)
	if err != nil{
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil{
		log.Fatal(err)
	}
	channel <- page{url, len(body)}
}

// This function doesn't uses go routine execution
func responseSize(url string) {
	response, err := http.Get(url)
	if err != nil{
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(url, ":", len(body))
}
