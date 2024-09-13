package main

import "fmt"

// import "fmt"

// import ("fmt")

func main() {

 cards := newDeck()

// cards.print()
// fmt.Println( newDeck())
 
//    var card string = "Ace of Spades"
//    fmt.Println([]byte(card))
// fmt.Println(cards.toString())
// cards.saveToFile("my_cards")
cards.print()
// cards.shuffle().print()

 newDeckFromFile("my_cards").shuffle().saveToFile("my_cards")
 fmt.Println(newDeckFromFile("my_cards"))		

}


// func newCard() string {
// 	return "Five of Diamonds"
// }