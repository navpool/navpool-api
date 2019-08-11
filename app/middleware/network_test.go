package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_it_will_default_to_the_main_network(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(NetworkSelect)
	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "") })

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "mainnet", w.Header().Get("X-Network"))
}

func Test_it_will_alert_when_network_header_is_invalid(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(NetworkSelect)
	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "") })

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("network", "invalid")
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"message\":\"Network not found\",\"status\":400}", w.Body.String())
}

func Test_it_can_select_the_network_from_the_header_successfully(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(NetworkSelect)
	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "") })

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("network", "testnet")
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "testnet", w.Header().Get("X-Network"))
}
