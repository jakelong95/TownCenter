package router

import (
	"testing"

	mockg "github.com/ghmeier/bloodlines/_mocks/gateways"
	"github.com/ghmeier/bloodlines/config"
	h "github.com/ghmeier/bloodlines/handlers"
	"github.com/jakelong95/TownCenter/handlers"
	mocks "github.com/jakelong95/TownCenter/_mocks"

	"github.com/stretchr/testify/assert"
	"gopkg.in/alexcesaro/statsd.v2"
)

func TestNewSuccess(t *testing.T) {
	assert := assert.New(t)

	r, err := New(&config.Root{SQL: config.MySQL{}})

	assert.NoError(err)
	assert.NotNil(r)
}

func getMockTownCenter() *TownCenter {
	sql := new(mockg.SQL)
	stats, _ := statsd.New()
	ctx := &h.GatewayContext{
		Sql:   sql,
		Stats: stats,
	}

	return &TownCenter{
		user:    handlers.NewUser(ctx),
		roaster: handlers.NewRoaster(ctx), 
	}
}

func mockUser() (*TownCenter, *mocks.UserI) {
	t := getMockTownCenter()
	mock := new(mocks.UserI)
	t.user = &handlers.User{Helper: mock, BaseHandler: &h.BaseHandler{Stats: nil}}
	InitRouter(t)

	return t, mock
}

func mockRoaster() (*TownCenter, *mocks.RoasterI) {
	t := getMockTownCenter()
	mock := new(mocks.RoasterI)
	t.roaster = &handlers.Roaster{Helper: mock, BaseHandler: &h.BaseHandler{Stats: nil}}
	InitRouter(t)

	return t, mock
}