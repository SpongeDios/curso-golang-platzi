package main

import "fmt"

func main() {
	c := make(chan int, 1) //<-- canal con buffer
	c <- 1
	fmt.Println(<-c)

	//sinBuffer := make(chan int) //<-- canal sin buffer
}
