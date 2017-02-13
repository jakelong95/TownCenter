package helpers

import (
	"github.com/ghmeier/bloodlines/gateways"
	"github.com/jakelong95/TownCenter/models"
)

type RoasterI interface {
	GetByID(string) (*models.Roaster, error)
	GetAll(int, int) ([]*models.Roaster, error)
	Insert(*models.Roaster) error
	Update(*models.Roaster, string) error	
	Delete(string) error
}

type Roaster struct {
	*baseHelper
}

func NewRoaster(sql gateways.SQL) *Roaster {
	return &Roaster{baseHelper: &baseHelper{sql: sql}}
}

func (r *Roaster) GetByID(id string) (*models.Roaster, error) {
	rows, err := r.sql.Select("SELECT id, name, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry FROM roaster WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	roasters, err := models.RoasterFromSQL(rows)
	if err != nil {
		return nil, err
	}

	if(len(roasters) == 0) {
		return nil, nil
	}

	return roasters[0], err
}

func (r *Roaster) GetAll(offset int, limit int) ([]*models.Roaster, error) {
	rows, err := r.sql.Select("SELECT id, name, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry FROM roaster ORDER BY id ASC LIMIT ?,?", offset, limit)
	if err != nil {
		return nil, err
	}

	roasters, err := models.RoasterFromSQL(rows)
	if err != nil {
		return nil, err
	}

	return roasters, err
}

func (r *Roaster) Insert(roaster *models.Roaster) error {
	err := r.sql.Modify(
		"INSERT INTO roaster (id, name, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry) VALUE (?,?,?,?,?,?,?,?,?,?)",
		roaster.ID, 
		roaster.Name,
		roaster.Email, 
		roaster.Phone, 
		roaster.AddressLine1, 
		roaster.AddressLine2, 
		roaster.AddressCity, 
		roaster.AddressState, 
		roaster.AddressZip, 
		roaster.AddressCountry, 
	)

	return err
}

func (r *Roaster) Update(roaster *models.Roaster, roasterId string) error {
	err := r.sql.Modify(
		"UPDATE roaster SET name=?, email=?, phone=?, addressLine1=?, addressLine2=?, addressCity=?, addressState=?, addressZip=?, addressCountry=? WHERE id=?",
		roaster.Name,
		roaster.Email, 
		roaster.Phone, 
		roaster.AddressLine1, 
		roaster.AddressLine2, 
		roaster.AddressCity, 
		roaster.AddressState, 
		roaster.AddressZip, 
		roaster.AddressCountry, 
		roasterId,
	)

	return err
}

func (r *Roaster) Delete(id string) error {
	err := r.sql.Modify("DELETE FROM roaster WHERE id=?", id)
	return err
}