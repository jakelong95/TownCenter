package models

import (
	"database/sql"
	"time"

	"github.com/pborman/uuid"
)

type Token struct {
	Value     string      `json:"value"`
	Email     string      `json:"email"`
	CreatedAt time.Time   `json:"createdAt"`
	Status    TokenStatus `json:"status"`
}

type ResetRequest struct {
	PassHash string `json:"passHash"`
}

func NewToken(email string) *Token {
	return &Token{
		Value:     uuid.New(),
		Email:     email,
		CreatedAt: time.Now(),
		Status:    ACTIVE,
	}
}

func TokenFromSQL(rows *sql.Rows) ([]*Token, error) {
	tokens := make([]*Token, 0)

	for rows.Next() {
		t := &Token{}

		var raw string
		rows.Scan(&t.Value, &t.Email, &t.CreatedAt, &raw)

		s, ok := toTokenStatus(raw)
		if !ok {
			s = INVALID
		}

		t.Status = s
		tokens = append(tokens, t)
	}

	return tokens, nil
}

func toTokenStatus(s string) (TokenStatus, bool) {
	switch s {
	case ACTIVE:
		return ACTIVE, true
	case INVALID:
		return INVALID, true
	case EXPIRED:
		return EXPIRED, true
	default:
		return "INVALID", false
	}
}

/*TokenStatus is an enum wrapper for valid TokenStatus strings*/
type TokenStatus string

/*valid TokenStatuses*/
const (
	ACTIVE  = "ACTIVE"
	INVALID = "INVALID"
	EXPIRED = "EXPIRED"
)
