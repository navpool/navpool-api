package tests

import (
	"bytes"
	"github.com/DATA-DOG/go-sqlmock"
	uuid "github.com/satori/go.uuid"
	"regexp"

	//"database/sql"
	"github.com/NavPool/navpool-api/app/middleware/authorization"
	"github.com/NavPool/navpool-api/app/routes"
	"github.com/NavPool/navpool-api/tests/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	//"regexp"
	"testing"
)

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_get_info_will_return_401_with_invalid_authorization() {
	router := routes.Routes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/info", nil)
	router.ServeHTTP(w, req)

	assert.Equal(s.T(), 401, w.Code)
	assert.Equal(s.T(), "{\"message\":\"Authentication failed\",\"status\":401}", w.Body.String())
}

func (s *Suite) Test_get_info_will_return_200_with_valid_authorization() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "accounts" WHERE ("accounts"."username" = $1) ORDER BY "accounts"."id" ASC LIMIT 1`)).
		WithArgs("tester").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "secret", "active"}).
			AddRow(uuid.NewV4().String(), "tester", "secret", true))

	router := routes.Routes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/info", bytes.NewBufferString(""))

	authorization.GetMessage(req.Method, req.RequestURI, authorization.GetBodyMd5(helpers.GetBody(req)))
	req.Header.Set("Authorization", "hmac tester:0e7abf7ad2b1eb52ed7a2ffe8849447969e8202ea1069be257bb39c4722ab17f")
	router.ServeHTTP(w, req)

	assert.Equal(s.T(), 200, w.Code)
	assert.Equal(s.T(), "{\"message\":\"Authentication failed\",\"status\":401}", w.Body.String())
}
