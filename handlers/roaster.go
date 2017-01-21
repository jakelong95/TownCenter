package handlers

import (
	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/handlers"
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
}

func NewRoaster(ctx *handlers.GatewayContext) RoasterI {
	stats := ctx.Stats.Clone(statsd.Prefix("api.roaster"))
	return &Roaster{
		BaseHandler: &handlers.BaseHandler{Stats: stats},
	}
}

func (p *Roaster) New(ctx *gin.Context) {
	//TODO

	p.Success(ctx, nil)
}

func (p *Roaster) ViewAll(ctx *gin.Context) {
	//TODO

	p.Success(ctx, nil)
}

func (p *Roaster) View(ctx *gin.Context) {
	//TODO

	p.Success(ctx, nil)
}

func (p *Roaster) Update(ctx *gin.Context) {
	//TODO

	p.Success(ctx, nil)
}

func (p *Roaster) Delete(ctx *gin.Context) {
	//TODO

	p.Success(ctx, nil)
}
