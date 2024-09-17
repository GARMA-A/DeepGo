package main

import "fmt"


type englishBot struct{}
type spanishBot struct{}

func main(){

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}


func (spanishBot) getGreeting() string {

	//vert unique logic for generating an spanish greeting
	return "Hola!"
}




func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}



type bot interface {
	getGreeting() string
}


func (englishBot) getGreeting() string {
	//vert unique logic for generating an english greeting
	return "Hi There!"
}
