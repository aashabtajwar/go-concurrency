// a more practical demonstration for concurrency
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // setting one goroutine to wait for
	go func() {
		count("cat")
		wg.Done() // remove goroutine
	}() // function is immediately invoked
	wg.Wait() // block until all goroutines are finished, (i.e. the counter Add() is 0)
}

func count(animal string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, animal)
		time.Sleep(time.Millisecond * 500)
	}
}
