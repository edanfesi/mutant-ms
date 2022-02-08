package settings

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type postgresEnvs struct {
	Host     string `envconfig:"PG_BD_HOST" required:"true"`
	Port     string `envconfig:"PG_BD_PORT" required:"true"`
	Name     string `envconfig:"PG_BD_DATABASE" required:"true"`
	Username string `envconfig:"PG_BD_USER" required:"true"`
	Password string `envconfig:"PG_BD_PASSWORD" required:"true"`
}

type PostgresSettings struct {
	Host            string
	Port            string
	DBName          string
	User            string
	Password        string
	MaxConnections  int
	ConnMaxLifeTime time.Duration
}

var PostgresSetting PostgresSettings

func init() {
	pEnvs := postgresEnvs{}

	if err := envconfig.Process("", &pEnvs); err != nil {
		panic(err.Error())
	}

	PostgresSetting.Host = pEnvs.Host
	PostgresSetting.Port = pEnvs.Port
	PostgresSetting.User = pEnvs.Username
	PostgresSetting.Password = pEnvs.Password
	PostgresSetting.DBName = pEnvs.Name
	PostgresSetting.MaxConnections = 10
	PostgresSetting.ConnMaxLifeTime = 30 * time.Minute
}
