package main

import "fmt"

// shorthand function for variable

// var name = "Mike"

func main() {
	/* Main types 
	   string
	   bool
	   int
	   int	 int8	 int16	 int32	 int64
	   uint	 uint8	 uint16	 uint32  uint64  uintptr
	   byte - alias for uint8
	   rune - alias for int32
	   float32 float64
	   complex64 complex128
	*/
	

	// Using var
	// var name = "Mike"
	var age int32 = 20
	const isCool = true

	// Shorthand
	// name := "Mike"
	// email := "mike@gmail.com"
	size := 1.3

	name, email := "Mike", "mike@gmail.com"

	fmt.Println(name, age, isCool, email)

	// get type
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
	
}