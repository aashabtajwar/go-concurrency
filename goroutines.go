// starter

package main

import (
	"fmt"
	"time"
)

func main() {
	go count("cat")
	count("mouse")
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 1000)
	}
}
