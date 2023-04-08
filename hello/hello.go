package main

import "fmt"

// package level scope scope
var cooper_daughter = "Murphy Cooper"
var y = 10

// get error
//something := "Something..."

func main() {
	fmt.Println("Hello, Amelia Brand!")
	fmt.Println("Hello,", 2040, "year!")

	// get error if don't use these variables
	bytes_number, errors := fmt.Println("Something here!")
	fmt.Println(bytes_number, errors)

	// automatic typing
	// Just possible in the code block
	numbers := 12345
	floats := 1.2
	strings := "Something..."
	bools := true
	fmt.Println(numbers, floats, strings, bools)

	// get type
	pilot := "Joseph Cooper"
	movie_duration := 169
	fmt.Printf("Hello, %v - %T\n", pilot, pilot)
	fmt.Printf("Movie duration, %v - %T\n", movie_duration, movie_duration)

	starship := "Enduraaaaaaance"
	fmt.Println("Wrong name to starship", starship)

	// get error if try override
	// starship := "Endurance"
	starship = "Enduranceeeee"
	fmt.Printf("Starship, %v\n", starship)

	starship, starship_size := "Endurance", 69
	fmt.Printf("Starship: %v\nSize: %v\n", starship, starship_size)

	// outiside scope
	fmt.Println(cooper_daughter)

	// Call function
	sum(y)
}

// Declare function
func sum(x int) {
	fmt.Println(x)
	// package level scope scope
	fmt.Println(y)
}
