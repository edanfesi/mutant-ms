package server

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"mutant-ms/routes"
	"mutant-ms/settings"
	storage "mutant-ms/storage"
)

type DnaSequence struct {
	Dna []string `json:"dna"`
}

func SetupServer(postgres storage.PostgresDrivers) {
	echo := echo.New()
	echo.HideBanner = true

	server := newAPI()

	routes.Setup(server.BaseRouter((echo)), postgres)

	echo.Logger.Fatal(
		echo.Start(fmt.Sprintf("%s:%s", settings.Commons.Host, settings.Commons.Port)),
	)
}
