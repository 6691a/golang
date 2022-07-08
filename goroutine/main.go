package main

import (
	"fmt"
	"time"
)

func print() {
	for {
		fmt.Printf("*")
		time.Sleep(2 * time.Second)
	}
}

func input_word(channel chan string) {
	fmt.Println("Type a word, then hit Enter.")
	var word string
	fmt.Scanf("%s", &word)
	channel <- word
}

func main() {
	defer fmt.Println("Bye")
	go print()

	channel := make(chan string)
	go input_word(channel)

	select {
	case word := <-channel:
		fmt.Println("\nReceived: ", word)
	}
}
