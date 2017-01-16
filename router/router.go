package router

import (
	"fmt"
	
	"gopkg.in/gin-gonic/gin.v1"
	
	"github.com/jakelong95/TownCenter/handlers"
	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/config"
)

/* TownCenter is the main server object which routes the requests */
type TownCenter struct {
	router		  *gin.Engine
	consumer	  handlers.ConsumerI
	provider	  handlers.ProviderI
}

/* Creates a ready-to-run TownCenter struct from the given config */
func New(config *config.Root) (*TownCenter, error) {
	sql, err := gateways.NewSQL(config.SQL)
	if err != nil {
		fmt.Println("Error: Could not connect to MySQL")
		fmt.Println(err.Error())
		return nil, err
	}
	
	tc := &TownCenter {
		consumer: handlers.NewConsumer(sql),
		provider: handlers.NewProvider(sql),
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