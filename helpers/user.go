package helpers

import (
	"fmt"
	"mime/multipart"

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
	Profile(string, string, multipart.File) error
}

type User struct {
	*baseHelper
	S3 gateways.S3
}

func NewUser(sql gateways.SQL, s3 gateways.S3) *User {
	return &User{
		baseHelper: &baseHelper{sql: sql},
		S3:         s3,
	}
}

func (u *User) GetByID(id string) (*models.User, error) {
	rows, err := u.sql.Select("SELECT id, passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry, roasterId, profileUrl FROM user WHERE id=?", id)

	if err != nil {
		return nil, err
	}

	users, err := models.UserFromSQL(rows)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users[0], err
}

func (u *User) GetAll(offset int, limit int) ([]*models.User, error) {
	rows, err := u.sql.Select("SELECT id, passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry, roasterId, profileUrl FROM user ORDER BY id ASC LIMIT ?,?", offset, limit)
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
		"INSERT INTO user (id, passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry, roasterId, profileUrl) VALUE (?,?,?,?,?,?,?,?,?,?,?,?,?)",
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
		user.ProfileURL,
	)

	return err
}

func (u *User) Update(user *models.User, id string) error {
	err := u.sql.Modify(
		"UPDATE user SET firstName=?, lastName=?, email=?, phone=?, addressLine1=?, addressLine2=?, addressCity=?, addressState=?, addressZip=?, addressCountry=?, roasterId=?, profileUrl=? WHERE id=?",
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
		user.ProfileURL,
		id,
	)

	if err != nil {
		return err
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
	rows, err := u.sql.Select("SELECT id, passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry, roasterId, profileUrl FROM user WHERE email=?", email)
	if err != nil {
		return nil, err
	}

	users, err := models.UserFromSQL(rows)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users[0], err
}

func (u *User) Profile(id string, name string, body multipart.File) error {
	filename := fmt.Sprintf("%s-%s", id, name)
	fmt.Println(filename)
	_, err := u.S3.Upload("profile", filename, body)
	return err
}
