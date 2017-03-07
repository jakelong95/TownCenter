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
	router  *gin.Engine
	user    handlers.UserI
	roaster handlers.RoasterI
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

	bloodlines := gateways.NewBloodlines(config.Bloodlines)
	if err != nil {
		fmt.Println(err.Error())
	}

	ctx := &h.GatewayContext{
		Sql:   		  sql,
		Stats: 		  stats,
		Bloodlines: bloodlines,
	}

	tc := &TownCenter{
		user:    handlers.NewUser(ctx),
		roaster: handlers.NewRoaster(ctx),
	}

	InitRouter(tc)

	return tc, nil
}

func InitRouter(tc *TownCenter) {
	tc.router = gin.Default()
	tc.router.Use(h.GetCors())

	user := tc.router.Group("/api/user")
	{
		user.Use(tc.user.Time())
		user.POST("", tc.user.New)
		user.POST("/login", tc.user.Login)
		user.Use(tc.user.GetJWT())
		user.GET("", tc.user.ViewAll)
		user.PUT("/:userId", tc.user.Update)
		user.DELETE("/:userId", tc.user.Delete)
		user.GET("/:userId", tc.user.View)

	}

	roaster := tc.router.Group("/api/roaster")
	{
		roaster.Use(tc.roaster.GetJWT())
		roaster.Use(tc.roaster.Time())
		roaster.POST("", tc.roaster.New)
		roaster.GET("", tc.roaster.ViewAll)
		roaster.PUT("/:roasterId", tc.roaster.Update)
		roaster.DELETE("/:roasterId", tc.roaster.Delete)
		roaster.GET("/:roasterId", tc.roaster.View)
	}
}

/* Starts the TownCenter server */
func (tc *TownCenter) Start(port string) {
	tc.router.Run(port)
}
