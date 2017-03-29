package handlers

import (
	"fmt"
	"time"

	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/handlers"
	bmodels "github.com/ghmeier/bloodlines/models"
	"github.com/jakelong95/TownCenter/helpers"
	"github.com/jakelong95/TownCenter/models"
)

type ResetI interface {
	Request(ctx *gin.Context)
	Get(ctx *gin.Context)
	Fulfill(ctx *gin.Context)
}

type Reset struct {
	*handlers.BaseHandler
	User       helpers.UserI
	Reset      helpers.ResetI
	Bloodlines gateways.Bloodlines
	Expiration time.Duration
}

func NewReset(ctx *handlers.GatewayContext) ResetI {
	stats := ctx.Stats.Clone(statsd.Prefix("api.reset"))
	return &Reset{
		BaseHandler: &handlers.BaseHandler{Stats: stats},
		User:        helpers.NewUser(ctx.Sql, ctx.S3),
		Reset:       helpers.NewReset(ctx.Sql),
		Bloodlines:  ctx.Bloodlines,
		Expiration:  time.Duration(time.Hour * 2),
	}
}

func (r *Reset) Request(ctx *gin.Context) {
	email := ctx.Query("email")

	if email == "" {
		r.UserError(ctx, "Error: must provied email parameter", nil)
		return
	}

	user, err := r.User.GetByEmail(email)
	if err != nil || user == nil {
		r.UserError(ctx, "ERROR: no user found for email", email)
		return
	}

	token := models.NewToken(email)

	err = r.Reset.Insert(token)
	if err != nil {
		r.ServerError(ctx, err, email)
		return
	}

	values := make(map[string]string)
	values["reset_link"] = fmt.Sprintf("https://expresso.store/reset/%s", token.Value)

	receipt, err := r.Bloodlines.ActivateTrigger("password_reset", &bmodels.Receipt{
		UserID: user.ID,
		Values: values,
	})

	if err != nil {
		fmt.Println(err.Error())
		r.ServerError(ctx, fmt.Errorf("Error: unable to sent reset email"), nil)
		return
	}

	r.Success(ctx, receipt)
}

func (r *Reset) Get(ctx *gin.Context) {
	value := ctx.Param("token")

	token, err := r.Reset.Get(value)
	if err != nil {
		r.ServerError(ctx, err, nil)
		return
	}

	if token == nil {
		r.NotFoundError(ctx, "Error: no token for that value")
	}

	if !r.valid(token) {
		r.UserError(ctx, "Error: token has expired, request a new one", nil)
		return
	}

	r.Success(ctx, token)
}

func (r *Reset) Fulfill(ctx *gin.Context) {
	value := ctx.Query("token")

	var json models.ResetRequest
	err := ctx.BindJSON(&json)
	if err != nil || json.PassHash == "" {
		r.UserError(ctx, "Error: unable to parse request", json)
		return
	}

	token, err := r.Reset.Get(value)
	if err != nil {
		r.ServerError(ctx, err, nil)
		return
	}

	if token == nil {
		r.NotFoundError(ctx, "Error: no token for that value")
		return
	}

	if !r.valid(token) {
		r.UserError(ctx, "Error: token has expired, request a new one", nil)
		return
	}

	user, err := r.User.GetByEmail(token.Email)
	if err != nil {
		r.ServerError(ctx, err, nil)
		return
	}

	if user == nil {
		r.NotFoundError(ctx, "Error: invlaid email")
		return
	}

	user.PassHash = json.PassHash
	err = r.User.Update(user, user.ID.String())
	if err != nil {
		r.ServerError(ctx, err, nil)
		return
	}

	r.Reset.SetStatus(token, models.INVALID)
	r.Success(ctx, nil)
}

func (r *Reset) valid(token *models.Token) bool {

	if token.Status == models.INVALID || token.Status == models.EXPIRED {
		return false
	}

	if time.Since(token.CreatedAt) >= r.Expiration {
		r.Reset.SetStatus(token, models.EXPIRED)
		return false
	}

	return true
}
