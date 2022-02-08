package server

import (
	"fmt"
	"mutant-ms/constants"

	"github.com/labstack/echo/v4"
)

type api struct{}

func newAPI() *api {
	return &api{}
}

func (api api) BaseRouter(server *echo.Echo) *echo.Group {
	return server.Group(fmt.Sprintf("/api/%s/", constants.Commons.ProjectName))
}
