package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	//introducimos dos mensajes dentro del canal
	channel <- "Mensaje 1"
	channel <- "Mensaje 2"

	fmt.Println(len(channel), cap(channel))

	//Range y Close
	close(channel) //Cierra el channel, por lo tanto no recibe mas datos

	for message := range channel {
		fmt.Println(message)
	}

	//Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje1", email1)
	go message("mensaje2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido desde el email1:", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido desde el email2:", m2)
		}
	}

}

func message(text string, channel chan string) {
	channel <- text
}
