package goMaps

import (
	"fmt"
)

// Use the make keyword to initialize a map
// Keys need to be unique, but values don't
// The type used for a map key in GO needs to have the equality operator defined for it
// slice and map types do not have the equality operator defined and can't be used
// Maps are reference types - behaves like a pointer
// Maps are not thread safe - avoid using maps concurrently
// Can insert, update, delete, check for existence
// https://golang.org/doc/effective_go.html#maps

func basicMap() {
	fmt.Println("map")
}

// MapBasic demonstrates the basic use of a map type and the insert operation
func MapBasic(name string) (prefix string) {
	var prefixMap map[string]string
	prefixMap = make(map[string]string)

	prefixMap["Annica"] = "Ms "
	prefixMap["Mitchel"] = "Mr "
	prefixMap["Joline"] = "Mrs "
	prefixMap["Jo"] = "Mr "

	return prefixMap[name]

}

// MapUpdate demonstrates a shorthand way to initialize and define a map and how to update a map
// Update and Insert use the same syntax
func MapUpdate(name string) (prefix string) {
	prefixMap := map[string]string{
		"Annica":  "Ms ",
		"Mitchel": "Mr ",
		"Joline":  "Mrs ",
		"Jo":      "Mr ",
	}
	// update our map
	prefixMap["Jo"] = "Mrs "

	return prefixMap[name]
}

// MapDelete demonstrates how to delete a member from a map and how to check for existence
func MapDelete(name string) (prefix string) {
	prefixMap := map[string]string{
		"Annica":  "Ms ",
		"Mitchel": "Mr ",
		"Joline":  "Mrs ",
		"Jo":      "Mr ",
	}
	// delete a member from our map
	delete(prefixMap, "Jo")

	if value, exists := prefixMap[name]; exists {
		return value
	}

	return "Dude "
}
