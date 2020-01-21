package main

import "fmt"

func sendValues(myIntChannel chan int) {

	for i := 0; i < 5; i++ {
		myIntChannel <- i //sending value
	}
	close(myIntChannel)
}

func main() {
	myIntChannel := make(chan int)

	go sendValues(myIntChannel)

	for i := 0; i < 6; i++ {
		value, open := <-myIntChannel
		if !open {
			break;
		}
		fmt.Println(value) //receiving value
	}
}
