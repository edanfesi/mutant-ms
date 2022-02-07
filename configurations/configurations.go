package configurations

import (
	"mutant-ms/controllers"
	sMutants "mutant-ms/services/mutants"
)

func SetLayers() {

	// Services
	mutantsService := &sMutants.MutantsServices{}

	// Controllers
	controllers.NewMutants(mutantsService)
}
