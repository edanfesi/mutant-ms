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
	MutantPath      string
	MutantStatsPath string
	ProjectName     string
	ProjectVersion  string
	services        mutantsService.Services
}

func NewMutantController(mService mutantsService.Services) *mutants {
	return &mutants{
		MutantPath:      "mutant",
		MutantStatsPath: "stats",
		ProjectName:     constants.Commons.ProjectName,
		ProjectVersion:  settings.Commons.ProjectVersion,
		services:        mService,
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

	err = mutantController.services.IsMutant(ctx, dnaSequence.Dna)
	if err != nil {
		log.Errorf("[is_mutant][err:%s]", err.Error())
		return c.JSON(mutantError.HandlerError(err))
	}

	return c.JSON(http.StatusOK, nil)
}

func (mutantController mutants) GetStats(c echo.Context) error {
	ctx, cancel := mutantContext.GeContextWithTimeout(c, 5*time.Minute)
	log := mutantContext.GetLogger(ctx)

	defer cancel()

	stats, err := mutantController.services.GetStats(ctx)
	if err != nil {
		log.Errorf("[is_mutant][err:%s]", err.Error())
		return c.JSON(mutantError.HandlerError(err))
	}

	return c.JSON(http.StatusOK, stats)
}
