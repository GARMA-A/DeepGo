package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.stackoverflow.com",
	}
	// the part that is commented is another way of doing the same
	// thing with array instead of one chanel
	
	// c := make([]chan string, len(links))

	
       c:= make(chan string)
	 
	for _, link := range links {
		go checkLink(link, c)
	}

	// for i,_ := range(links) {
	// 	fmt.Println(<-c[i])
	// }

	for range(links) {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		c <- link + " might be down!"
		return
	}

	c <- link + " is up"
}
