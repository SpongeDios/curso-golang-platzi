package main

import "fmt"

func main() {

	perrito := perro{}
	pezz := pez{}
	pajarito := pajaro{}

	moverAnimal(perrito)
	moverAnimal(pezz)
	moverAnimal(pajarito)

}

type animal interface {
	mover() string
}
type perro struct {
}
type pez struct {
}
type pajaro struct {
}

func (perro) mover() string {
	return "Soy un perro y estoy caminando"
}

func (pez) mover() string {
	return "Soy un pez y estoy nadando"
}

func (pajaro) mover() string {
	return "Soy un pajaro y estoy volando"
}

func moverAnimal(a animal) {
	fmt.Println(a.mover())
}
