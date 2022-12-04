package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/roymina/golang-gorm-postgres/controllers"
	"github.com/roymina/golang-gorm-postgres/initializers"
	"github.com/roymina/golang-gorm-postgres/routes"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRoute routes.AuthRoute

	UserController      controllers.UserController
	UserRoute routes.UserRoute
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRoute = routes.NewAuthRoute(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRoute = routes.NewUserRoute(UserController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig :=cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials=true

	server.Use(cors.New(corsConfig))

	router:=server.Group("/api")

	
	router.GET("/health", func(ctx *gin.Context) {
		message := "welcome to golang gorm postgres server"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	AuthRoute.AuthRoute(router)
	UserRoute.UserRoute(router)
	
	log.Fatal(server.Run(":" + config.ServerPort))
}
