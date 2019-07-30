module github.com/NavPool/navpool-api

go 1.12

require (
	github.com/certifi/gocertifi v0.0.0-20190506164543-d2eda7129713 // indirect
	github.com/getsentry/raven-go v0.2.0
	github.com/gin-contrib/gzip v0.0.1
	github.com/gin-gonic/gin v1.4.0
	github.com/jinzhu/gorm v1.9.10
	github.com/kr/pretty v0.1.0 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.3.2
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c
	gopkg.in/go-playground/validator.v8 v8.18.2
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
