package main

import (
	"bmkg/controllers"
	"bmkg/modules"
	"bmkg/repositories"
	"bmkg/services"
	"bmkg/utils"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func setup(app *mvc.Application) {
	// register dependencies
	app.Register(
		services.NewEarthquakeService(
			repositories.NewEarthquakeRepository(),
			repositories.NewCacheRepository(),
		),
	)

	// register controllers
	app.Handle(new(controllers.EarthquakeController))
	app.Handle(new(controllers.WeatherController))

	// register error handler
	app.HandleError(func(ctx iris.Context, err error) {
		_, _ = ctx.JSON(map[string]interface{}{
			"message": err.Error(),
		})
	})
}

func main() {
	// load environment variable
	if envErr := godotenv.Load(); envErr != nil {
		log.Fatal("error while loading environment file")
	}

	// connect to redis server
	if redErr := modules.InitializeRedis(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_USERNAME"), os.Getenv("REDIS_PASSWORD")); redErr != nil {
		log.Fatal("error while connecting to redis server")
	}

	// init constants
	utils.InitializeConstant()

	// initialize app and mvc module
	app := iris.New()
	mvc.Configure(app.Party("/"), setup)

	// listen to http port
	if err := app.Listen(":"+os.Getenv("APP_PORT"), iris.WithoutBodyConsumptionOnUnmarshal); err != nil {
		log.Fatal("unable to start server")
	}
}
