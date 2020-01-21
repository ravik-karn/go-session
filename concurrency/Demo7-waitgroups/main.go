package main

import (
	"fmt"
	"sync"
	"time"
)

func WelcomeMessage() {
	fmt.Println("Welcome")
}

func test() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		WelcomeMessage()
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("Hello World!")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done!")
}

func testAnother() {
	go func() {
		WelcomeMessage()
	}()
	go func() {
		fmt.Println("Hello World!")
	}()
	fmt.Println("Done!")
}

func main() {
	//test()
	testAnother()
}
