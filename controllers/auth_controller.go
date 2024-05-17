package controllers

import (
	"gin-api/dto"
	"gin-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	SignUp(ctx *gin.Context)
}

type AuthController struct {
	service services.IAuthService
}

func NewAuthController(service services.IAuthService) IAuthController {
	return &AuthController{service: service}
}

func (c *AuthController) SignUp(ctx *gin.Context) {
	var input dto.SignUpInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.SignUp(input.Email, input.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "success"})
}
