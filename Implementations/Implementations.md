# [Go to Theory](./Theory/Theory.md)
# [Go back](../README.md)

<hr/>

# cards game

### receiver function

```Go
package main
import ("fmt")

type deck []string

func (d deck) print(){


	for i, card := range d {
		fmt.Println(i,card)

	}

}

```

### so receiver functions is functions that works

### on variables that we made

## see that

```Go
package main

func main() {
 cards := deck{"Ace of Diamonds", newCard()};
 cards.print()

}
```

#### cards now can access the print funtion because it type is deck

## the model function of go contain three parts

### 1. receiver

### 2. arguments

### 3. returner

```Go
type deck []string
deck.toString()
//   receiver            // arguments       // return value
func (d deck) saveToFile(filename string) (error){

 return os.WriteFile(filename , []byte(d.toString()), 0666)

}
```

## full read write file example

```Go
// create a new type of 'deck'
// which is a slice of strings
package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func (d deck) print(){
	for i, card := range d {
		fmt.Println(i,card)
	}
}

func (d deck) toString() (string){
   return strings.Join([]string(d), ",\n")
}

func (d deck) saveToFile(filename string) (error){
 return os.WriteFile(filename , []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck{

	bs, err := os.ReadFile(filename)
	if err != nil{
		fmt.Println("Error: ", err)
		os.Exit(1)
	}



	s := strings.Split(string(bs), ",\n")
	return deck(s)
}

func main(){
cards := newDeck()
cards.toString()
cards.saveToFile("my_cards")
newDeckFromFile("my_cards").print()
}

```

# StructsFolder

#### struct in go it's very similar to that on c and c++

```Go

type ContactInfo struct{
	email string
	zipCode int
}

type Person struct{
	firstName string
	lastName string
}

alex := Person{firstName: "Alex", lastName: "Anderson"}
alex2 := Person{"Alex","Anderson"}
var alex3 Person
fmt.Println(alex) // {Alex Anderson}
fmt.Println(alex2)  // {Alex Anderson}
fmt.Println(alex3)  // {}
fmt.Printf("%+v",alex) // {firstName:Alex lastName:Anderson}%

```

```Go


type ContactInfo struct{
	email string
	zipCode int
}

type Person struct{
	firstName string
	lastName string
	contact  ContactInfo
}

jim := Person{
	firstName: "Jim",
	lastName: "Party",
	contact: ContactInfo{
		email: "jim@gmail.com",
		zipCode: 94000,
	},
}

fmt.Printf("%+v",jim)
// {firstName:Jim lastName:Party contact:{email:jim@gmail.com zipCode:94000}}%


```

#### how to update the name of that struct the answer is (pointers)

```Go
func (p *Person) updateName(newFirstName string){
	(*p).firstName = newFirstName
}

```

### if you know c or c++ it,s the same thing

```Go

func (p *Person) updateName(newFirstName string){
	(*p).firstName = newFirstName
}

func main()
{


	girgis := Person{
		firstName: "Girgis",
		lastName: "Girgis",
		contact: ContactInfo{
			email: "girgis@email.com",
			zipCode: 94000,
		},

	}

     var girgisPointer *Person = &girgis
     girgisPointer.updateName("Girgis2")
     girgis.print()
	// {firstName:Girgis2 lastName:Girgis contact:{email:girgis@email.com zipCode:94000}}
	/*****************/
       /* you can just do that*/
	/*****************/
      girgis.updateName("Girgis3")
      girgis.print()
	// {firstName:Girgis3 lastName:Girgis contact:{email:girgis@email.com zipCode:94000}}
}

```

### we need to talk about pass by value and refernce in Go

#### so whenever you create struct and pass it to function as

#### we see go create new copy of it inside the function by default(pass by value)

#### but that not the case for every thing

```go

func updateSlice(s []string){
	s[0] = "Bye"
}

func main()
{

mySlice := []string{"Hi","There","How","Are","You"}
updateSlice(mySlice)
fmt.Println(mySlice)
// output : [Bye There How Are You]
}
```

#### what is happen why the output like that ?

#### you need to understand the difference between

#### the `value types` and the `refernce types` in Go

## value types

### `int` `float` `string` `bool` `struct`

## Referce types

### `map` `slice` `channels` `pointer` `function`

#### so the value type(pass by value) if you want to change it's value inside function

#### you need to use pointers

#### the refernce type by default will change tha main element on any function(pass by refernce)

#### so you need to make manualy copy of it first if you do not want to change the original

# map Folder

### ways of declare map

```Go
func main(){

	var colors1 map[string]string


	colors2 := map[string]string{
		"red": "#ff0000",
		"orange" : "#ff4500",
		"blue" : "#0000ff",
		"green" : "#008000",
	}
	colors3 := make(map[string]string)

	fmt.Println(colors1)
	fmt.Println(colors3)
	fmt.Println(colors2)
	// output:
	// map[]
	// map[]
	// map[blue:#0000ff green:#008000 orange:#ff4500 red:#ff0000]

}

```

### asign value to empty map

```Go

	colors3 := make(map[string]string)
	colors3["white"] = "#ffffff"
       fmt.Println(colors3)
	// output
	// map[white:#ffffff]

```

### you need to use `[]` on mape not the `.key` like struct

### delete value from map

```Go


	colors3 := make(map[string]string)
	colors3["white"] = "#ffffff"

       delete(colors3 , "white")
       fmt.Println(colors3)
	// output
	// map[]



```

### iterate over a map

```Go


func printMap(c map[string]string){
	for color, hex := range c{
		fmt.Println("Hex code for", color, "is", hex)
	}
}


func main(){

	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
     printMap(colors)

    // Hex code for red is #ff0000
    // Hex code for green is #4bf745
    // Hex code for white is #ffffff

}

```

## `Map` vs `struct`

| Map                                         | Struct                                   |
| ------------------------------------------- | ---------------------------------------- |
| Reference Type                              | Value Type                               |
| all keys must be with same type             | all key are just some vars names         |
| all values must be with same type           | each value can be with different type    |
| Key are indexed we can iterate throw        | key do not support indexing no iteration |
| do not need to know all the keys at compile | must know all the keys at compile        |

# Interface Folder

### why i need Interface

#### so if i have some thing like that

```Go


type englishBot struct{}
type spanishBot struct{}



func (englishBot) getGreeting() string {
	//vert unique logic for generating an english greeting
	return "Hi There!"
}


func (spanishBot) getGreeting() string {
	//vert unique logic for generating an spanish greeting
	return "Hola!"
}
/* ERORR HERE 👇*/
func printGreeting(eb englishBot){
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot){
	fmt.Println(sb.getGreeting())
}
/*Connot declare two functions with same name without unique receiver*/

func main(){
eb := englishBot{}
sb := spanishBot{}

printGreeting(eb)
printGreeting(sb)

}

```

### this is not valid code because you can't declare two functions with the same name of course

### if you on java or C# there is something called funcion overloading to do that but not in go

### there is solution with receiver and with Interface

## solution with interface

```Go

type englishBot struct{}
type spanishBot struct{}
/* new ! 🆕 */
type bot interface {
	getGreeting() string
}
/*******/

func (englishBot) getGreeting() string {
	//vert unique logic for generating an english greeting
	return "Hi There!"
}


func (spanishBot) getGreeting() string
	//vert unique logic for generating an spanish greeting
	return "Hola!"
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}

func main(){

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

```

### what i just did is instead of create the same function twice

### i make type `bot` and tell my go programm to notice any type

### on my programm and detect if it has on it `getGreeting() string`

### if it has it then automaticly make them also of type `bot`

### that mean i can now use `englishBot` `spanishBot` as type bot

### and any type has getGreeting() on it also will be of type `bot`

## normal types vs interface type

| topic  | normal type                        | interface type                      |
| ------ | ---------------------------------- | ----------------------------------- |
| syntax | type name (int,float64,map,struct) | type name interface                 |
| assign | can assign them to many vars       | can't assign to any another var     |
| usage  | make vars static types             | add extra functionalty to some type |



