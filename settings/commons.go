package settings

import (
	"strings"

	"github.com/kelseyhightower/envconfig"
)

type commons struct {
	Port string `envconfig:"PORT" required:"true" default:"8080"`
	Host string `envconfig:"HOST" required:"true" default:"0.0.0.0"`

	MicroServiceURL string `envconfig:"MICROSERVICES_URL" required:"true"`

	XApplicationID string `envconfig:"APPLICATION_ID" default:"mutant-ms/1.0.0"`

	Country string `envconfig:"COUNTRY" required:"true"`

	ProjectVersion string
}

var Commons commons

func init() {
	err := envconfig.Process("", &Commons)
	if err != nil {
		panic(err.Error())
	}

	infoApp := strings.Split(Commons.XApplicationID, "/")
	Commons.ProjectVersion = getVersion(infoApp[1:])

	Commons.Country = strings.ToUpper(Commons.Country)
}

func getVersion(preVersion []string) string {
	return strings.Join(preVersion, "-")
}
