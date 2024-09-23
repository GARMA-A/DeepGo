package main

import (
	"fmt"
)



func main(){
	fmt.Println("first line of main")
	defer func(){ 
		if r:=recover(); r!=nil{
			fmt.Println("Recovered from ", r)
			fmt.Println("here on recover  do some clean up or continue the program on other functions") 
		}
	}()
	var input int
	fmt.Scanln(&input)
	switch input{
	case 1:
		panic("panic1")
	case 2:
		panic("panic2")
	case 3:
		panic("panic3")
	default:
		fmt.Println("default")
	}
	 fmt.Println("This won't be printed if there is panic happen")
}




