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
	Update(*models.User) error	
}

type User struct {
	*baseHelper
}

func NewUser(sql gateways.SQL) *User {
	return &User{baseHelper: &baseHelper{sql: sql}}
}

func (u *User) GetByID(id string) (*models.User, error) {
	rows, err := u.sql.Select("SELECT * FROM user WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	users, err := models.UserFromSQL(rows)
	if err != nil {
		return nil, err
	}

	//Don't pass the password hash around
	users[0].PassHash = ""

	return users[0], err
}

func (u *User) GetAll(offset int, limit int) ([]*models.User, error) {
	rows, err := u.sql.Select("SELECT * FROM user ORDER BY id ASC LIMIT ?,?", offset, limit)
	if err != nil {
		return nil, err
	}

	users, err := models.UserFromSQL(rows)
	if err != nil {
		return nil, err
	}

	//Don't pass the password hash around
	for _, user := range users {
		user.PassHash = ""
	}

	return users, err
}

func (u *User) Insert(user *models.User) error {
	err := u.sql.Modify(
		"INSERT INTO user (id, passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry, roasterId, isRoaster) VALUE (?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
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
		user.IsRoaster,
	)

	return err
}

func (u *User) Update(user *models.User) error {
	err := u.sql.Modify(
		"UPDATE user SET passHash=?, firstName=?, lastName=?, email=?, phone=?, addressLine1=?, addressLine2=?, addressCity=?, addressState=?, addressZip=?, addressCountry=?, roasterId=?, isRoaster=? WHERE id=?",
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
		user.IsRoaster,
		user.ID,
	)

	return err
}