package authorization

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/NavPool/navpool-api/app/container"
	"github.com/NavPool/navpool-api/app/helpers"
	"github.com/NavPool/navpool-api/app/model/account"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func Authorization(c *gin.Context) {
	username, digest, err := getAuthorization(c.GetHeader("Authorization"))
	if err != nil {
		helpers.HandleError(c, err, 401)
		return
	}

	body, _ := c.GetRawData()
	message := getMessage(c.Request.Method, c.Request.RequestURI, getBodyMd5(body))
	log.Print(message)

	account, err := model_account.AccountRepository().GetByUsername(username)
	if err != nil {
		log.Printf("Invalid secret (uername:%s)", username)
		helpers.HandleError(c, ErrAuthenticationFailed, 401)
		return
	}

	expectedMAC := getExpectedMAC(message, account.Secret)
	actualMAC, _ := hex.DecodeString(digest)

	log.Print(hex.EncodeToString(expectedMAC))

	if !hmac.Equal(expectedMAC, actualMAC) {
		log.Printf("Invalid HMAC (uername:%s)", username)
		helpers.HandleError(c, ErrAuthenticationFailed, 401)
		return
	}

	container.Container.Account = *account
}

func getAuthorization(authHeader string) (username string, digest string, err error) {
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "hmac" {
		log.Printf("Authorization header malformed (%s)", authHeader)
		err = ErrAuthenticationFailed
		return
	}

	parts = strings.Split(parts[1], ":")
	if len(parts) != 2 {
		log.Printf("Authorization header malformed (%s)", authHeader)
		err = ErrAuthenticationFailed
		return
	}

	return parts[0], parts[1], nil
}

func getMessage(method string, uri string, body string) string {
	return strings.Trim(fmt.Sprintf("%s+%s+%s", method, uri, body), "+")
}

func getExpectedMAC(message string, secret string) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))

	return mac.Sum(nil)
}

func getBodyMd5(body []byte) string {
	hasher := md5.New()
	hasher.Write(body)

	return hex.EncodeToString(hasher.Sum(nil))
}

var (
	ErrAuthenticationFailed = errors.New("Authentication failed")
)
