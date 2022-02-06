package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"mutant-ms/constants"
	mutantModel "mutant-ms/models/mutant"
	mutantService "mutant-ms/services/mutant"
	"mutant-ms/settings"
	mutantContext "mutant-ms/utils/context"
	mutantError "mutant-ms/utils/errors"
)

type mutant struct {
	BasePath       string
	ProjectName    string
	ProjectVersion string
	services       mutantService.Services
}

var mService mutantService.Services

func NewMutants(services mutantService.Services) {
	mService = services
}

func NewMutantController() *mutant {
	return &mutant{
		BasePath:       "mutant",
		ProjectName:    constants.Commons.ProjectName,
		ProjectVersion: settings.Commons.ProjectVersion,
		services:       mService,
	}
}

func (mutantController mutant) IsMutant(c echo.Context) error {
	ctx, cancel := mutantContext.GeContextWithTimeout(c, 5*time.Minute)
	log := mutantContext.GetLogger(ctx)

	defer cancel()

	dnaSequence := new(mutantModel.DnaSequence)
	if err := c.Bind(dnaSequence); err != nil {
		log.Errorf("[is_mutant][err:%s]", err.Error())
		return c.JSON(mutantError.HandlerError(err))
	}

	err := mutantController.services.ValidateDna(ctx, dnaSequence.Dna)
	if err != nil {
		log.Errorf("[is_mutant][err:%s]", err.Error())
		return c.JSON(mutantError.HandlerError(err))
	}

	result, err := mutantController.services.IsMutant(ctx, dnaSequence.Dna)
	if err != nil {
		log.Errorf("[is_mutant][err:%s]", err.Error())
		return c.JSON(mutantError.HandlerError(err))
	}

	response := mutantModel.IsMutantResponse{
		IsMutant: result,
	}

	return c.JSON(http.StatusOK, response)
}
