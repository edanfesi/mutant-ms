package server

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"mutant-ms/routes"
	"mutant-ms/settings"
)

type DnaSequence struct {
	Dna []string `json:"dna"`
}

func SetupServer() {
	echo := echo.New()
	echo.HideBanner = true

	server := newAPI()

	routes.Setup(server.BaseRouter((echo)))

	echo.Logger.Fatal(
		echo.Start(fmt.Sprintf("%s:%s", settings.Commons.Host, settings.Commons.Port)),
	)
}
