package main

import (
	"net/http"
	
	"github.com/labstack/echo"
	"github.com/fernferret/b0xexample/staticfiles"
)

func main() {
	e := echo.New()
	e.GET("/*", echo.WrapHandler(staticfiles.Handler))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}