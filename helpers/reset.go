package helpers

import (
	"github.com/ghmeier/bloodlines/gateways"
	"github.com/jakelong95/TownCenter/models"
)

type ResetI interface {
	Insert(*models.Token) error
	Get(string) (*models.Token, error)
	SetStatus(*models.Token, models.TokenStatus) (*models.Token, error)
}

type Reset struct {
	*baseHelper
}

func NewReset(sql gateways.SQL) *Reset {
	return &Reset{
		baseHelper: &baseHelper{sql: sql},
	}
}

func (r *Reset) Insert(token *models.Token) error {
	err := r.sql.Modify("INSERT INTO token (value, email, createdAt, status) VALUE (?,?,?,?) ON DUPLICATE KEY UPDATE value=?, createdAt=?, status=?",
		token.Value,
		token.Email,
		token.CreatedAt,
		string(token.Status),
		token.Value,
		token.CreatedAt,
		string(token.Status),
	)

	return err
}

func (r *Reset) Get(value string) (*models.Token, error) {
	rows, err := r.sql.Select("SELECT value, email, createdAt, status FROM token WHERE value=?", value)
	if err != nil {
		return nil, err
	}

	// cannot return an error
	tokens, _ := models.TokenFromSQL(rows)

	if len(tokens) == 0 {
		return nil, nil
	}

	return tokens[0], nil
}

func (r *Reset) SetStatus(token *models.Token, status models.TokenStatus) (*models.Token, error) {
	err := r.sql.Modify("UPDATE token SET status=? WHERE email=?", token.Email, string(status))
	if err != nil {
		return token, err
	}

	token.Status = status
	return token, err
}
