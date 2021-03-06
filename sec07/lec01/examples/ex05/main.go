// Section 07 - Lecture 01 : Introduction to Channels
package main

import (
	"fmt"
)

func main() {
	// declaring channels
	// ----
	var ch chan int
	fmt.Printf("ch: %v, len: %v, cap: %v\n", ch, len(ch), cap(ch))

	// making a channel without capacity (unbuffer)
	// ----
	ch = make(chan int)
	fmt.Printf("ch: %v, len: %v, cap: %v\n", ch, len(ch), cap(ch))

	// sending data to unbuffered channel (without a receiver)
	// ----
	ch <- 2 // fails while trying to send

}
