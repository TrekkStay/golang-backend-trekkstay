package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"trekkstay/modules/user/domain/usecase"
)

type UserHandler interface {
	HandleCreateUser(c *gin.Context)
	HandleLoginUser(c *gin.Context)
	HandleChangePassword(c *gin.Context)
	HandleForgotPassword(c *gin.Context)
	HandleRefreshToken(c *gin.Context)
}

type userHandler struct {
	requestValidator      *validator.Validate
	createUserUseCase     usecase.CreateUserUseCase
	loginUserUseCase      usecase.LoginUserUseCase
	changePasswordUseCase usecase.ChangePasswordUseCase
	forgotPasswordUseCase usecase.ForgotPasswordUseCase
	refreshTokenUseCase   usecase.RefreshTokenUseCase
}

func NewUserHandler(
	requestValidator *validator.Validate,
	createUserUseCase usecase.CreateUserUseCase,
	loginUserUseCase usecase.LoginUserUseCase,
	changePasswordUseCase usecase.ChangePasswordUseCase,
	forgotPasswordUseCase usecase.ForgotPasswordUseCase,
	refreshTokenUseCase usecase.RefreshTokenUseCase,
) UserHandler {
	return &userHandler{
		requestValidator:      requestValidator,
		createUserUseCase:     createUserUseCase,
		loginUserUseCase:      loginUserUseCase,
		changePasswordUseCase: changePasswordUseCase,
		forgotPasswordUseCase: forgotPasswordUseCase,
		refreshTokenUseCase:   refreshTokenUseCase,
	}
}
