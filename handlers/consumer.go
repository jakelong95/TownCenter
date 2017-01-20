package handlers

import (
	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/handlers"
)

type ConsumerI interface {
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type Consumer struct {
	*handlers.BaseHandler
}

func NewConsumer(ctx *handlers.GatewayContext) ConsumerI {
	stats := ctx.Stats.Clone(statsd.Prefix("api.consumer"))
	return &Consumer{
		BaseHandler: &handlers.BaseHandler{Stats: stats},
	}
}

func (c *Consumer) New(ctx *gin.Context) {
	//TODO

	c.Success(ctx, nil)
}

func (c *Consumer) ViewAll(ctx *gin.Context) {
	//TODO

	c.Success(ctx, nil)
}

func (c *Consumer) View(ctx *gin.Context) {
	//TODO

	c.Success(ctx, nil)
}

func (c *Consumer) Update(ctx *gin.Context) {
	//TODO

	c.Success(ctx, nil)
}

func (c *Consumer) Delete(ctx *gin.Context) {
	//TODO

	c.Success(ctx, nil)
}
