package goLoops

import (
	"fmt"

	"github.com/annicaburns/learngo/greeting"
)

// GO has only one looping keyword (for), but it's not true that there is only one type of loop.
// Easy to create the equivalent of a while loop and a collection loop with the FOR keyword because elements are all optionsl;
// https://golang.org/doc/effective_go.html#for

func vendSalutation() (salutation greeting.Salutation) {
	return greeting.Salutation{Name: "Annica", Greeting: "Hello"}
}

// BasicForLoop demonstrates
func BasicForLoop(times int) {
	sal := vendSalutation()
	for i := 0; i < times; i++ {
		fmt.Println(sal.Greeting+", ", sal.Name)
	}
}

// WhileLoop demonstrates a FOR loop with a condition
func WhileLoop(times int) {
	salutation := vendSalutation()
	i := 0
	for i < times {
		fmt.Println(salutation.Greeting+", ", salutation.Name)
		i++
	}
}

// InfiniteLoop demonstrates a FOR loop that will never end unless you call the break keyword at some point
func InfiniteLoop(salutation greeting.Salutation, times int) {
	i := 0
	for {
		i++
		fmt.Println(salutation.Greeting+", ", salutation.Name)
		if i >= times {
			break
		}
	}
}

// LoopWithContinue demonstrates the use of the continue keyword
// continue short circuts the loop so that all code in the loop below the continue keyword does not get executed, meanwhile the loop starts again
// with a times param of 6 - this should only print 3 of them because only 3 are odd numbers
func LoopWithContinue(salutation greeting.Salutation, times int) {
	i := 0
	for {
		if i >= times {
			break
		}
		if i%2 == 0 {
			i++
			continue
		}
		fmt.Println(salutation.Greeting+", ", salutation.Name)
		i++
	}
}

// CollectionLoop demonstrates a FOR loop with a range
// ranges can work on 4 types:
//array or slice
//string (will get a rune for each char in the string)
//map
//channel (waiting for some data to come into the channel)
func CollectionLoop() {
	slice := []greeting.Salutation{
		{Name: "Annica", Greeting: "Hello"},
		{Name: "Mitchel", Greeting: "Hi"},
	}
	for _, s := range slice {
		fmt.Println(s.Greeting+", ", s.Name)

	}
}
