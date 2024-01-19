package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"trekkstay/modules/user/domain/usecase"
)

type UserHandler interface {
	HandleCreateUser(c *gin.Context)
}

type userHandler struct {
	requestValidator  *validator.Validate
	createUserUseCase usecase.CreateUserUseCase
}

func NewUserHandler(
	requestValidator *validator.Validate,
	createUserUseCase usecase.CreateUserUseCase,
) UserHandler {
	return &userHandler{
		requestValidator:  requestValidator,
		createUserUseCase: createUserUseCase,
	}
}
