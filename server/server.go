package server

import (
	"github.com/labstack/echo"
)

var server *echo.Echo

func init() {
	server = echo.New()
}

func Run(port interface{}) {
	if err := server.Start(port.(string)); err != nil {
		server.Logger.Fatal(err.Error())
	}
}
func Get() *echo.Echo {
	return server
}
