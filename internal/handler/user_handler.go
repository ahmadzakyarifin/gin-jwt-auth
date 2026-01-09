package handler

import (
	"net/http"

	"github.com/ahmadzakyarifin/gin-jwt-auth/internal/dto"
	"github.com/ahmadzakyarifin/gin-jwt-auth/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.AuthService
}

func NewUserHandler(s service.AuthService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Register(ctx *gin.Context){
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.Register(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Register Success",
	})
}

func (h *UserHandler) Login(ctx *gin.Context){
	var req dto.LoginRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error" : err.Error(),
		})
		return
	}
	


	_,err = h.service.Login(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error" : err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"message" : "Selamat datang di APK ini",
	})
}
