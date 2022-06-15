package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	//instanciamos echo
	e := echo.New()

	//Ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hola mundo, mi primer servidor web")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
