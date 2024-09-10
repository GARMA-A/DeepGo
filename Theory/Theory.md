# [Go Back](../README.md)
# [Go to Implementations](./Implementations/Implementations.md)
<hr/>


# CH1

# Go philosophy
#### Go has a strong enough type system to avoid most careless mistakes found in dynamic languages.
#### Its type system is simpler compared to other statically typed languages.
#### Go can lead to some "untyped" pockets within its type-safe framework.
#### Go programmers do not use types for safety proofs like in C++ or Haskell.
#### Despite a simpler type system, Go provides safety and performance benefits without much complexity.
#### Go emphasizes modern system design, focusing on data locality.
#### Its data types and libraries work naturally without needing explicit initialization or constructors.
#### Go avoids hidden memory allocations and writes in code.

<hr/>

## Hello World
```Go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}

```

#### Go code is organized into packages,
#### which are similar to libraries or modules in other
#### languages. A package consists of one or more .go source files
#### in a single directory that
#### define what the package does. Each source file begins with a
#### package declaration, here `package main`, that states which package the file belongs to,
#### followed by a list of other packages that it imports, and then the declarations
#### of the program that are stored in that file.
<hr/>

# package `main`

#### Package `main` is special. It defines a standalone executable
#### program, not a library. Within `package main`, the
#### function `main` is
#### also special—it’s where execution of the program begins. Whatever `main`
#### does is what the program does. Of course, `main` will normally call upon
#### functions in other packages to do much of the work, such as
#### the function `fmt.Println`.

# import()

#### The import declarations must follow the package declaration. 
#### After that,a program consists of the declarations of functions, variables,
#### constants, and types
#### (introduced by the keywords func, var, const, and type);

<hr/>

# variables
### you can declare a variable with
```Go
s := ""
var s string
var s = ""
var s string = ""
```
### some very usefull format expresions
|sympol                       | mean                                        |
|:---------------------------:|:--------------------------------------------|
| `%d`                        | decimal integer                             |
| `%x` , `%o` , `%b`          | integer in hexadecimal,cotal,binary         |
| `%f` , `%g` , `%e`          | floating number 3.14 , 3.1415926 , 3.14e+00 |
| `%t`                        | boolean                                     |
| `%c`                        | rune (unicode code point)                   |               
| `%s`                        | string                                      |
| `%q`                        | quoted string "abc"                         |
| `%v`                        | any value in a natural format               |
| `%T`                        | Type of any value                           |
| `%%`                        | literal percent sign (no operand)

### we will using them on `Printf()` and `Scanf()` a lot !
 


#### Go comes with an extensive standard library of useful packages, and the Go
#### community has created and shared many more.

#### Programming is often more about using existing packages than about
#### writing original code of one’s own.

#### Throughout the book, we will point out a couple of dozen of the most important
#### standard package

# CH2
 
# Go has 25 keywords they can’t be used as names.
```go
break
case
chan
const
continue
default
defer
else
fallthrough
for
func
go
goto
if
import
interface
map
package
range
return
select
struct
switch
type
var
```

### In addition, there are about three dozen predeclared
### names like int and true for built-in constants , types, functions

## Constants: 
### true  false  iota  nil

## Types:
###  int  int8  int16  int32  int64 uint
###  uint8  uint16  uint32  uint64
###  uintptr float32  float64  complex128  
### complex64 bool  byte  rune  string  error


## Functions:  
### make  len  cap  new  append
###  copy  close  deletecomplex  real  imagpanic  recover


## Access declared functions and variables in Go

```Go
package main


import("fmt")

// Boiling prints the boiling point of water.
const boilingF = 212.0
func main11() {
    var f = boilingF
    var c = (f - 32) * 5 / 9
    fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
    // Output:
    // boiling point = 212°F or 100°C
}
```

#### The constant boilingFis a package-level declaration (as is
#### main), whereas the variables fand care
#### local to the function main.

#### The name of each package-level entity is
#### visible not only throughout the source file that contains its
#### declaration, but throughout all the files of the package.


#### By contrast, local declarations are visible only within
#### the function in which they are declared and perhaps only within a
#### small part of it.


## pointers in Go
```Go
var x, y int
fmt.Println(&x == &x, &x == &y, &x == nil) 
// "true false false"
/************/
var p = f()
func f() *int {
    v := 1
    return &v
}

// each call return a different value
fmt.Println(f() == f()) // "false"

/**********************/

func incr(p *int) int {
    *p++ // increments what p points to; does not change p
    return *p
}

/***********************/

p := new(int)

fmt.Printf(*p)



```
# Types in Go
```Go
// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC     Celsius = 0
    BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
```
