package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/roymina/golang-gorm-postgres/controllers"
	"github.com/roymina/golang-gorm-postgres/middlewares"
)

type UserRoute struct {
	userController controllers.UserController
}

func NewUserRoute(userController controllers.UserController) UserRoute {
	return UserRoute{userController}
}

func (uc *UserRoute) UserRoute(rg *gin.RouterGroup){
	router:=rg.Group("users")
	router.GET("/me",middleware.DeserializeUser(),uc.userController.GetMe)
}