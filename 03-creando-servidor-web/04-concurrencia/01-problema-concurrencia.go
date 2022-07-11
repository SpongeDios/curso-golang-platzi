package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()

	servidores := []string{
		"https://platzi.com",
		"https://google.com",
		"https://facebook.com",
		"https://instagram.com",
		"https://twitter.com",
		"https://twitch.tv",
	}

	canal := make(chan string)
	/*i := 0

	for {
		if i > 2 {
			break
		}
		for _, servidor := range servidores {
			go revisarServidor(servidor, canal)
		}

		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)
		i++
	}*/

	for _, servidor := range servidores {
		go revisarServidor(servidor, canal)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	tiempoPaso := time.Since(inicio)
	defer fmt.Printf("Tiempo de ejecucion %s\n", tiempoPaso)

}

func revisarServidor(servidor string, canal chan<- string) {
	_, err := http.Get(servidor)
	if err != nil {
		//fmt.Println(servidor, "no se encuentra disponible :(")
		canal <- servidor + "no se encuentra disponible :("
	} else {
		//fmt.Println(servidor, " esta funcionando normalmente")
		canal <- servidor + " esta funcionando normalmente"
	}
}
