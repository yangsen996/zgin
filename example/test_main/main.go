package main

import (
	. "github.com/yangsen996/zgin/controllers"
	. "github.com/yangsen996/zgin/middlewares"
	"github.com/yangsen996/zgin/zgin"
)

func main() {
	zgin.NewZGin().
		Middlewares(NewUserMiddleware()).
		Formation("v1", NewUserController()).Launch()
}
