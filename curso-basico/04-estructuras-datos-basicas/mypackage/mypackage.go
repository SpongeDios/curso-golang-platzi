package mypackage

import "fmt"

// CarPublic Car con acceso publico: Este comentario se ve si pasas el mouse por encima
type CarPublic struct {
	brand string
	year  int
}

func (c *CarPublic) Brand() string {
	return c.brand
}

func (c *CarPublic) SetBrand(brand string) {
	c.brand = brand
}

func (c *CarPublic) Year() int {
	return c.year
}

func (c *CarPublic) SetYear(year int) {
	c.year = year
}

// PrintMessage imprime un mensaje
func PrintMessage(message string) {
	fmt.Println(message)
}
