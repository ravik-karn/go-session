package main

import "fmt"

func main() {
	mychannel := make(chan int)
	mychannel <- 10
	fmt.Println(<-mychannel)
}

//Resolve
//	mychannel := make(chan int)
//	go func(){
//		mychannel <- 10
//	}()
//	fmt.Println(<-mychannel)

