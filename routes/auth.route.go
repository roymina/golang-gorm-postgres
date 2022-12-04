package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/roymina/golang-gorm-postgres/controllers"
)

type AuthRoute struct {
	authController controllers.AuthController
}

func NewAuthRoute(authController controllers.AuthController) AuthRoute {
	return AuthRoute{authController}
}

func (rc *AuthRoute) AuthRoute(rg *gin.RouterGroup){
	router:=rg.Group("/auth")

	router.POST("/register",rc.authController.SignUpUser)
	router.POST("/login",rc.authController.SignInUser)
	router.GET("/refresh",rc.authController.RefreshAccessToken)
	router.GET("/logout",rc.authController.LogoutUser)

}