package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"mutant-ms/constants"
	mutantModel "mutant-ms/models/mutants"
	mutantsService "mutant-ms/services/mutants"
	"mutant-ms/settings"
	mutantContext "mutant-ms/utils/context"
	mutantError "mutant-ms/utils/errors"
)

type mutants struct {
	BasePath       string
	ProjectName    string
	ProjectVersion string
	services       mutantsService.Services
}

var mService mutantsService.Services

func NewMutants(services mutantsService.Services) {
	mService = services
}

func NewMutantController() *mutants {
	return &mutants{
		BasePath:       "mutant",
		ProjectName:    constants.Commons.ProjectName,
		ProjectVersion: settings.Commons.ProjectVersion,
		services:       mService,
	}
}

func (mutantController mutants) IsMutant(c echo.Context) error {
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

	result := mutantController.services.IsMutant(ctx, dnaSequence.Dna)

	response := mutantModel.IsMutantResponse{
		IsMutant: result,
	}

	return c.JSON(http.StatusOK, response)
}
