package goInterfaces

import (
	"fmt"
)

// https://golang.org/doc/effective_go.html#methods
// https://golang.org/doc/effective_go.html#interfaces

// GO methods can operate on any named type in the package in which they are defined
// GO functions operate on a specific type, while methods operate on a named type (type alias)
// To write a method for int, create a type alias in your package and then write a method to operate on it
// GO methods can also operate on pointers to named typed

// In GO, any type that implements same methods as an Interface implements that interface
// To accomplish Any in Swift, use and empty interface - interface{} - because all types in go will automatically conform
// See goSwitch.SwitchType for an example

// Salutation is a single object
type Salutation struct {
	Name           string
	CasualGreeting string
	FormalGreeting string
}

type renamable interface {
	rename(newName string)
}

func (salutation *Salutation) rename(newName string) {
	salutation.Name = newName
}

// Salutations is a named type representing a slice of Salutations
type Salutations []Salutation

// ChannelGreeter is a method on Salutations that fills a channel with a slice of Salutations
func (salutations Salutations) ChannelGreeter(channel chan Salutation) {
	for _, s := range salutations {
		channel <- s
	}
	close(channel)
}

// VendSalutations can be used program wide to produce a starter slice of Salutations
func VendSalutations() (salutations Salutations) {
	salutations = Salutations{
		{"Annica", "Howdy", "Hello"},
		{"Mitchel", "Hey", "Hello"},
		{"Marisol", "Salud", "Hello"},
	}
	return
}

// this greet function is a method that operations on our named type - Salutations
func (salutations Salutations) greet(isFormal bool) {
	for _, s := range salutations {
		var greeting = s.CasualGreeting
		if formalGreeting := s.FormalGreeting; isFormal {
			greeting = formalGreeting
		}
		fmt.Println(greeting + ", " + s.Name)
	}
}

// PrintGreetings is used to demonstrate calling a method
func PrintGreetings() {
	var salutations = VendSalutations()
	salutations[0].rename("Jessica")
	salutations.greet(false)
}

func renameToFrog(r renamable) {
	r.rename("Frog")
}

// PrintRenamable is used to demonstrate calling a method that takes an interface parameter
func PrintRenamable() {
	var salutations = VendSalutations()
	renameToFrog(&salutations[0])
	salutations.greet(false)
}

// Implementing the GO Writer interface
func (salutation *Salutation) Write(p []byte) (n int, err error) {
	s := string(p)
	salutation.rename(s)
	n = len(s)
	err = nil
	return
}

// PrintWriterType is used to demonstrate calling a method on a type that implements an interface
func PrintWriterType() {
	var salutations = VendSalutations()

	fmt.Fprintf(&salutations[0], "%d New Name", 1)
	fmt.Println(salutations[0])
	fmt.Println(salutations[1])

}
