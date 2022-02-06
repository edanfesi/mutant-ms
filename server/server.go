package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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

	data, _ := json.MarshalIndent(echo.Routes(), "", "  ")
	ioutil.WriteFile("routes.json", data, 0644)

	echo.Logger.Fatal(
		echo.Start(fmt.Sprintf("%s:%s", settings.Commons.Host, settings.Commons.Port)),
	)
}

/*
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.POST("/mutant", func(c echo.Context) error {
		dnaSequence := new(DnaSequence)
		if err := c.Bind(dnaSequence); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, dnaSequence)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
*/
