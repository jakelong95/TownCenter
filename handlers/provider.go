package handlers

import
(
	"github.com/pborman/uuid"
	"gopkg.in/gin-gonic/gin.v1"
	
	"github.com/jakelong95/TownCenter/gateways"
)

type ProviderI struct
{
	New(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type Provider struct
{
	sql *gateways.Sql
}

func NewProvider(sql gateways.SQL) *Provider
{
	return &Provider{sql: sql}
}

func (p *Provider) New(ctx *gin.Context)
{
	//TODO
	
	ctx.JSON(200, empty())
}

func (p *Provider) ViewAll(ctx *gin.Context)
{
	//TODO
	
	ctx.JSON(200, empty())
}

func (p *Provider) View(ctx *gin.Context)
{
	//TODO
	
	ctx.JSON(200, empty())
}

func (p *Provider) Update(ctx *gin.Context)
{
	//TODO
	
	ctx.JSON(200, empty())
}

func (p *Provider) Delete(ctx *gin.Context)
{
	//TODO
	
	ctx.JSON(200, empty())
}