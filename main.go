package main

import (
	"bmkg/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func setup(app *mvc.Application) {
	app.Handle(new(controllers.EarthquakeController))
}

func main() {
	app := iris.New()

	mvc.Configure(app.Party("/"), setup)

	app.Listen(":8080", iris.WithLogLevel("debug"))

}
