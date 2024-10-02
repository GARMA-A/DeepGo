package main

import (
	. "fmt"
)

type key struct {
	key int
}

func (k key) getKey() int {
	return k.key
}

type value struct {
	value int
}

func (v value) getValue() int {
	return v.value
}

type intrfc interface {
	getKey() int
	getValue() int
}

func printTheKey(intr intrfc) {
	Println("KEY =", intr.getKey())
	Println("value =", intr.getValue())
}



func main() {

	x := key{100}
	Printf("%#v", x)

}
