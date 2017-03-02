package helpers

import (
	"gopkg.in/alexcesaro/statsd.v2"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/jakelong95/TownCenter/models"
)

type baseHelper struct {
	sql   gateways.SQL
	stats *statsd.Client
}

type UserI interface {
	GetByID(string) (*models.User, error)
	GetAll(int, int) ([]*models.User, error)
	Insert(*models.User) error
	Update(*models.User, string) error
	Delete(string) error
	GetByEmail(string) (*models.User, error)
}

type User struct {
	*baseHelper
}

func NewUser(sql gateways.SQL) *User {
	return &User{baseHelper: &baseHelper{sql: sql}}
}

func (u *User) GetByID(id string) (*models.User, error) {
	rows, err := u.sql.Select("SELECT id, passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry, roasterId FROM user WHERE id=?", id)

	if err != nil {
		return nil, err
	}

	users, err := models.UserFromSQL(rows)
	if err != nil {
		return nil, err
	}

	if(len(users) == 0) {
		return nil, nil
	}

	return users[0], err
}

func (u *User) GetAll(offset int, limit int) ([]*models.User, error) {
	rows, err := u.sql.Select("SELECT id, passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry, roasterId FROM user ORDER BY id ASC LIMIT ?,?", offset, limit)
	if err != nil {
		return nil, err
	}

	users, err := models.UserFromSQL(rows)
	if err != nil {
		return nil, err
	}

	return users, err
}

func (u *User) Insert(user *models.User) error {
	err := u.sql.Modify(
		"INSERT INTO user (id, passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry, roasterId) VALUE (?,?,?,?,?,?,?,?,?,?,?,?,?)",
		user.ID,
		user.PassHash,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.AddressLine1,
		user.AddressLine2,
		user.AddressCity,
		user.AddressState,
		user.AddressZip,
		user.AddressCountry,
		user.RoasterId,
	)

	return err
}

func (u *User) Update(user *models.User, id string) error {
	err := u.sql.Modify(
		"UPDATE user SET firstName=?, lastName=?, email=?, phone=?, addressLine1=?, addressLine2=?, addressCity=?, addressState=?, addressZip=?, addressCountry=?, roasterId=? WHERE id=?",
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.AddressLine1,
		user.AddressLine2,
		user.AddressCity,
		user.AddressState,
		user.AddressZip,
		user.AddressCountry,
		user.RoasterId,
		id,
	)

	if err != nil {
		return nil
	}

	if user.PassHash != "" {
		err = u.sql.Modify(
			"UPDATE user SET passHash=? WHERE id=?",
			user.PassHash,
			id,
		)
	}

	return err
}

func (u *User) Delete(id string) error {
	err := u.sql.Modify("DELETE FROM user WHERE id=?", id)
	return err
}

func (u *User) GetByEmail(email string) (*models.User, error) {
	rows, err := u.sql.Select("SELECT id, passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry, roasterId FROM user WHERE email=?", email)
	if err != nil {
		return nil, err
	}

	users, err := models.UserFromSQL(rows)
	if err != nil {
		return nil, err
	}

	if(len(users) == 0) {
		return nil, nil
	}

	return users[0], err
}
