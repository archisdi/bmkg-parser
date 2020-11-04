package main

import (
	"bmkg/controllers"
	"bmkg/modules"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func initialize() {
	if envErr := godotenv.Load(); envErr != nil {
		log.Fatal("error while loading environment file")
	}

	modules.InitializeRedis(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_USERNAME"), os.Getenv("REDIS_PASSWORD"))
}

func setupMvc(app *mvc.Application) {
	app.Handle(new(controllers.EarthquakeController))
}

func main() {
	initialize()

	app := iris.New()
	mvc.Configure(app.Party("/"), setupMvc)

	if err := app.Listen(":8080", iris.WithLogLevel("debug")); err != nil {
		log.Fatal("unable to start server")
	}
}
