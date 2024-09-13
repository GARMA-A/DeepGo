# [Go to Theory](./Theory/Theory.md)
# [Go back](../README.md)
<hr/>

# cards game

### reciver function

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
### so reciver functions is functions that works 
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
### 1. reciver
### 2. arguments
### 3. returner
```Go
type deck []string
deck.toString()
//   reciver            // arguments       // return value
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

return <div> hello , degen</div>


