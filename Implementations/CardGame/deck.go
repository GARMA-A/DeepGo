// create a new type of 'deck'
// which is a slice of strings
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck{
  
  
	cards := deck{}

   
	
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}
	
	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) print(){


	for i, card := range d {
		fmt.Println(i,card)
		
	}

}



func deal(d deck , handSize int) (deck, deck){
	
	return d[:handSize], d[handSize:]
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


func (d deck) shuffle()(deck){
       source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d{
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}




