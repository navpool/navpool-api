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

	database.CreateConnection()
	migrations.Migrate()

	if config.Get().Sentry.Active {
		raven.SetDSN(config.Get().Sentry.DSN)
	}

	r := routes.Routes()
	r.Run(":" + config.Get().Server.Port)
}

func setReleaseMode() {
	if config.Get().Debug == false {
		log.Printf("Mode: %s", gin.ReleaseMode)
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.Printf("Mode: %s", gin.DebugMode)
		gin.SetMode(gin.DebugMode)
	}
}
