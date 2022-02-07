package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"mutant-ms/constants"
	mHealthCheck "mutant-ms/models/healthcheck"
	"mutant-ms/settings"
)

type healthCheck struct {
	BasePath       string
	ProjectName    string
	ProjectVersion string
}

func NewHealthCheckController() *healthCheck {
	return &healthCheck{
		BasePath:       "health-check",
		ProjectName:    constants.Commons.ProjectName,
		ProjectVersion: settings.Commons.ProjectVersion,
	}
}

func (cHealthCheck healthCheck) GetHealthCheck(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusOK,
		mHealthCheck.Response{
			Name:    cHealthCheck.ProjectName,
			Version: settings.Commons.ProjectVersion,
			Date:    time.Now().UTC(),
		},
	)
}
