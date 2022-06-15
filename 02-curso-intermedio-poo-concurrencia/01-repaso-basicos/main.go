package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	parseInt, err := strconv.ParseInt("a", 0, 64)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Println(parseInt)
	}

	mapa := make(map[string]int)
	mapa["Key"] = 6
	fmt.Println(mapa["Key"])

	slice := []int{1, 2, 3}
	slice = append(slice, 16)
	for index, value := range slice {
		fmt.Println("INDEX:", index)
		fmt.Println("VALUE:", value)
	}

	//channel := make(chan int)
	//go doSomething(channel)
	//<-channel

	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

func doSomething(c chan int) {
	time.Sleep(time.Second * 3)
	fmt.Println("Done")
	c <- 1
}
