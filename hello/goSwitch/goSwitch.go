package goSwitch

import (
	"fmt"

	"annicaburns.com/hello/greeting"
)

// Cases can actually be expressions - the first one that evaluates to true will be executed
// Can switch on types (rather than the traditional switch on the value of a variable)
// https://golang.org/doc/effective_go.html#switch

// SwitchBasic demonstrates the basic switch statement in GO - If Annica evaluates to true, it will return "Ms""
func SwitchBasic(name string) (prefix string) {
	switch name {
	case "Annica":
		prefix = "Ms"
	case "Mitchel":
		prefix = "Mr"
	default:
		prefix = "Dude"
	}
	return
}

// SwitchFallthrough demonstrates using the fallthrough keyword - If Annica evaluates to true, it will return "Mr"
func SwitchFallthrough(name string) (prefix string) {
	switch name {
	case "Annica":
		prefix = "Ms"
		fallthrough
	case "Mitchel", "Tom":
		prefix = "Mr"
	default:
		prefix = "Dude"
	}
	return
}

// SwitchNothing demonstrates the fact that you don't have to switch on a value.
// The cases will each become an expression, and the first case that evaluates to true will be executed
// So it's basically just one big if/else statement
// Result will be: "b"
func SwitchNothing() (result string) {
	var a = false
	var b = true
	var c = false
	switch {
	case a:
		result = "a"
	case b, 2 == 3:
		result = "b"
	case c:
		result = "c"
	default:
		result = "nothing"
	}
	return
}

// SwitchType demonstrates switching on a type
// "interface{}" means the input parameter can be of any type - like Any in Swift
func SwitchType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case greeting.Salutation:
		fmt.Println("salutation")
	default:
		fmt.Println("unknown")
	}
}
