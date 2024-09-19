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

### When printing numbers using the fmt package, we can control the radix and for mat with the
### %d, %o, and %x verbs, as shown in this example:

```go
o := 0666    // convert from octal to decimal on first output
fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
x := int64(0xdeadbeef)
fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
// Output:
// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
```
### Usually a Printf format string containing multiple % verbs
### would require the same number of extra operands, but the [1] ‘‘adverbs’’ after % tell Printf to
### use the first operand over and over again. Second, the # adverb for %o or %x or %X tells Printf
### to emita 0 or 0x or 0X prefix respectively.

```go 
ascii := 'a'
unicode := 'D'
newline := '\n'
fmt.Printf("%d %[1]c %[1]q\n", ascii) // "97 a 'a'"
fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 D 'D'"
fmt.Printf("%d %[1]q\n", newline) // "10 '\n'"
```

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
```Go
true  false  iota  nil
```
## Types:
```Go
  int  int8  int16  int32  int64 uint
  uint8  uint16  uint32  uint64
  uintptr float32  float64  complex128  
 complex64 bool  byte  rune  string  error
```


## Functions:  
```Go
 make  len  cap  new  append  copy
 close  delete complex  real  imag 
 panic  recover
```


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
fmt.Printf(*p) // 0
/*
The new built-in function allocates memory. The first argument is 
a type, not a value, and the value returned is a pointer to a 
newly allocated zero value of that type.
*/


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


# Ch3


## More on Types 


### Go provides both signed and unsigned integer arithmetic.
### There are four distinct sizes of
### signed integers—8, 16, 32, and 64 bits—represented by the types
### int8, int16, int32, and int64, and corresponding unsigned
### versions uint8, uint16, uint32, and uint64
### the number after the type is the numbers  of bits 

```Go
	// Boolean type
	var isGoFun bool = true // stored on 8 bits
	fmt.Println("Boolean value:", isGoFun)

	// Integer types
	var intVal int = 42                    // Default integer type (depends on platform architecture) mostly 32 or 64
	var int8Val int8 = -128                // 8-bit signed integer (-128 to 127)
       var int16Val int16 = 32767             // 16-bit signed integer (-32,768 to 32,767)
	var int32Val int32 = 2147483647        // 32-bit signed integer (-2,147,483,648 to 2,147,483,647)
	var int64Val int64 = 9223372036854775807 // 64-bit signed integer (-2^63 to 2^63-1)
	fmt.Println("Integers:", intVal, int8Val, int16Val, int32Val, int64Val)

	// Unsigned integer types
	var uintVal uint = 42                    // Default unsigned integer type (depends on platform architecture)
	var uint8Val uint8 = 255                 // 8-bit unsigned integer (0 to 255)
	var uint16Val uint16 = 65535             // 16-bit unsigned integer (0 to 65535)
	var uint32Val uint32 = 4294967295        // 32-bit unsigned integer (0 to 4,294,967,295)
	var uint64Val uint64 = 18446744073709551615 // 64-bit unsigned integer (0 to 2^64-1)
	fmt.Println("Unsigned Integers:", uintVal, uint8Val, uint16Val, uint32Val, uint64Val)
```

## Important Note on types
### As an example familiar from other contexts, consider this sequence:
```go
var apples int32 = 1
var oranges int16 = 2
var compote int = apples + oranges // compile error
```
### At tempting to compile these three declarations produces an error message:
### invalid operation: apples + oranges (mismatched types int32 and int16)
### This typ e mismatch can be fixe d in several ways, most direc tly by converting everything to a
### common type:

```go
 var compote = int(apples) + int(oranges)
```

## Floating numbers

```Go
	// Boolean type
	var isGoFun bool = true
	fmt.Println("Boolean value:", isGoFun)

	// Integer types
	var intVal int = 42                    // Default integer type (depends on platform architecture)
	var int8Val int8 = -128                // 8-bit signed integer (-128 to 127)
	var int16Val int16 = 32767             // 16-bit signed integer (-32,768 to 32,767)
	var int32Val int32 = 2147483647        // 32-bit signed integer (-2,147,483,648 to 2,147,483,647)
	var int64Val int64 = 9223372036854775807 // 64-bit signed integer (-2^63 to 2^63-1)
	fmt.Println("Integers:", intVal, int8Val, int16Val, int32Val, int64Val)

	// Unsigned integer types
	var uintVal uint = 42                    // Default unsigned integer type (depends on platform architecture)
	var uint8Val uint8 = 255                 // 8-bit unsigned integer (0 to 255)
	var uint16Val uint16 = 65535             // 16-bit unsigned integer (0 to 65535)
	var uint32Val uint32 = 4294967295        // 32-bit unsigned integer (0 to 4,294,967,295)
	var uint64Val uint64 = 18446744073709551615 // 64-bit unsigned integer (0 to 2^64-1)
	fmt.Println("Unsigned Integers:", uintVal, uint8Val, uint16Val, uint32Val, uint64Val)

	// Floating-point types
	var float32Val float32 = 3.14  // 32-bit floating-point number
	var float64Val float64 = 3.141592653589793 // 64-bit floating-point number
	fmt.Println("Floats:", float32Val, float64Val)
```

### A float32 provides approximately six decimal digits of precision, where as a float64
### prov ides about 15 digits; float64 should be preferred for most purposes because float32
### computation sacc umulate error rapidly unless one is quite careful

```go
var f float32 = 16777216 // 1 << 24
fmt.Println(f == f+1) // "true"!
```
### how realy the floating numbers work in Go
```go
for x := 0; x < 8; x++ {
fmt.Printf("x = %d eA = %8.3f\n", x, math.Exp(float64(x)))
}
/*
x = 0 eA =    1.000
x = 1 eA =    2.718
x = 2 eA =    7.389
x = 3 eA =   20.086
x = 4 eA =   54.598
x = 5 eA =  148.413
x = 6 eA =  403.429
x = 7 eA = 1096.633
*/
```
### notice that when i say %8.3f i say that you must print the number
### on at least 8 digits width and with three decimal digits of precision
### if the number smaller than 8 digits add space at left of it or add '0' at the end

## Strings 

```Go
       // normal way
       var carType string = "BMW"
       // let go give the type
       carType := "BMW"
       // one line declaration
 
	   // String type
	   var str string = "Hello, Go!"
	   fmt.Println("String:", str)

       // normal way
       var carType string = "BMW"
       // let go give the type
       carType := "BMW"
       // one line declaration
       car,price := "BMW",20_000_000

```
### how go handle the Strings
#### you need to know that strings are just slice of bytes no more so it you try that :
```go

s := "hello, world"
fmt.Println(len(s)) // "12"
fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')

/**********/

fmt.Println(s[0:5]) // "hello"
fmt.Println(s[:5]) // "hello"
fmt.Println(s[7:]) // "world"
fmt.Println(s[:]) // "hello, world"
s[0] = 'L' // compile error: cannot assign to s[0]
// strings in go are immutable

```
#### Immutability means that it is safe for two copies of a string to share the same underlying
#### memory, making it cheap to copy strings of anylength. Similarly, a strings and a substring
#### like s[7:] may safely share the same data, so the substring operation is also cheap. No new
#### memory is allocated in either case. Figure 3.4 illustrates the arrangement of astring and two
#### of its subst rings sharing the same underlying byte array.

<img  src="GoStringMemory.png" width="600px" height="300px"/>

#### there is one more inmportant thing on Strings
#### if you use  `` ` ``  instead of `` ' `` the /n /t ...etc has no effect
```go
fmt.Println(`hello, \n\tworld`)
// hello, \n\tworld
```
## convert str -> int || str <- int

```go
x := 123
y := fmt.Sprintf("%d", x)
fmt.Println(y, strconv.Itoa(x)) // "123 123"
fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
s := fmt.Sprintf("x=%b", x) // "x=1111011"
x, err := strconv.Atoi("123") // x is an int
y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits

```


## Complex,rune,byte

```Go
	// Complex number types
	var complex64Val complex64 = 1 + 2i    // 64-bit complex number
	var complex128Val complex128 = 3.14 + 2.71i // 128-bit complex number
	fmt.Println("Complex numbers:", complex64Val, complex128Val)

	// Rune type (alias for int32, represents a Unicode code point)
	var runeVal rune = 'G'
	fmt.Println("Rune (character):", runeVal)

	// Byte type (alias for uint8, represents a single byte)
	var byteVal byte = 255
	fmt.Println("Byte:", byteVal)
```

# complex numbers

```Go
var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i
fmt.Println(x*y)                 // "(-5+10i)"
fmt.Println(real(x*y))           // "-5"
fmt.Println(imag(x*y))           // "10"

```

## Go Escape Sequences Table


| Escape Sequence | Character Represented       |
|-----------------|-----------------------------|
| `\\`            | Backslash (`\`)             |
| `\'`            | Single Quote (`'`)          |
| `\"`            | Double Quote (`"`)          |
| `\n`            | Newline (`\n`)              |
| `\r`            | Carriage Return (`\r`)      |
| `\t`            | Tab (`\t`)                  |
| `\v`            | Vertical Tab (`\v`)         |
| `\f`            | Form Feed (`\f`)            |
| `\xNN`          | Hexadecimal (`\xNN`)        |
| `\uNNNN`        | Unicode (`\uNNNN`)          |
| `\U00NNNNNN`    | Unicode (`\U00NNNNNN`)      |

**Note:**  
- `\xNN` represents a character with hexadecimal code `NN`.  
- `\uNNNN` represents a Unicode character with hexadecimal code `NNNN`.  
- `\U00NNNNNN` represents a Unicode character with a longer hexadecimal code `00NNNNNN`.

## constants on go 

```Go

const ( 
	a = 1
	b 
	c 
	d 
	e 
	f =2 
	g 
)
fmt.Println(a,b,c,d,e,f) // 1 1 1 1 1 2 2 

```

## Go Bitwise
``` go
& //bit wise AND
| //bit wise OR
^ //bit wise XOR
&^ //bit clear (AND NOT)
<< //left shift
>> //right shift
```