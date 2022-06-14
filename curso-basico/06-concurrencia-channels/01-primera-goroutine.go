package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//WaitGroup te permite manejar de manera nativa las goroutines por lo tanto es mas eficiente.
	//de contra, es mas dificil de mantener
	var wg sync.WaitGroup
	fmt.Println("Hello")
	wg.Add(1)
	go otherSay("World", &wg)

	wg.Wait()

	//time.Sleep() no es recomendable
	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1)
}

func otherSay(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}
