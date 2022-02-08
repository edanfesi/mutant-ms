package main

import (
	"mutant-ms/server"
	"mutant-ms/settings"
	mutantPostgres "mutant-ms/storage/postgres"
)

func main() {
	server.SetupServer(
		mutantPostgres.NewPostgresStorage(settings.PostgresSetting),
	)
}
