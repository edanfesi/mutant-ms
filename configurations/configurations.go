package configurations

import (
	"mutant-ms/controllers"
	sMutants "mutant-ms/services/mutant"
)

func SetLayers() {

	// Services
	mutantService := &sMutants.MutantServices{}

	// Controllers
	controllers.NewMutants(mutantService)
}
