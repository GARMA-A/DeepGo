package main

import (
	// "bufio"
	. "fmt"
	// "sort"
	// "sort"
	// "sort"
	// "strings"
	// "os"
)

// how to make a function that is the user can call it once !
func once(fn func(args ...interface{}) interface{}) func(args ...interface{}) interface{} {

	var done bool
	return func(args ...interface{}) interface{} {
		if !done {
			done = true
			return fn(args...)
		}
		return nil
	}
}

// make function that can called only twice
func twice(fn func(args ...interface{}) interface{}) func(args ...interface{}) interface{} {

	var count int
	return func(args ...interface{}) interface{} {
		if count < 2 {
			count++
			return fn(args...)
		}
		return nil
	}
}

func mainCh4() {

	myFunc := func(args ...interface{}) interface{} {
		Println("function call! , args : ", args)
		return args
	}

	onceFunc := once(myFunc)

	onceFunc(1, 2, 3)
	onceFunc(4, 5, 6)

	//-------------------------------

	twiceFunc := twice(myFunc)

	twiceFunc(1, 2, 3)
	twiceFunc(4, 5, 6)
	twiceFunc(7, 8, 9)

}
