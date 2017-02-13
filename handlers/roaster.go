package handlers

import (
	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/handlers"
	"github.com/jakelong95/TownCenter/helpers"
	"github.com/jakelong95/TownCenter/models"
)

type RoasterI interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type Roaster struct {
	*handlers.BaseHandler
	Helper helpers.RoasterI
}

func NewRoaster(ctx *handlers.GatewayContext) RoasterI {
	stats := ctx.Stats.Clone(statsd.Prefix("api.roaster"))
	return &Roaster{
		BaseHandler: &handlers.BaseHandler{Stats: stats},
		Helper:      helpers.NewRoaster(ctx.Sql),
	}
}

func (r *Roaster) New(ctx *gin.Context) {
	//Bind the json to a roaster object
	var json models.Roaster
	err := ctx.BindJSON(&json)
	if err != nil {
		r.UserError(ctx, "Error: Unable to parse json", err)
		return
	}

	//Create the new roaster in the database
	roaster := models.NewRoaster(json.Name, json.Email, json.Phone, json.AddressLine1, json.AddressLine2, json.AddressCity, json.AddressState, json.AddressZip, json.AddressCountry)
	err = r.Helper.Insert(roaster)
	if err != nil {
		r.ServerError(ctx, err, json)
		return
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
		r.UserError(ctx, "Error: Roaster does not exist", roasterId)
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