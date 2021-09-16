package main

import "fmt"

func ClosedChannel() {
	c := make(chan int)
	go sendIn(c)
	receiveOut(c)
	close(c)
	fmt.Println("about to exit")

}
func sendIn(c chan<- int){
	c <- 22
}
func receiveOut(c <-chan int){
	fmt.Println("the value received from the channel:",<-c)
}

