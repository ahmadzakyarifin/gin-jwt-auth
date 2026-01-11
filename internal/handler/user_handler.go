package handler

import (
	"net/http"

	"github.com/ahmadzakyarifin/gin-jwt-auth/internal/dto"
	"github.com/ahmadzakyarifin/gin-jwt-auth/internal/service"
	"github.com/ahmadzakyarifin/gin-jwt-auth/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
    service service.AuthService
}

func NewUserHandler(s service.AuthService) *UserHandler {
    return &UserHandler{service: s}
}

func (h *UserHandler) Register(ctx *gin.Context) {
    var req dto.RegisterRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
		errors := utils.FormatValidationEror(err)
        utils.APIErrorResponse(ctx, "Input tidak valid", http.StatusBadRequest, "error", errors)
        return
    }

    if err := h.service.Register(&req); err != nil {
        utils.APIErrorResponse(ctx, "Register Gagal", http.StatusBadRequest, "error", nil)
        return
    }

    utils.APIResponse(ctx, "Register Berhasil", http.StatusCreated, "success", nil)
}

func (h *UserHandler) Login(ctx *gin.Context) {
    var req dto.LoginRequest

    err := ctx.ShouldBindJSON(&req)
    if err != nil {
		errors := utils.FormatValidationEror(err)
        utils.APIErrorResponse(ctx, "Input tidak valid", http.StatusBadRequest, "error", errors)
        return
    }

    token, err := h.service.Login(&req)
    if err != nil {
        utils.APIErrorResponse(ctx, "Login Gagal", http.StatusUnauthorized, "error", nil)
        return
    }
    ctx.SetCookie("token",token,3600,"/","localhost",false,true)

    utils.APIResponse(ctx, "Login Berhasil", http.StatusOK, "success", gin.H{"token": token})
}