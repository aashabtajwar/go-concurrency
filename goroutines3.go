// GOROUTINES AND CHANNELS
// a channel is a way for goroutines to communicate with each other
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string) // making a channel of type string
	go count("mouse", c)
	// using a normal loop (and if the sender does not close the channel)
	// it will result in a deadlock
	// cuz the count function is finished but the main function
	// is still waiting for messages
	for {
		msg, open := <-c // a second value is also received to determin whether the channel is still open
		if !open {
			// if not open, we can break out of the loop
			break
		}
		fmt.Println(msg)
	}
}

func count(animal string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- animal
		time.Sleep(time.Millisecond * 500)
	}
	// close the channel
	close(c)
}

// NOTE: As a receiver, don't close the channel!
// cuz you don't know when the sender will send the next message
// if the sender finds the receiving channel closed, the program
// will panic

/*
	a better loop
	for msg := range c {
		fmt.Println(msg)
	}


*/
