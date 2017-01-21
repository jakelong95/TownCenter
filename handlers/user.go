package handlers

import (
	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/handlers"
)

type UserI interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type User struct {
	*handlers.BaseHandler
}

func NewUser(ctx *handlers.GatewayContext) UserI {
	stats := ctx.Stats.Clone(statsd.Prefix("api.user"))
	return &User{
		BaseHandler: &handlers.BaseHandler{Stats: stats},
	}
}

func (c *User) New(ctx *gin.Context) {
	//TODO

	c.Success(ctx, nil)
}

func (c *User) ViewAll(ctx *gin.Context) {
	//TODO

	c.Success(ctx, nil)
}

func (c *User) View(ctx *gin.Context) {
	//TODO

	c.Success(ctx, nil)
}

func (c *User) Update(ctx *gin.Context) {
	//TODO

	c.Success(ctx, nil)
}

func (c *User) Delete(ctx *gin.Context) {
	//TODO

	c.Success(ctx, nil)
}
