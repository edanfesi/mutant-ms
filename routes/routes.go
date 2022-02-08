package routes

import (
	"github.com/labstack/echo/v4"

	"mutant-ms/controllers"
	mutantRepo "mutant-ms/repositories/mutants"
	mutantService "mutant-ms/services/mutants"
	"mutant-ms/storage"
)

func Setup(baseRoute *echo.Group, postgres storage.PostgresDrivers) {

	// Repositories
	mutantReposPostgres := mutantRepo.NewMutantsRepositories(postgres)

	// Services
	mutantsService := mutantService.NewMutantService(mutantReposPostgres)

	// Controllers
	healthCheckController := controllers.NewHealthCheckController()
	baseRoute.GET(healthCheckController.BasePath, healthCheckController.GetHealthCheck)

	mutantController := controllers.NewMutantController(mutantsService)
	baseRoute.POST(mutantController.MutantPath, mutantController.IsMutant)
	baseRoute.GET(mutantController.MutantStatsPath, mutantController.GetStats)

}
