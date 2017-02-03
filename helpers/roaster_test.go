package helpers

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/jakelong95/TownCenter/models"
	"github.com/ghmeier/bloodlines/gateways"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRoasterGetByID(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	s, mock, _ := sqlmock.New()
	r := getMockRoaster(s)

	mock.ExpectQuery("SELECT id, name, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry FROM roaster").
		WithArgs(id.String()).
		WillReturnRows(getRoasterMockRows().AddRow(id.String(), "Name", "Email", "Phone", "AddressLine1", "AddressLine2", "AddressCity", "AddressState", "AddressZip", "AddressCountry"))

	roaster, err := r.GetByID(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(roaster.ID, id)
	assert.Equal(roaster.Name, "Name")
	assert.Equal(roaster.Email, "Email")
	assert.Equal(roaster.Phone, "Phone")
	assert.Equal(roaster.AddressLine1, "AddressLine1")
	assert.Equal(roaster.AddressLine2, "AddressLine2")
	assert.Equal(roaster.AddressCity, "AddressCity")
	assert.Equal(roaster.AddressState, "AddressState")
	assert.Equal(roaster.AddressZip, "AddressZip")
	assert.Equal(roaster.AddressCountry, "AddressCountry")
}

func TestRoasterGetByIDError(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	s, mock, _ := sqlmock.New()
	r := getMockRoaster(s)

	mock.ExpectQuery("SELECT id, name, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry FROM roaster").
		WithArgs(id.String()).
		WillReturnError(fmt.Errorf("This is an error"))

	_, err := r.GetByID(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestRoasterGetAll(t *testing.T) {
	assert := assert.New(t)

	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	r := getMockRoaster(s)

	mock.ExpectQuery("SELECT id, name, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry FROM roaster").
		WithArgs(offset, limit).
		WillReturnRows(getRoasterMockRows().
			AddRow(uuid.New(), "Name", "Email", "Phone", "AddressLine1", "AddressLine2", "AddressCity", "AddressState", "AddressZip", "AddressCountry").
			AddRow(uuid.New(), "Name", "Email", "Phone", "AddressLine1", "AddressLine2", "AddressCity", "AddressState", "AddressZip", "AddressCountry"))

	roasters, err := r.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(2, len(roasters))
}

func TestRoasterGetAllError(t *testing.T) {
	assert := assert.New(t)

	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	r := getMockRoaster(s)

	mock.ExpectQuery("SELECT id, name, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry FROM roaster").
		WithArgs(offset, limit).
		WillReturnError(fmt.Errorf("This is an error"))

	_, err := r.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestRoasterInsert(t *testing.T) {
	assert := assert.New(t)

	roaster := getDefaultRoaster()
	s, mock, _ := sqlmock.New()
	r := getMockRoaster(s)

	mock.ExpectPrepare("INSERT INTO roaster").
		ExpectExec().
		WithArgs(roaster.ID.String(), roaster.Name, roaster.Email, roaster.Phone, roaster.AddressLine1, roaster.AddressLine2, roaster.AddressCity, roaster.AddressState, roaster.AddressZip, roaster.AddressCountry).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Insert(roaster)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestRoasterInsertError(t *testing.T) {
	assert := assert.New(t)

	roaster := getDefaultRoaster()
	s, mock, _ := sqlmock.New()
	r := getMockRoaster(s)

	mock.ExpectPrepare("INSERT INTO roaster").
		ExpectExec().
		WithArgs(roaster.ID.String(), roaster.Name, roaster.Email, roaster.Phone, roaster.AddressLine1, roaster.AddressLine2, roaster.AddressCity, roaster.AddressState, roaster.AddressZip, roaster.AddressCountry).
		WillReturnError(fmt.Errorf("This is an error"))

	err := r.Insert(roaster)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestRoasterUpdate(t *testing.T) {
	assert := assert.New(t)

	roaster := getDefaultRoaster()
	s, mock, _ := sqlmock.New()
	r := getMockRoaster(s)

	mock.ExpectPrepare("UPDATE roaster").
		ExpectExec().
		WithArgs(roaster.Name, roaster.Email, roaster.Phone, roaster.AddressLine1, roaster.AddressLine2, roaster.AddressCity, roaster.AddressState, roaster.AddressZip, roaster.AddressCountry, roaster.ID.String()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Update(roaster, roaster.ID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestRoasterUpdateError(t *testing.T) {
	assert := assert.New(t)

	roaster := getDefaultRoaster()
	s, mock, _ := sqlmock.New()
	r := getMockRoaster(s)

	mock.ExpectPrepare("UPDATE roaster").
		ExpectExec().
		WithArgs(roaster.Name, roaster.Email, roaster.Phone, roaster.AddressLine1, roaster.AddressLine2, roaster.AddressCity, roaster.AddressState, roaster.AddressZip, roaster.AddressCountry, roaster.ID.String()).
		WillReturnError(fmt.Errorf("This is an error"))

	err := r.Update(roaster, roaster.ID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func getDefaultRoaster() *models.Roaster {
	return models.NewRoaster("Name", "Email", "Phone", "AddressLine1", "AddressLine2", "AddressCity", "AddressState", "AddressZip", "AddressCountry")
}

func getRoasterMockRows() sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "name", "email", "phone", "addressLine1", "addressLine2", "addressCity", "addressState", "addressZip", "addressCountry"})
}

func getMockRoaster(s *sql.DB) *Roaster {
	return NewRoaster(&gateways.MySQL{DB: s})
}