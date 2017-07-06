package main

import (
	"github.com/annicaburns/learngo/goSwitch"
	"github.com/annicaburns/learngo/greeting"
)

func main() {
	// greeting.PointerExample()
	var sal = greeting.Salutation{Name: "Annica", Greeting: "Dearest"}
	// fmt.Println(goSwitch.SwitchNothing())
	goSwitch.SwitchType(sal)
	// goLoops.CollectionLoop()
	// fmt.Println(goMaps.MapDelete("Jo"))
	// goCollections.PrintSmallerSlice()
	// goInterfaces.PrintRenamable()
	// goInterfaces.PrintWriterType()
	// greeting.PrintVariadicGreet()
	// goConcurrency.ConcurrencySelect()
}

/*
* https://golang.org/doc/effective_go.html
* Basic Types
    * bool
    * string
    * int, int8, int16, int32, int64 - initialized to zero by default
        * types aren’t inter compatible, meaning for example, you can’t assign an int to an int16- it won’t automatically map to the other. have to cast or convert
    * uint, uint8, uint16, uint32, uint64, uintptr
    * byte (uint8)
        * byte and uint8 are compatible - will map directly to one or another
    * rune (int32) like a char
        * use in places where we would have a single symbol or unicode character
    * float32, float64
    * complex64, complex128
* Other Types
    * Array
        * an indexed collection with a fixed size and of a particular type
    * Slice
        * like a vector or list - an array that can grow in size
        * they have the flexibility of some of the more advanced collection types, but with the memory and performance characteristics of something more basic like an array
    * Struct
    * Pointer - used to share data. Function parameters are always passed by value (copied). Use pointers to pass by reference.
    * Function -
        * act more like types -
        * can declare them as variables and pass them around - function literals.
        * Also, variables declared inside a function retain their scope - closures
        * https://golang.org/doc/codewalk/functions/
        *
    * Interface
    * Map - dictionary
        * first class citizens
        * without classes in go, we will often use Maps or Slices as a replacement
        * https://golang.org/doc/effective_go.html#for
    * Channel
* built in functions
    * len() will give you the count of elements in a slice
    * for (loops)
        * https://golang.org/doc/effective_go.html#for
* Commands: https://golang.org/cmd/go/
    * go build (compiles)
    * go install (compiles and creates an executable file in the GOBIN directory
        * in terminal, navigate to the directory that contains your main package and run this command with no params: go install
        * from then on you can just run “go run” instead of “go run ./main.go" - can't get this to work
        * The name of executable is based on the parent directory name - not on the name of the main pkg go file
    * := declaration plus initialization
* Package reference: https://golang.org/pkg/
* Language Reference: https://golang.org/ref/spec
*/
