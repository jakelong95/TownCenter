package handlers

import
(
	"github.com/pborman/uuid"
	"gopkg.in/gin-gonic/gin.v1"
	
	"github.com/jakelong95/TownCenter/gateways"
)

type ConsumerI struct
{
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type Consumer struct
{
	sql *gateways.Sql
}

func NewConsumer(sql gateways.SQL) *Consumer
{
	return &Consumer{sql: sql}
}

func (c *Consumer) New(ctx *gin.Context)
{
	//TODO
	
	ctx.JSON(200, empty())
}

func (c *Consumer) ViewAll(ctx *gin.Context)
{
	//TODO
	
	ctx.JSON(200, empty())
}

func (c *Consumer) View(ctx *gin.Context)
{
	//TODO
	
	ctx.JSON(200, empty())
}

func (c *Consumer) Update(ctx *gin.Context)
{
	//TODO
	
	ctx.JSON(200, empty())
}

func (c *Consumer) Delete(ctx *gin.Context)
{
	//TODO
	
	ctx.JSON(200, empty())
}