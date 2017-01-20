package handlers

import (
	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/handlers"
)

type ProviderI interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type Provider struct {
	*handlers.BaseHandler
}

func NewProvider(ctx *handlers.GatewayContext) ProviderI {
	stats := ctx.Stats.Clone(statsd.Prefix("api.provider"))
	return &Provider{
		BaseHandler: &handlers.BaseHandler{Stats: stats},
	}
}

func (p *Provider) New(ctx *gin.Context) {
	//TODO

	p.Success(ctx, nil)
}

func (p *Provider) ViewAll(ctx *gin.Context) {
	//TODO

	p.Success(ctx, nil)
}

func (p *Provider) View(ctx *gin.Context) {
	//TODO

	p.Success(ctx, nil)
}

func (p *Provider) Update(ctx *gin.Context) {
	//TODO

	p.Success(ctx, nil)
}

func (p *Provider) Delete(ctx *gin.Context) {
	//TODO

	p.Success(ctx, nil)
}
