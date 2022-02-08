package routes

import (
	"github.com/labstack/echo/v4"

	"mutant-ms/configurations"
	"mutant-ms/controllers"
	"mutant-ms/storage"
)

func Setup(baseRoute *echo.Group, postgres storage.PostgresDrivers) {

	configurations.SetLayers(postgres)

	healthCheckController := controllers.NewHealthCheckController()
	baseRoute.GET(healthCheckController.BasePath, healthCheckController.GetHealthCheck)

	mutantController := controllers.NewMutantController()
	baseRoute.POST(mutantController.MutantPath, mutantController.IsMutant)
	baseRoute.GET(mutantController.MutantStatsPath, mutantController.GetStats)

}
