package middleware

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/NavPool/navpool-api/config"
	"github.com/NavPool/navpool-api/helpers"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func Authentication(c *gin.Context) {
	username, digest, err := getAuthentication(c.GetHeader("Authentication"))
	if err != nil {
		helpers.HandleError(c, err, 401)
		return
	}

	message := getMessage(c)
	log.Print(message)
	expectedMAC := getExpectedMAC(message, getSecret(username))
	actualMAC, _ := hex.DecodeString(digest)

	log.Print(hex.EncodeToString(expectedMAC))

	if !hmac.Equal(expectedMAC, actualMAC) {
		log.Printf("Invalid HMAC (uername:%s)", username)
		helpers.HandleError(c, ErrAuthenticationFailed, 401)
		return
	}

	config.Get().Account = username
}

func getAuthentication(authHeader string) (username string, digest string, err error) {
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "hmac" {
		log.Printf("Authentication header malformed (%s)", authHeader)
		err = ErrAuthenticationFailed
		return
	}

	parts = strings.Split(parts[1], ":")
	if len(parts) != 2 {
		log.Printf("Authentication header malformed (%s)", authHeader)
		err = ErrAuthenticationFailed
		return
	}

	return parts[0], parts[1], nil
}

func getMessage(c *gin.Context) string {
	return strings.Trim(fmt.Sprintf("%s+%s+%s", c.Request.Method, c.Request.RequestURI, getBodyMd5(c)), "+")
}

func getSecret(username string) string {
	return "secret"
}

func getExpectedMAC(message string, secret string) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))

	return mac.Sum(nil)
}

func getBodyMd5(c *gin.Context) string {
	body, _ := c.GetRawData()
	if len(body) == 0 {
		return ""
	}

	hasher := md5.New()
	hasher.Write(body)

	return hex.EncodeToString(hasher.Sum(nil))
}

var (
	ErrAuthenticationFailed = errors.New("Authentication failed")
)
