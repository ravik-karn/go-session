package main

import "fmt"

func WelcomeMessage() {
	fmt.Println("Welcome to Educative!")
}

func main() {
	go WelcomeMessage()

	fmt.Println("Hello World!")
}
