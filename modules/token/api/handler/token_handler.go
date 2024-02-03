package handler

import (
	"github.com/gin-gonic/gin"
	"trekkstay/modules/token/domain/usecase"
)

type TokenHandler interface {
	HandleRefreshToken(c *gin.Context)
}

type tokenHandler struct {
	refreshTokenUseCase usecase.RefreshTokenUseCase
}

func NewTokenHandler(refreshTokenUseCase usecase.RefreshTokenUseCase) TokenHandler {
	return &tokenHandler{
		refreshTokenUseCase: refreshTokenUseCase,
	}
}
