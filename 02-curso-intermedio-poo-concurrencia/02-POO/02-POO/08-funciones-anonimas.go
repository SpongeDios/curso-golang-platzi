package main

import (
	"fmt"
	"time"
)

func main() {
	/*x := 5
	y := func() int {
		return x * 2
	}()

	fmt.Println(y)*/

	c := make(chan int)

	go func() {
		fmt.Println("Ejecutando la funcion anonima")
		time.Sleep(5 * time.Second)
		fmt.Println("Fin de la go routine")
		c <- 1
	}()

	<-c
}
