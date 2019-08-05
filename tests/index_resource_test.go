package tests

import (
	"github.com/NavPool/navpool-api/app/routes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_it_will_return_the_name_on_the_index_page(t *testing.T) {
	router := routes.Routes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome to NavPool Node API!", w.Body.String())
}
