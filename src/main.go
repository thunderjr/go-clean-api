package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/thunderjr/go-clean-api/src/application/router"
	migrations "github.com/thunderjr/go-clean-api/src/infra/database"
	local_database "github.com/thunderjr/go-clean-api/src/infra/database/local"
)

func main() {
	db := local_database.Init()
	err := migrations.Migrate(db)
	if err != nil {
		log.Panic("Error migrating database", err)
	}

	echoServer := echo.New()
	router.Init(echoServer)

	echoServer.Logger.Fatal(echoServer.Start(":1323"))
}
