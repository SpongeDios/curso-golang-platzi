package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	respuesta, err := http.Get("https://pokeapi.co/api/v2/pokemon/?offset=&limit=20")
	if err != nil {
		fmt.Println(err)
	}
	e := escritorWeb{}
	io.Copy(e, respuesta.Body)
}

type escritorWeb struct {
}

func (escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}
