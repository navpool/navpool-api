package main

import (
	"github.com/NavPool/navpool-api/app/config"
	"github.com/NavPool/navpool-api/app/database"
	"github.com/NavPool/navpool-api/app/database/migrations"
	"github.com/NavPool/navpool-api/app/routes"
	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	setReleaseMode()

	err := database.CreateConnection()
	if err != nil {
		log.Fatal(err)
	}
	migrations.Migrate()

	if config.Get().Sentry.Active {
		raven.SetDSN(config.Get().Sentry.DSN)
	}

	r := routes.Routes()
	r.Run(":" + config.Get().Server.Port)
}

func setReleaseMode() {
	if config.Get().Debug == true {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	log.Printf("Mode: %s", gin.Mode())
}
