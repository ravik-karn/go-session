package main

import (
	"fmt"
)

func main() {
	channel1 := make(chan string)
	channel1 = nil
	channel2 := make(chan string)

	select {
	case message1 := <-channel1:
		fmt.Println(message1)
	case message2 := <-channel2:
		fmt.Println(message2)
	default:
		fmt.Println("No channel is ready")
	}
}