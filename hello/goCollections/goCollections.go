package goCollections

import (
	"fmt"

	"annicaburns.com/hello/greeting"
)

// BasicArray demonstrates the characteristics and semantics of a GO array
// No initialization required, always zero based
// Not a pointer - value type
// Two arrays of different types or sizes cannot be compared or mapped because they are different types.
// For this reason they are a bad choice as function parameters because if the param was an array with 3 members,then only arrays with 3 members can be passed in.
// https://golang.org/doc/effective_go.html#arrays
func BasicArray() {
	// var items [3]int
}

// BasicSlice demonstrates the characteristics and semantics of a GO slice
// Think of a Slice as an abstraction over an array
// Two slices of different types cannot be compared or mapped, but the size of the slice is not part of it's underlying type
// Use the make keyword to initialize the slice - otherwise it's nil
// A slice is a pointer to an underlying array - reference type
// Fixed size, but can be re-allocated with append to make it grow
// You can make a slice of a slice, and the new slice will still point to the underlying data in the original
// https://golang.org/doc/effective_go.html#slices
func BasicSlice() (salutationSlice []greeting.Salutation) {
	//var items []int = make([]int,3,5) - initial slice has 3 items (it's length), but it has a capacity of 5

	// var items = make([]int, 3) // length and capacity are the same - both set to 3
	// items[0] = 1
	// items[1] = 2
	// items[2] = 3

	// items := []int{1, 2, 3}

	salutationSlice = []greeting.Salutation{
		{Name: "Annica", Greeting: "Hello"},
		{Name: "Mitchel", Greeting: "Howdy"},
		{Name: "Joline", Greeting: "Welcome"},
	}

	return
}

// SlicingASlice demonstrates filtering down a slice by location
// This operation includes the start index but excludes the end index - [1:2] only includes the item at index 1
// [:2] will include Annica and Mitchel
// [1:] will include Mitchel and Joline
func SlicingASlice(startingSlice []greeting.Salutation) (finalSlice []greeting.Salutation) {
	// finalSlice = startingSlice[1:2]
	finalSlice = startingSlice[1:]
	return
}

// PrintFilteredSlice demonstrates SlicingASlice
func PrintFilteredSlice() {
	var finalSlice = SlicingASlice(BasicSlice())
	fmt.Println(finalSlice)
	fmt.Println(len(finalSlice))
}

func appendingASlice(startingSlice []greeting.Salutation) (finalSlice []greeting.Salutation) {
	// Can add a single element to a slice
	var biggerSlice = append(startingSlice, greeting.Salutation{Name: "Tammy", Greeting: "Salud"})
	var filteredSlice = biggerSlice[3:]
	// Or can add a slice to a slice
	finalSlice = append(filteredSlice, filteredSlice...)
	return
}

// PrintBiggerSlice demonstrates appendingASlice
func PrintBiggerSlice() {
	var finalSlice = appendingASlice(BasicSlice())
	fmt.Println(finalSlice)
}

func deletingASlice(startingSlice []greeting.Salutation) (finalSlice []greeting.Salutation) {
	// use append to cobble together all the elements you want to keep, omitting the ones you don't
	finalSlice = append(startingSlice[:1], startingSlice[2:]...)
	return
}

// PrintSmallerSlice demonstrates deletingASlice
func PrintSmallerSlice() {
	var finalSlice = deletingASlice(BasicSlice())
	fmt.Println(finalSlice)
}
