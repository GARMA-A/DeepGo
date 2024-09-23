package main

import (
	"fmt"
)

func reverse(s []int) {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}


}


func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
	if zcap < 2*len(x) {
		zcap = 2 * len(x)
	}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
		z[len(x)] = y
		return z
}
func howCopyWork(){

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}


func do(m1 *map[rune]int){
	(*m1)['c'] = 0
	*m1 = make(map[rune]int)
	(*m1)['d'] = 4
	fmt.Println("the m1 :" , m1)
}

func main(){
	m:= map[rune]int{'d':1 , 'g':2 , 'h':3}
	fmt.Println("the m : " , m )
	do(&m)
	fmt.Println("the m after do() : " , m)
}






