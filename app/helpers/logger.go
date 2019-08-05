package helpers

import (
	"github.com/NavPool/navpool-api/app/config"
	"github.com/getsentry/raven-go"
	"log"
)

func LogError(err error) {
	log.Print(err)
	if config.Get().Sentry.Active {
		raven.CaptureErrorAndWait(err, nil)
	}
}
