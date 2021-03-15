// BUFFERED CHANNELS
// the previous demonstration has limitations
package main

import "fmt"

func main() {
	// making channel
	/*
		c := make(chan string)
		c <- "hello"

		msg := <-c
		fmt.Println(msg)


		this wont work
		cuz the sender channel will get blocked
		until there is something to receive (the two must simultaneously be active)
		so the program never progresses pass the sender block
		to make this work, the receiver must be in a separate go-routine
	*/
	// An alternative to make a separte goroutine
	// make a buffered channel by giving it a capacity
	// it will work without a separate receiver in a goroutine
	// and it won't be blocked until the channel is full

	c := make(chan string, 2)
	c <- "Hello"
	c <- "World"

	msg := <-c
	fmt.Println(msg)
	msg = <-c
	fmt.Println(msg)

}
