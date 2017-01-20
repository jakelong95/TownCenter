package router

import (
	"fmt"

	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/config"
	"github.com/ghmeier/bloodlines/gateways"
	h "github.com/ghmeier/bloodlines/handlers"
	"github.com/jakelong95/TownCenter/handlers"
)

/* TownCenter is the main server object which routes the requests */
type TownCenter struct {
	router   *gin.Engine
	consumer handlers.ConsumerI
	provider handlers.ProviderI
}

/* Creates a ready-to-run TownCenter struct from the given config */
func New(config *config.Root) (*TownCenter, error) {
	sql, err := gateways.NewSQL(config.SQL)
	if err != nil {
		fmt.Println("ERROR: could not connect to mysql.")
		fmt.Println(err.Error())
		return nil, err
	}

	stats, err := statsd.New(
		statsd.Address(config.Statsd.Host+":"+config.Statsd.Port),
		statsd.Prefix(config.Statsd.Prefix),
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	ctx := &h.GatewayContext{
		Sql:   sql,
		Stats: stats,
	}

	tc := &TownCenter{
		consumer: handlers.NewConsumer(ctx),
		provider: handlers.NewProvider(ctx),
	}

	InitRouter(tc)

	return tc, nil
}

func InitRouter(tc *TownCenter) {
	tc.router = gin.Default()

	consumer := tc.router.Group("/api/consumer")
	{
		consumer.POST("", tc.consumer.New)
		consumer.GET("", tc.consumer.ViewAll)
		consumer.PATCH("/:consumerId", tc.consumer.Update)
		consumer.DELETE("/:consumerId", tc.consumer.Delete)
		consumer.GET("/:consumerId", tc.consumer.View)
	}

	provider := tc.router.Group("/api/provider")
	{
		provider.POST("", tc.provider.New)
		provider.GET("", tc.provider.ViewAll)
		provider.PATCH("/:providerId", tc.provider.Update)
		provider.DELETE("/:providerId", tc.provider.Delete)
		provider.GET("/:providerId", tc.provider.View)
	}
}

/* Starts the TownCenter server */
func (tc *TownCenter) Start(port string) {
	tc.router.Run(port)
}
