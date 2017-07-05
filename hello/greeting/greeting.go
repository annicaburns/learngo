package greeting

import (
	"fmt"
)

// Capitalize the name "Salutation" to "export" it (make it visible) outside of this package
type Salutation struct {
	Name     string
	Greeting string
}

type printer func(string)

const (
	pi       = 3.14
	language = "GO"
	// iota represents successive untyped integer constants - since A is the 3rd constant in this group, it's value will be 2
	A = iota
	B = iota
	C = iota
)
const (
	// sinc a is the 1st constant in this group, it's value will be 0
	a = iota
	b
	c
)

// Return multiple values - tuple. Name the return values to assign them at different times
func createMessage(name, greeting string) (message string, alternate string) {
	message = greeting + ", " + name
	alternate = "Hey, " + name
	return
}

// Use an underscore to ignore one of the return values
// Example of a function type being passed as an argument
func Greet(salutation Salutation, passedFunctionLiteral printer) {
	_, alternate := createMessage(salutation.Name, salutation.Greeting)
	passedFunctionLiteral(alternate)
}

// If statement example - using the embedded statement format of the if statement
func IfGreet(salutation Salutation, passedFunctionLiteral printer, isFormal bool) {
	message, alternate := createMessage(salutation.Name, salutation.Greeting)
	if extraSugar := " (sweetheart)"; isFormal {
		passedFunctionLiteral(message + extraSugar)
	} else {
		passedFunctionLiteral(alternate)
	}
}

// Example of a creating a Closure
func createPrintFunction(custom string) printer {
	return func(s string) { fmt.Println(s + custom) }
}

// Example of a using a Closure
func useClosure() {
	var sal = Salutation{"Annica", "Dearest"}
	Greet(sal, createPrintFunction("000"))
}

func printString(s string) {
	fmt.Print(s)
}

func printLine(s string) {
	fmt.Println(s)
}

// Variadic functions - a variable number of parameters of a certain type - has to come as the last parameter
func variadicMessage(name string, greeting ...string) (result string) {
	result = greeting[2]
	return
}

func variadicGreet(salutation Salutation) {
	result := variadicMessage(salutation.Name, salutation.Greeting, "greeting1", "greeting2")
	println("result: ", result)
}

// PrintVariadicGreet demonstrates calling a variadic function
func PrintVariadicGreet() {
	saluation := Salutation{"Annica", "Hi"}
	variadicGreet(saluation)
}

func constantExample() {
	println(a, b, c)
}

func userDefinedTypeExample() {
	// var s = Salutation{}
	// s.name = "Annica"
	// s.greeting = "Hello"
	// var s = Salutation{greeting: "Bye", name: "Annica"}
	var s = Salutation{"Annica", "Hi"}

	fmt.Println(s.Name, s.Greeting)

}

// PointerExample demonstrates passing by reference through pointers
func PointerExample() {
	message := "Hello, little chickies"
	// &message passes a reference to the message - so it is not a copy
	var greeting = &message
	message = message + "!"
	fmt.Println(message, *greeting)
}
