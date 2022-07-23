package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Exemplo(c echo.Context) error {
	return c.String(http.StatusOK, "Exemplo 4Linux")
}
func main() {
	e := echo.New()
	e.GET("/", Exemplo)
	e.Logger.Fatal(e.Start(":1323"))
}
