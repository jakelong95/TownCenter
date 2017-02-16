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

func TestUserGetByID(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	s, mock, _ := sqlmock.New()
	u := getMockUser(s)

	mock.ExpectQuery("SELECT id, passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry, roasterId, isRoaster FROM user").
		WithArgs(id.String()).
		WillReturnRows(getUserMockRows().AddRow(id.String(), "", "FirstName", "LastName", "Email", "Phone", "AddressLine1", "AddressLine2", "AddressCity", "AddressState", "AddressZip", "AddressCountry", nil, "0"))

	user, err := u.GetByID(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(user.ID, id)
	assert.Equal(user.PassHash, "")
	assert.Equal(user.FirstName, "FirstName")
	assert.Equal(user.LastName, "LastName")
	assert.Equal(user.Email, "Email")
	assert.Equal(user.Phone, "Phone")
	assert.Equal(user.AddressLine1, "AddressLine1")
	assert.Equal(user.AddressLine2, "AddressLine2")
	assert.Equal(user.AddressCity, "AddressCity")
	assert.Equal(user.AddressState, "AddressState")
	assert.Equal(user.AddressZip, "AddressZip")
	assert.Equal(user.AddressCountry, "AddressCountry")
	assert.Equal(user.RoasterId, uuid.UUID(nil))
	assert.Equal(user.IsRoaster, 0)
}

func TestUserGetByIDError(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	s, mock, _ := sqlmock.New()
	u := getMockUser(s)

	mock.ExpectQuery("SELECT id, passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry, roasterId, isRoaster FROM user").
		WithArgs(id.String()).
		WillReturnError(fmt.Errorf("This is an error"))

	_, err := u.GetByID(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestUserGetByIDDoesNotExist(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	s, mock, _ := sqlmock.New()
	u := getMockUser(s)

	mock.ExpectQuery("SELECT id, passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry, roasterId, isRoaster FROM user").
		WithArgs(id.String()).
		WillReturnRows(getUserMockRows())

	user, err := u.GetByID(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Nil(user)
	assert.NoError(err)
}

func TestUserGetByEmail(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	s, mock, _ := sqlmock.New()
	u := getMockUser(s)

	mock.ExpectQuery("SELECT id, passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry, roasterId, isRoaster FROM user").
		WithArgs("Email").
		WillReturnRows(getUserMockRows().AddRow(id.String(), "", "FirstName", "LastName", "Email", "Phone", "AddressLine1", "AddressLine2", "AddressCity", "AddressState", "AddressZip", "AddressCountry", nil, "0"))

	user, err := u.GetByEmail("Email")

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(user.ID, id)
	assert.Equal(user.PassHash, "")
	assert.Equal(user.FirstName, "FirstName")
	assert.Equal(user.LastName, "LastName")
	assert.Equal(user.Email, "Email")
	assert.Equal(user.Phone, "Phone")
	assert.Equal(user.AddressLine1, "AddressLine1")
	assert.Equal(user.AddressLine2, "AddressLine2")
	assert.Equal(user.AddressCity, "AddressCity")
	assert.Equal(user.AddressState, "AddressState")
	assert.Equal(user.AddressZip, "AddressZip")
	assert.Equal(user.AddressCountry, "AddressCountry")
	assert.Equal(user.RoasterId, uuid.UUID(nil))
	assert.Equal(user.IsRoaster, 0)
}

func TestUserGetByEmailError(t *testing.T) {
	assert := assert.New(t)

	s, mock, _ := sqlmock.New()
	u := getMockUser(s)

	mock.ExpectQuery("SELECT id, passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry, roasterId, isRoaster FROM user").
		WithArgs("Email").
		WillReturnError(fmt.Errorf("This is an error"))

	_, err := u.GetByEmail("Email")

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestUserGetAll(t *testing.T) {
	assert := assert.New(t)

	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	u := getMockUser(s)

	mock.ExpectQuery("SELECT id, passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry, roasterId, isRoaster FROM user").
		WithArgs(offset, limit).
		WillReturnRows(getUserMockRows().
			AddRow(uuid.New(), "PassHash", "FirstName", "LastName", "Email", "Phone", "AddressLine1", "AddressLine2", "AddressCity", "AddressState", "AddressZip", "AddressCountry", nil, "0").
			AddRow(uuid.New(), "PassHash", "FirstName", "LastName", "Email", "Phone", "AddressLine1", "AddressLine2", "AddressCity", "AddressState", "AddressZip", "AddressCountry", nil, "0"))

	users, err := u.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
	assert.Equal(2, len(users))
}

func TestUserGetAllError(t *testing.T) {
	assert := assert.New(t)

	offset, limit := 0, 20
	s, mock, _ := sqlmock.New()
	u := getMockUser(s)

	mock.ExpectQuery("SELECT id, passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry, roasterId, isRoaster FROM user").
		WithArgs(offset, limit).
		WillReturnError(fmt.Errorf("This is an error"))

	_, err := u.GetAll(offset, limit)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestUserInsert(t *testing.T) {
	assert := assert.New(t)

	user := getDefaultUser()
	s, mock, _ := sqlmock.New()
	u := getMockUser(s)

	mock.ExpectPrepare("INSERT INTO user").
		ExpectExec().
		WithArgs(user.ID.String(), user.PassHash, user.FirstName, user.LastName, user.Email, user.Phone, user.AddressLine1, user.AddressLine2, user.AddressCity, user.AddressState, user.AddressZip, user.AddressCountry, user.RoasterId.String(), user.IsRoaster).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := u.Insert(user)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestUserInsertError(t *testing.T) {
	assert := assert.New(t)

	user := getDefaultUser()
	s, mock, _ := sqlmock.New()
	u := getMockUser(s)

	mock.ExpectPrepare("INSERT INTO user").
		ExpectExec().
		WithArgs(user.ID.String(), user.PassHash, user.FirstName, user.LastName, user.Email, user.Phone, user.AddressLine1, user.AddressLine2, user.AddressCity, user.AddressState, user.AddressZip, user.AddressCountry, user.RoasterId.String(), user.IsRoaster).
		WillReturnError(fmt.Errorf("This is an error"))

	err := u.Insert(user)

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)

	user := getDefaultUser()
	s, mock, _ := sqlmock.New()
	u := getMockUser(s)

	mock.ExpectPrepare("UPDATE user").
		ExpectExec().
		WithArgs(user.PassHash, user.FirstName, user.LastName, user.Email, user.Phone, user.AddressLine1, user.AddressLine2, user.AddressCity, user.AddressState, user.AddressZip, user.AddressCountry, user.RoasterId.String(), user.IsRoaster, user.ID.String()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := u.Update(user, user.ID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestUpdateError(t *testing.T) {
	assert := assert.New(t)

	user := getDefaultUser()
	s, mock, _ := sqlmock.New()
	u := getMockUser(s)

	mock.ExpectPrepare("UPDATE user").
		ExpectExec().
		WithArgs(user.PassHash, user.FirstName, user.LastName, user.Email, user.Phone, user.AddressLine1, user.AddressLine2, user.AddressCity, user.AddressState, user.AddressZip, user.AddressCountry, user.RoasterId.String(), user.IsRoaster, user.ID.String()).
		WillReturnError(fmt.Errorf("This is an error"))

	err := u.Update(user, user.ID.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	s, mock, _ := sqlmock.New()
	u := getMockUser(s)

	mock.ExpectPrepare("DELETE FROM user").
		ExpectExec().
		WithArgs(id.String()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := u.Delete(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.NoError(err)
}

func TestDeleteUserError(t *testing.T) {
	assert := assert.New(t)

	id := uuid.NewUUID()
	s, mock, _ := sqlmock.New()
	u := getMockUser(s)

	mock.ExpectPrepare("DELETE FROM user").
		ExpectExec().
		WithArgs(id.String()).
		WillReturnError(fmt.Errorf("This is an error"))

	err := u.Delete(id.String())

	assert.Equal(mock.ExpectationsWereMet(), nil)
	assert.Error(err)
}

func getDefaultUser() *models.User {
	return models.NewUser("passhash", "Firstname", "Lastname", "Email", "Phone", "AddressLine1", "AddressLine2", "AddressCity", "AddressState", "AddressZip", "AddressCountry")
}

func getUserMockRows() sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "passHash", "firstName", "lastName", "email", "phone", "addressLine1", "addressLine2", "addressCity", "addressState", "addressZip", "addressCountry", "roasterId", "isRoaster"})
}

func getMockUser(s *sql.DB) *User {
	return NewUser(&gateways.MySQL{DB: s})
}