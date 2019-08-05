package tests

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NavPool/navpool-api/app/database"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log"
)

type Suite struct {
	suite.Suite
	mock sqlmock.Sqlmock
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	require.NoError(s.T(), err)

	database.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	database.DB.LogMode(true)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}
