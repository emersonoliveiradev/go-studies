package main

import "fmt"

// int, float32, string, bool
// slice, array, struct, map

var x int = 1
var z int

// get error if declare here
// z = 2

var empty_int int
var empty_float float32
var empty_string string
var empty_bool bool

type my_type int

var t_my_type my_type

var t_int int = 10

func main() {
	fmt.Println(x)

	// get error if try set float in the int type variable
	// x = 1.1

	z = 2
	fmt.Println(z)
	fmt.Println(empty_int)
	fmt.Println(empty_float)
	fmt.Println(empty_string)
	fmt.Println(empty_bool)

	fmt.Printf("%T\n", t_my_type)
	fmt.Printf("%T\n", t_int)

	// get error to try override type value
	// t_int = t_my_type
	t_int = int(t_my_type)
	fmt.Printf("%T\n", t_int)
}
