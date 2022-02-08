package configurations

import (
	"mutant-ms/controllers"
	mutantRepo "mutant-ms/repositories/mutants"
	mutantService "mutant-ms/services/mutants"
	"mutant-ms/storage"
)

func SetLayers(postgres storage.PostgresDrivers) {

	// Repositories
	mutantReposPostgres := mutantRepo.NewMutantsRepositories(postgres)

	// Services
	mutantsService := mutantService.NewMutantService(mutantReposPostgres)

	// Controllers
	controllers.NewMutantController(mutantsService)
}
