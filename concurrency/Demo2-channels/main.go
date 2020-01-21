package main

import "fmt"

func sendValues(myIntChannel chan int) {
	for i := 0; i < 2; i++ {
		myIntChannel <- i //sending value
	}

}

func main() {
	myIntChannel := make(chan int)

	go sendValues(myIntChannel)

	for i := 0; i < 3; i++ {
		fmt.Println(<-myIntChannel) //receiving value
	}
}
