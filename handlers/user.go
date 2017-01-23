package handlers

import (
	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/handlers"
	"github.com/jakelong95/TownCenter/helpers"
	"github.com/jakelong95/TownCenter/models"
)

type UserI interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type User struct {
	*handlers.BaseHandler
	Helper helpers.UserI
}

func NewUser(ctx *handlers.GatewayContext) UserI {
	stats := ctx.Stats.Clone(statsd.Prefix("api.user"))
	return &User{
		BaseHandler: &handlers.BaseHandler{Stats: stats},
		Helper:      helpers.NewUser(ctx.Sql),
	}
}

func (u *User) New(ctx *gin.Context) {
	//Bind the json to a user object
	var json models.User
	err := ctx.BindJSON(&json)
	if err != nil {
		u.UserError(ctx, "Error: Unable to parse json", err)
		return
	}

	//Create the new user in the database
	user := models.NewUser(json.PassHash, json.FirstName, json.LastName, json.Email, json.Phone,
		                   json.AddressLine1, json.AddressLine2, json.AddressCity, json.AddressState, json.AddressZip,
		                   json.AddressCountry)
	err = u.Helper.Insert(user)
	if err != nil {
		u.ServerError(ctx, err, json)
		return
	}

	//Don't need to pass the password hash back
	user.PassHash = ""

	u.Success(ctx, user)
}

func (u *User) ViewAll(ctx *gin.Context) {
	//TODO

	u.Success(ctx, nil)
}

func (u *User) View(ctx *gin.Context) {
	userId := ctx.Param("userId")
	
	//Query the database for the user
	user, err := u.Helper.GetByID(userId)
	if err != nil {
		u.ServerError(ctx, err, userId)
		return
	}

	u.Success(ctx, user)
}

func (u *User) Update(ctx *gin.Context) {
	//TODO

	u.Success(ctx, nil)
}

func (u *User) Delete(ctx *gin.Context) {
	//TODO

	u.Success(ctx, nil)
}

func (u *User) Login(ctx *gin.Context) {
	//TODO: ALso return user id

	u.Success(ctx, nil)
}