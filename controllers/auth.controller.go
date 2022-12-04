package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/roymina/golang-gorm-postgres/models"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

func (ac *AuthController) SignUp(ctx *gin.Context){
	var payload *models.SignUpInput
	if err:=ctx.ShouldBindJSON(&payload);err!=nil{
		
	}
}