package main

import (
	"strings"
	"fmt"
	"flag"
)

// Boiling prints the boiling point of water.
const boilingF = 212.0
func main11() {
    var f = boilingF
    var c = (f - 32) * 5 / 9
    fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
    // Output:
    // boiling point = 212°F or 100°C
}




// Ftoc prints two Fahrenheit-to-Celsius conversions
func main12() {
    const freezingF, boilingF = 32.0, 212.0
    fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
    fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"

}

func fToC(f float64) float64 {
    return (f-32) * 5 / 9
}



// Echo4 prints its command-line arguments.
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main13() {
    flag.Parse()
    fmt.Print(strings.Join(flag.Args(), *sep))
    if !*n {
        fmt.Println()    
    }

}





var global *int
func g() {
     y := new(int)
     *y = 1
}   



func arrPrint()  {  
  
   var i, j int = 0, 1       
   var a []string = []string{"5", "5"}
   a[i], a[j] = a[j], a[i]
  
 
    fmt.Printf("i=%d||j=%d|| a[i]=%s || a[j]=%s" , i ,j,a[i],a[j])

}






func print(){
    fmt.Println("hello")
}