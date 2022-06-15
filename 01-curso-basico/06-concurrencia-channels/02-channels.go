package main

import "fmt"

func main() {
	//Menos eficiente que los wildgroups. A menos que necesitas una enorme optimizacion, los channels son la mejor opcion
	//El segundo parametro del channel es el espacio de goroutines, en este caso es solo 1
	c := make(chan string, 1)
	fmt.Println("Hello")
	go say("Bye", c)
	fmt.Println(<-c)
}

//La <- al lado de chan, dice que ese channels solo recibe datos
func say(text string, c chan<- string) {
	//Le anhadimos al channel la variable text
	c <- text
}
