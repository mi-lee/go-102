// Declare a struct type to maintain information about a person.  Declare a
// function that creates new values of your type.  Call this function from main
// and display the value.
//
// Template available at: http://play.golang.org/p/ta6oFzjgwn
package main

// Add your imports here.
import (
	"fmt"
)

// Declare a struct type `person` to maintain information about a person.
type person struct {
	age    int
	gender bool
	name   string
}

// Declare a function that creates new values of your `person` type.
func createPerson(age int, gender bool, name string) person {
	return person{age, gender, name}
}

func main() {
	// Use you function to create a new value of type `person`.
	michelle := createPerson(25, true, "Michelle")

	// Output the value of your person.
	fmt.Printf("%s is %d years old and %b", michelle.name, michelle.age, michelle.gender)
}
