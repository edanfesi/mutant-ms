package configurations

import (
	"mutant-ms/controllers"
	mutantRepo "mutant-ms/repositories/mutants"
	sMutants "mutant-ms/services/mutants"
	"mutant-ms/storage"
)

func SetLayers(postgres storage.PostgresDrivers) {

	// Repositories
	mutantRepository := mutantRepo.NewMutantsPostgres(postgres)

	// Services
	mutantsService := &sMutants.MutantsServices{
		PostgresRepo: mutantRepository,
	}

	// Controllers
	controllers.NewMutants(mutantsService)
}
