package goConcurrency

import (
	"fmt"
	"time"

	"github.com/annicaburns/learngo/goInterfaces"
)

// The concurrency problem around eggs - don't use the data to communicate information - don't share data
// When John and his wife were both counting the eggs to determine when to buy more eggs, they both bought eggs on the same day
// Instead, assign control of/access to the data to a single resource (thread - Goroutines) and use commication (channels) to control application behavior
// Goroutine is a lightweight thread managed by the GO runtime. Use keyword "go" to execute a "routine" concurrently. run this and keep going, we don't want to block anything waiting fot this to finish
// https://golang.org/doc/effective_go.html#goroutines

// BasicConcurrency demonstrates Go's built in concurrency handling
func BasicConcurrency() {
	go iterateAndPrint(3, true)
	iterateAndPrint(3, false)
	// If we don't add in a "wait" period, the BasicConcurrency function will exit before the  the asyncronous call (go iterateAndPrint)
	// has time to spin up and finish.
	// We have to keep this method alive for long enough to finish
	time.Sleep(100 * time.Millisecond)
}

func printGreeting(salutation goInterfaces.Salutation, isFormal bool) {
	var greeting = salutation.CasualGreeting
	if formalGreeting := salutation.FormalGreeting; isFormal {
		greeting = formalGreeting
	}
	fmt.Println(greeting+", ", salutation.Name)
}

func iterateAndPrint(times int, isFormal bool) {
	var salutations = goInterfaces.VendSalutations()
	for i := 0; i < times; i++ {
		printGreeting(salutations[i], isFormal)
	}
}

// ChannelConcurrency demonstrates using a channel to communicate to the main thread when the asynchronous call is finished
// This allows the routine to finish before we exit the function
// Channels are blocking - meaning that if a value is waiting to be written to a full channel (without enough buffer space)
// the channel will block it's thread until enough values are read out of the channel for the waiting content to enter.
// Also... if a channel is being read in a for loop, that loop will not exit until the channel is closed by
// whatever is feeding it
func ChannelConcurrency() {
	// create the channel
	done := make(chan bool)
	// create and execute an anonymous function to augment iterateAndPrint with the ability to communicate over a channel
	// this anonymous function is also a closure and can access the value of the done variable
	go func() {
		iterateAndPrint(3, true)
		done <- true
	}()
	iterateAndPrint(3, false)
	// we could create a variable to read the value out of done, but it's not necessary
	// because this line will block until we can read a value out of done, which won't happen until we write to done
	<-done
}

// UnBufferedChannel demonstrates that we can only have one item on a channel at a time.
// Unbuffered channels are serial - channel processes can only be run one at a time
// Buffered channels process as many routines as they can before they block
func UnBufferedChannel() {
	// create the channel
	done := make(chan bool)
	go func() {
		iterateAndPrint(3, true)
		done <- true
		// This second true will never be allowed to get onto the channel because it's unbuffered. This will block
		// indefinitely, but as soon as the first done moves onto the channel, the function will exit and the println
		// will never be reached
		done <- true
		println("Done!")
	}()
	iterateAndPrint(3, false)
	// we could create a variable to read the value out of done, but it's not necessary
	// because this line will block until we can read a value out of done, which won't happen until we write to done
	<-done
}

// BufferedChannel demonstrates that if we change the code above to be a buffered channel with room for 2 items
// then we are able to reach our println before the funciton exits
// The value of the buffer size indicates how many items can be written onto the channel before the channel has to be read
// a buffer size of one means the channel will get read after the first item is written onto the channel
// a buffer size of two means the channel won't get read until after the second item is written onto the channel
// But this code actually creates a race condition because SOMETIMES the println won't be reached before the function exists
func BufferedChannel() {
	// create the channel
	done := make(chan bool, 2)
	go func() {
		iterateAndPrint(3, true)
		done <- true
		// This second true will never be allowed to get onto the channel because it's unbuffered. This will block
		// indefinitely, but as soon as the first done moves onto the channel, the function will exit and the println
		// will never be reached
		done <- true
		println("Done!")
	}()
	iterateAndPrint(3, false)
	// we could create a variable to read the value out of done, but it's not necessary
	// because this line will block until we can read a value out of done, which won't happen until we write to done
	<-done
}

// FixedChannel demonstrates code that solves the race condition that exists in the BufferedChannel function
// race condition: SOMETIMES the println won't be reached before the function exists
//
func FixedChannel() {
	// create the channel
	done := make(chan bool, 2)
	go func() {
		iterateAndPrint(3, true)
		done <- true
		// Introducing a sleep here demonstrates that the race condition exists
		time.Sleep(100 * time.Millisecond)
		done <- true
		println("Done!")
	}()
	iterateAndPrint(3, false)
	// we could create a variable to read the value out of done, but it's not necessary
	// because this line will block until we can read a value out of done, which won't happen until we write to done
	<-done
	// The infinite loop below will block and the function will never exit on it's own,
	// which means the println above will always be reached - but this obviously isn't a good solution.
	// Instead we would need to read a value and react by manually exiting the function
	for {
		time.Sleep(100 * time.Millisecond)
	}
}

// ChannelWithRange demonstrates
func ChannelWithRange() {
	var salutations = goInterfaces.VendSalutations()
	// create a channel that will hold Salutations
	salChannel := make(chan goInterfaces.Salutation)
	// Use a goroutine to fill the channel with salutations.
	// It will run asynchronously and fill the channel with a new value each time a value gets read out the other end
	// Eventually, when all values have been fed into the channel, the channel will be closed by ChannelGreeter.
	go salutations.ChannelGreeter(salChannel)
	for salutation := range salChannel {
		fmt.Println(salutation.Name)
		// This loop will run as long as the channel is open and pull values out of the channel (by reading and printing)
		// them until it receives a "channel closed" message after the last salutation.
		// This loop will then exit and the function will exit.
	}
}

// ConcurrencySelect demonstrates the Select statement which is like a switch statement - but on communications
// Rules
// execute case that is "ready"
// if more than one is "ready", execute one at random
// if none are ready, block unless a default is defined
func ConcurrencySelect() {
	var salutations = goInterfaces.VendSalutations()
	salChannel1 := make(chan goInterfaces.Salutation)
	salChannel2 := make(chan goInterfaces.Salutation)
	go salutations.ChannelGreeter(salChannel1)
	go salutations.ChannelGreeter(salChannel2)
	for {
		select {
		case salutation, ok := <-salChannel1:
			if ok {
				fmt.Println(salutation.Name, ":1")
			} else {
				return
			}
		case salutation, ok := <-salChannel2:
			if ok {
				fmt.Println(salutation.Name, ":2")
			} else {
				return
			}
		default:
			fmt.Println("waiting")
		}
	}
}
