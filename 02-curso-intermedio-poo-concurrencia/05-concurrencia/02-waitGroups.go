package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Inicio la funcion", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Termino la funcion", i)
}
