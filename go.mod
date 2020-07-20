module github.com/NavPool/navpool-api

go 1.13

require (
	github.com/certifi/gocertifi v0.0.0-20190506164543-d2eda7129713 // indirect
	github.com/getsentry/raven-go v0.2.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-contrib/gzip v0.0.1
	github.com/gin-gonic/gin v1.5.0
	github.com/joho/godotenv v1.3.0
	github.com/pkg/errors v0.8.1 // indirect
	github.com/sarulabs/dingo/v3 v3.1.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.3.2 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
