package helpers

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/jakelong95/TownCenter/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInsertResetSuccess(t *testing.T) {
	assert := assert.New(t)

	s, mock, _ := sqlmock.New()
	r := getMockReset(s)
	token := getDefaultToken()

	mock.ExpectPrepare("INSERT INTO token").
		ExpectExec().
		WithArgs(token.Value, token.Email, token.CreatedAt, string(token.Status), token.Value, token.CreatedAt, string(token.Status)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Insert(token)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestGetTokenSuccess(t *testing.T) {
	assert := assert.New(t)

	s, mock, _ := sqlmock.New()
	r := getMockReset(s)
	token := getDefaultToken()

	mock.ExpectQuery("SELECT value, email, createdAt, status FROM token").
		WithArgs(token.Value).
		WillReturnRows(getResetMockRows().
			AddRow(token.Value, token.Email, token.CreatedAt, token.Status))

	res, err := r.Get(token.Value)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.NotNil(res)
}

func TestGetTokenEmpty(t *testing.T) {
	assert := assert.New(t)

	s, mock, _ := sqlmock.New()
	r := getMockReset(s)
	token := getDefaultToken()

	mock.ExpectQuery("SELECT value, email, createdAt, status FROM token").
		WithArgs(token.Value).
		WillReturnRows(getResetMockRows())

	res, err := r.Get(token.Value)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Nil(res)
}

func TestGetTokenError(t *testing.T) {
	assert := assert.New(t)

	s, mock, _ := sqlmock.New()
	r := getMockReset(s)
	token := getDefaultToken()

	mock.ExpectQuery("SELECT value, email, createdAt, status FROM token").
		WithArgs(token.Value).
		WillReturnError(fmt.Errorf("some error"))

	res, err := r.Get(token.Value)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
	assert.Nil(res)
}

func TestSetStatusSuccess(t *testing.T) {
	assert := assert.New(t)

	s, mock, _ := sqlmock.New()
	r := getMockReset(s)
	token := getDefaultToken()

	mock.ExpectPrepare("UPDATE token SET").
		ExpectExec().
		WithArgs(token.Email, string(models.INVALID)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := r.SetStatus(token, models.INVALID)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.EqualValues(models.INVALID, res.Status)
}

func TestSetStatusError(t *testing.T) {
	assert := assert.New(t)

	s, mock, _ := sqlmock.New()
	r := getMockReset(s)
	token := getDefaultToken()

	mock.ExpectPrepare("UPDATE token SET").
		ExpectExec().
		WithArgs(token.Email, string(models.INVALID)).
		WillReturnError(fmt.Errorf("some error"))

	res, err := r.SetStatus(token, models.INVALID)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
	assert.EqualValues(token.Status, res.Status)
}

func getDefaultToken() *models.Token {
	return models.NewToken("test")
}

func getResetMockRows() sqlmock.Rows {
	return sqlmock.NewRows([]string{"value", "email", "createdAt", "status"})
}

func getMockReset(s *sql.DB) *Reset {
	return NewReset(&gateways.MySQL{DB: s})
}
