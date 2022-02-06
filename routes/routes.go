package routes

import (
	"github.com/labstack/echo/v4"

	"mutant-ms/configurations"
	"mutant-ms/controllers"
)

func Setup(baseRoute *echo.Group) {

	configurations.SetLayers()

	healthCheckController := controllers.NewHealthCheckController()
	baseRoute.GET(healthCheckController.BasePath, healthCheckController.GetHealthCheck)

	mutantController := controllers.NewMutantController()
	baseRoute.POST(mutantController.BasePath, mutantController.IsMutant)
}
