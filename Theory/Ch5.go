package main

import (
	"fmt"
	"time"
)

func fanIn(input1, input2 <-chan string ) <-chan string{
	c := make(chan string)

	go func(){
        
	 for {
		select {
			case s := <-input1: c<-s
			case s := <-input2: c<-s
			default : fmt.Println("no one was ready to communicate")
		}
	 }

	}()

	return c
}

func boring(msg string) <-chan string {

	ch := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {

			ch <- fmt.Sprintf("%s %d" , msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func main() {

	c := boring("joe")
	timeout := time.After(4 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("Bro you are taking to much time to answer !!")
			return 
		}
	}
}
