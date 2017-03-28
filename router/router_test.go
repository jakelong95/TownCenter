package router

import (
	"testing"

	mockg "github.com/ghmeier/bloodlines/_mocks/gateways"
	"github.com/ghmeier/bloodlines/config"
	h "github.com/ghmeier/bloodlines/handlers"
	m "github.com/ghmeier/bloodlines/models"
	mocks "github.com/jakelong95/TownCenter/_mocks"
	"github.com/jakelong95/TownCenter/handlers"
	"github.com/jakelong95/TownCenter/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
		reset:   handlers.NewReset(ctx),
	}
}

func mockUser() (*TownCenter, *mocks.UserI) {
	t := getMockTownCenter()
	userMock := new(mocks.UserI)
	bloodlines := new(mockg.Bloodlines)
	bloodlines.On("NewPreference", mock.AnythingOfType("uuid.UUID")).Return(&m.Preference{}, nil)

	t.user = &handlers.User{
		Helper:      userMock,
		BaseHandler: &h.BaseHandler{Stats: nil},
		Bloodlines:  bloodlines,
	}

	InitRouter(t)

	return t, userMock
}

func mockRoaster() (*TownCenter, *mocks.RoasterI) {
	t := getMockTownCenter()
	roasterMock := new(mocks.RoasterI)
	userHelper := new(mocks.UserI)
	userHelper.On("GetByID", mock.AnythingOfType("string")).Return(&models.User{}, nil)
	userHelper.On("Update", mock.AnythingOfType("*models.User"), mock.AnythingOfType("string")).Return(nil)

	t.roaster = &handlers.Roaster{
		Helper:      roasterMock,
		BaseHandler: &h.BaseHandler{Stats: nil},
		UserHelper:  userHelper,
	}
	InitRouter(t)

	return t, roasterMock
}

func mockReset() (*TownCenter, *mocks.UserI, *mocks.ResetI) {
	t := getMockTownCenter()
	userHelper := new(mocks.UserI)
	resetMock := new(mocks.ResetI)

	t.reset = &handlers.Reset{
		User:        userHelper,
		BaseHandler: &h.BaseHandler{Stats: nil},
		Reset:       resetMock,
	}
	InitRouter(t)

	return t, userHelper, resetMock
}
