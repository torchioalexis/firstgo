package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	//INSTANCIAR
	e := echo.New()

	//RUTA
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "¡¡VAMOS CANAYA!!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
