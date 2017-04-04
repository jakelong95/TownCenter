package handlers

import (
	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/handlers"
	"github.com/jakelong95/TownCenter/helpers"
	"github.com/jakelong95/TownCenter/models"

	"github.com/pborman/uuid"
)

type RoasterI interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Upload(ctx *gin.Context)
	Time() gin.HandlerFunc
	GetJWT() gin.HandlerFunc
}

type Roaster struct {
	*handlers.BaseHandler
	Helper     helpers.RoasterI
	UserHelper helpers.UserI
}

type RoasterInfo struct {
	Roaster models.Roaster `json:"roaster"`
	UserID  uuid.UUID      `json:"userId"`
}

func NewRoaster(ctx *handlers.GatewayContext) RoasterI {
	stats := ctx.Stats.Clone(statsd.Prefix("api.roaster"))
	return &Roaster{
		BaseHandler: &handlers.BaseHandler{Stats: stats},
		Helper:      helpers.NewRoaster(ctx.Sql, ctx.S3),
		UserHelper:  helpers.NewUser(ctx.Sql, ctx.S3),
	}
}

func (r *Roaster) New(ctx *gin.Context) {
	//Bind the json to a roaster object
	var json RoasterInfo
	err := ctx.BindJSON(&json)
	if err != nil {
		r.UserError(ctx, "Error: Unable to parse json", err)
		return
	}

	//Create the new roaster in the database
	roaster := models.NewRoaster(json.Roaster.Name, json.Roaster.Email, json.Roaster.Phone, json.Roaster.AddressLine1, json.Roaster.AddressLine2, json.Roaster.AddressCity, json.Roaster.AddressState, json.Roaster.AddressZip, json.Roaster.AddressCountry)
	err = r.Helper.Insert(roaster)
	if err != nil {
		r.ServerError(ctx, err, json)
		return
	}

	//Update the user with their roaster ID
	user, err := r.UserHelper.GetByID(json.UserID.String())

	if err != nil {
		r.ServerError(ctx, err, json)
		return
	}

	if user == nil {
		r.NotFoundError(ctx, "Error: User with ID "+json.UserID.String()+" does not exist")
		return
	}

	user.RoasterId = roaster.ID
	user.PassHash = ""
	err = r.UserHelper.Update(user, json.UserID.String())

	if err != nil {
		r.ServerError(ctx, err, json.UserID)
	}

	r.Success(ctx, roaster)
}

func (r *Roaster) ViewAll(ctx *gin.Context) {
	//Use paging when getting lists of roasters
	offset, limit := r.GetPaging(ctx)

	//Query the database for all roasters
	roasters, err := r.Helper.GetAll(offset, limit)
	if err != nil {
		r.ServerError(ctx, err, roasters)
		return
	}

	r.Success(ctx, roasters)
}

func (r *Roaster) View(ctx *gin.Context) {
	roasterId := ctx.Param("roasterId")

	//Query the database for the roaster
	roaster, err := r.Helper.GetByID(roasterId)
	if err != nil {
		r.ServerError(ctx, err, roasterId)
		return
	}

	if roaster == nil {
		r.NotFoundError(ctx, "Error: Roaster with ID "+roasterId+" does not exist")
		return
	}

	r.Success(ctx, roaster)
}

func (r *Roaster) Update(ctx *gin.Context) {
	roasterId := ctx.Param("roasterId")

	//Bind the json to a roaster object
	var json models.Roaster
	err := ctx.BindJSON(&json)
	if err != nil {
		r.UserError(ctx, "Error: Unable to parse json", err)
		return
	}

	//Update the roaster in the database
	err = r.Helper.Update(&json, roasterId)
	if err != nil {
		r.ServerError(ctx, err, roasterId)
		return
	}

	r.Success(ctx, json)
}

func (r *Roaster) Upload(ctx *gin.Context) {
	id := ctx.Param("roasterId")
	file, headers, err := ctx.Request.FormFile("profile")
	if err != nil {
		r.ServerError(ctx, err, nil)
		return
	}
	if file == nil {
		r.UserError(ctx, "ERROR: unable to find body", nil)
		return
	}
	defer file.Close()

	err = r.Helper.Profile(id, headers.Filename, file)
	if err != nil {
		r.ServerError(ctx, err, id)
		return
	}

	r.Success(ctx, nil)
}

func (r *Roaster) Delete(ctx *gin.Context) {
	roasterId := ctx.Param("roasterId")

	//Delete the roaster from the database
	err := r.Helper.Delete(roasterId)
	if err != nil {
		r.ServerError(ctx, err, roasterId)
		return
	}

	r.Success(ctx, nil)
}
