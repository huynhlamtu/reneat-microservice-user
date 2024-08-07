package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"reneat-microservice-user/config"
	"reneat-microservice-user/controllers"
	"reneat-microservice-user/docs"
	"reneat-microservice-user/middleware"
)

func RouteInit(engine *gin.Engine) {
	authCtr := new(controllers.AuthController)

	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Management Promotion Service API")
	})
	engine.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	engine.Use(middleware.Recovery())
	engine.Use(middleware.CORSMiddleware())

	docs.SwaggerInfo.BasePath = "/v1"

	cfg := config.GetConfig()
	if cfg.GetString("env.mode") != "debug" {
		serviceType := cfg.GetString("server.permission_service.type")
		serviceName := cfg.GetString("server.permission_service.name")
		docs.SwaggerInfo.BasePath = "/" + serviceType + "-" + serviceName + "/v1"
	}

	apiV1 := engine.Group("/v1")

	apiV1.Use(middleware.RequestLog())
	{
		apiV1.POST("/users/register", authCtr.Register)
		apiV1.POST("/users/login", authCtr.Login)
		apiV1.GET("/users/user/:username", authCtr.Detail)
		apiV1.GET("/users/info/:uuid", authCtr.Info)
	}

	//apiV1.Use(middleware.VerifyAuth())

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
